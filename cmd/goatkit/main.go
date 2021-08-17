package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/togettoyou/goat/cmd/goatkit/internal/project"
	"github.com/togettoyou/goat/cmd/goatkit/internal/run"
	"github.com/togettoyou/goat/cmd/goatkit/internal/swag"
	"github.com/togettoyou/goat/cmd/goatkit/internal/upgrade"
)

var (
	version = "v1.0.0"

	rootCmd = &cobra.Command{
		Use:     "goatkit",
		Short:   "goatkit: An elegant toolkit for Go services.",
		Long:    `goatkit: An elegant toolkit for Go services.`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(swag.CmdSwag)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
