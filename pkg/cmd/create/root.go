package create

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new projects or resources",
	Long: `Choose a resource to create, e.g.
	nitric create stack
`,
}

// Flags
var force bool

var createStackCmd = &cobra.Command{
	Use:   "stack [name] [template]",
	Short: "create a new application stack",
	Long:  `Creates a new Nitric application stack from a template.`,
	Run: func(cmd *cobra.Command, args []string) {
		notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
		notice("Don't forget this... %v", force)
	},
	Args: cobra.MaximumNArgs(2),
}

func init() {
	createStackCmd.Flags().BoolVarP(&force, "force", "f", false, "force stack creation, even in non-empty directories.")
	createCmd.AddCommand(createStackCmd)
}

func RootCommand() *cobra.Command {
	return createCmd
}
