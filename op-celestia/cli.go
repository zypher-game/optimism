package celestia

import (
	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
)

const (
	// RPCFlagName defines the flag for the rpc url
	RPCFlagName = "da.rpc"
	// AuthTokenFlagName defines the flag for the auth token
	AuthTokenFlagName = "da.auth_token"
	// NamespaceFlagName defines the flag for the namespace
	NamespaceFlagName = "da.namespace"
	// EthFallbackDisabledFlagName defines the flag for disabling eth fallback
	EthFallbackDisabledFlagName = "da.eth_fallback_disabled"

	// NamespaceSize is the size of the hex encoded namespace string
	NamespaceSize = 58

	// defaultRPC is the default rpc dial address
	defaultRPC = "grpc://localhost:26650"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    RPCFlagName,
			Usage:   "dial address of the data availability rpc client; supports grpc, http, https",
			Value:   defaultRPC,
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_RPC"),
		},
		&cli.StringFlag{
			Name:    AuthTokenFlagName,
			Usage:   "authentication token of the data availability client",
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_AUTH_TOKEN"),
		},
		&cli.StringFlag{
			Name:    NamespaceFlagName,
			Usage:   "namespace of the data availability client",
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_NAMESPACE"),
		},
		&cli.BoolFlag{
			Name:    EthFallbackDisabledFlagName,
			Usage:   "disable eth fallback",
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_ETH_FALLBACK_DISABLED"),
		},
	}
}

type Config struct {
	Rpc                 string
	AuthToken           string
	Namespace           string
	EthFallbackDisabled bool
}

func (c Config) Check() error {
	return nil
}

func NewCLIConfig() Config {
	return Config{
		Rpc: defaultRPC,
	}
}

func ReadCLIConfig(ctx *cli.Context) Config {
	return Config{
		Rpc:                 ctx.String(RPCFlagName),
		AuthToken:           ctx.String(AuthTokenFlagName),
		Namespace:           ctx.String(NamespaceFlagName),
		EthFallbackDisabled: ctx.Bool(EthFallbackDisabledFlagName),
	}
}
