package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gotree",
	Short:   "gotree is a tree command using go",
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		if enabledColor, err := cmd.PersistentFlags().GetBool("color"); err == nil && enabledColor {
			setColor()
		}
		arg := getArg(args)
		w := TreeWriter{}
		Tree(w, arg)
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("color", true, "color output")
}

// Execute exec command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getArg(args []string) string {
	var arg string
	if len(args) == 0 {
		arg = "."
	} else {
		arg = args[0]
	}
	return arg
}
