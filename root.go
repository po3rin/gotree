package gotree

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gotree",
	Short:   "gotree is a tree command using go",
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		if enabledColor, err := cmd.PersistentFlags().GetBool("c"); err == nil && enabledColor {
			setColor()
		}
		if level, err := cmd.PersistentFlags().GetInt("l"); err == nil && level > 0 {
			setLevel(level)
		}
		if d, err := cmd.PersistentFlags().GetBool("d"); err == nil && d {
			setDirectoryOnly()
		}
		arg := getArg(args)
		w := writer{}
		Tree(w, arg)
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("c", true, "Display with color")
	rootCmd.PersistentFlags().Int("l", 10, "Max display depth of the directory tree")
	rootCmd.PersistentFlags().Bool("d", false, "Display directory only")
}

// Execute exec command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
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
