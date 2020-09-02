package discovery

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/cavaliercoder/grab"
	"github.com/go-semantic-release/semantic-release/v2/pkg/analyzer"
	"github.com/go-semantic-release/semantic-release/v2/pkg/condition"
	"github.com/go-semantic-release/semantic-release/v2/pkg/generator"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
	"github.com/go-semantic-release/semantic-release/v2/pkg/provider"
	"github.com/go-semantic-release/semantic-release/v2/pkg/updater"
)

const PluginDir = ".semrel"

var osArchDir = runtime.GOOS + "_" + runtime.GOARCH

func getPluginPath(name string) string {
	pElem := append([]string{PluginDir}, osArchDir, name)
	return path.Join(pElem...)
}

func ensurePluginDir(pth string) error {
	_, err := os.Stat(pth)
	if os.IsNotExist(err) {
		return os.MkdirAll(pth, 0755)
	}
	return err
}

type apiPluginAsset struct {
	FileName string
	URL      string
	OS       string
	Arch     string
	Checksum string
}

type apiPluginRelease struct {
	CreatedAt time.Time
	Assets    []*apiPluginAsset
}

func (r *apiPluginRelease) getMatchingAsset() *apiPluginAsset {
	for _, a := range r.Assets {
		if a.OS == runtime.GOOS && a.Arch == runtime.GOARCH {
			return a
		}
	}
	return nil
}

type apiPlugin struct {
	Type          string
	Name          string
	LatestRelease string
	Versions      map[string]*apiPluginRelease
}

func getPluginInfo(name string) (*apiPlugin, error) {
	res, err := http.Get(fmt.Sprintf("https://plugins.go-semantic-release.xyz/api/v1/plugins/%s.json", name))
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.New("invalid response")
	}
	defer res.Body.Close()
	var plugin *apiPlugin
	if err := json.NewDecoder(res.Body).Decode(&plugin); err != nil {
		return nil, err
	}
	return plugin, nil
}

func fetchPlugin(name, pth string, cons *semver.Constraints) (string, error) {
	pluginInfo, err := getPluginInfo(name)
	if err != nil {
		return "", err
	}

	foundVersion := ""
	if cons == nil {
		foundVersion = pluginInfo.LatestRelease
	} else {
		versions := make(semver.Collection, 0)
		for v := range pluginInfo.Versions {
			pv, err := semver.NewVersion(v)
			if err != nil {
				return "", err
			}
			versions = append(versions, pv)
		}
		sort.Sort(sort.Reverse(versions))
		for _, v := range versions {
			if cons.Check(v) {
				foundVersion = v.String()
				break
			}
		}
	}

	if foundVersion == "" {
		return "", errors.New("version not found")
	}

	releaseAsset := pluginInfo.Versions[foundVersion].getMatchingAsset()
	if releaseAsset == nil {
		return "", errors.New("release not found")
	}

	targetPath := path.Join(pth, foundVersion, releaseAsset.FileName)

	req, err := grab.NewRequest(targetPath, releaseAsset.URL)
	if err != nil {
		return "", err
	}
	if releaseAsset.Checksum != "" {
		sum, err := hex.DecodeString(releaseAsset.Checksum)
		if err != nil {
			return "", err
		}
		req.SetChecksum(sha256.New(), sum, true)
	}

	res := grab.DefaultClient.Do(req)
	if err := res.Err(); err != nil {
		return "", err
	}
	if err := os.Chmod(res.Filename, 0755); err != nil {
		return "", err
	}

	return res.Filename, nil
}

func getMatchingVersionDir(pth string, cons *semver.Constraints) (string, error) {
	vDirs, err := ioutil.ReadDir(pth)
	if err != nil {
		return "", err
	}
	foundVers := make(semver.Collection, 0)
	for _, f := range vDirs {
		if f.IsDir() {
			fVer, err := semver.NewVersion(f.Name())
			if err != nil {
				continue
			}
			foundVers = append(foundVers, fVer)
		}
	}

	if len(foundVers) == 0 {
		return "", errors.New("no installed version found")
	}
	sort.Sort(sort.Reverse(foundVers))

	if cons == nil {
		return path.Join(pth, foundVers[0].String()), nil
	}

	for _, v := range foundVers {
		if cons.Check(v) {
			return path.Join(pth, v.String()), nil
		}
	}
	return "", errors.New("no matching version found")
}

func findPluginLocally(pth string, cons *semver.Constraints) (string, error) {
	vPth, err := getMatchingVersionDir(pth, cons)
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(vPth)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", errors.New("no plugins found")
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if f.Mode()&0100 == 0 {
			continue
		}
		return path.Join(vPth, f.Name()), nil
	}
	return "", errors.New("no matching plugin found")
}

func getPluginType(t string) string {
	switch t {
	case analyzer.CommitAnalyzerPluginName:
		return "commit-analyzer"
	case condition.CIConditionPluginName:
		return "condition"
	case generator.ChangelogGeneratorPluginName:
		return "changelog-generator"
	case provider.PluginName:
		return "provider"
	case updater.FilesUpdaterPluginName:
		return "files-updater"
	}
	return ""
}

func FindPlugin(t, name string) (*plugin.PluginOpts, error) {
	pType := getPluginType(t)
	if pType == "" {
		return nil, errors.New("invalid plugin type")
	}

	var cons *semver.Constraints
	if ve := strings.SplitN(name, "@", 2); len(ve) > 1 {
		v, err := semver.NewConstraint(ve[1])
		if err != nil {
			return nil, err
		}
		name = ve[0]
		cons = v
	}

	pName := strings.ToLower(pType + "-" + name)
	pPath := getPluginPath(pName)
	if err := ensurePluginDir(pPath); err != nil {
		return nil, err
	}

	binPath, err := findPluginLocally(pPath, cons)
	if err != nil {
		binPath, err = fetchPlugin(pName, pPath, cons)
		if err != nil {
			return nil, err
		}
	}

	return &plugin.PluginOpts{
		Type: t,
		Cmd:  exec.Command(binPath),
	}, nil
}
