package cmd

import (
	"context"

	"github.com/ivange94/junkdb/internal/buildinfo"
	"github.com/ivange94/junkdb/internal/config"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use: "junkdb", //TODO: add more fields for documentation
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { //TODO: verify that this is called only once and not for command and all subcommands
		if cmd == versionCmd {
			return nil
		}
		cfg := config.MustLoad()
		cmd.SetContext(withConfig(cmd.Context(), cfg))
		return nil
	},
	Version: buildinfo.Version,
}

type cfgKey struct{}

func withConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, cfgKey{}, cfg)
}

func mustConfig(ctx context.Context) *config.Config {
	cfg, ok := ctx.Value(cfgKey{}).(*config.Config)
	if !ok || cfg == nil {
		panic("config missing from context")
	}
	return cfg
}
