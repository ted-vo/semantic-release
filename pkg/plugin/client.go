package plugin

import (
	"os/exec"
	"sync"

	"github.com/go-semantic-release/semantic-release/pkg/analyzer/commit"
	"github.com/go-semantic-release/semantic-release/pkg/plugin/wrapper"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type PluginOpts struct {
	Type string
	Cmd  *exec.Cmd
}

var runningClientsMx sync.Mutex
var runningClients = make([]*plugin.Client, 0)

func KillAllPlugins() {
	for _, c := range runningClients {
		c.Kill()
	}
}

func startPlugin(opts *PluginOpts) (interface{}, error) {
	runningClientsMx.Lock()
	defer runningClientsMx.Unlock()
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: {
				opts.Type: &wrapper.GRPC{
					Type: opts.Type,
				},
			},
		},
		Cmd:              opts.Cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           hclog.NewNullLogger(),
	})

	rpcClient, err := client.Client()
	if err != nil {
		client.Kill()
		panic(err)
	}
	raw, err := rpcClient.Dispense(opts.Type)
	if err != nil {
		client.Kill()
		return nil, err
	}
	runningClients = append(runningClients, client)
	return raw, nil
}

func StartCommitAnalyzerPlugin(opts *PluginOpts) (commit.Analyzer, error) {
	raw, err := startPlugin(opts)
	if err != nil {
		return nil, err
	}
	return raw.(commit.Analyzer), nil
}