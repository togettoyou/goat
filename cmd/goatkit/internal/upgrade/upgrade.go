package upgrade

import (
	"fmt"
	"github.com/togettoyou/goat/cmd/goatkit/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the goatkit tools",
	Long:  "Upgrade the goatkit tools. Example: goatkit upgrade",
	Run:   Run,
}

// Run upgrade the goatkit tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoGet(
		"github.com/togettoyou/goat/cmd/goatkit",
	)
	if err != nil {
		fmt.Println(err)
	}
}
