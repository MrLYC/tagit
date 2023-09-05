package cmd

import (
	"fmt"

	tagit "github.com/mrlyc/tagit/taggit"
	"github.com/spf13/cobra"
)

// commitIDCmd print current commit id directly
var commitIDCmd = &cobra.Command{
	Use:   "commit-id",
	Short: "Print current commit id directly",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		path, _ := flags.GetString("path")

		commitID, err := tagit.GetCurrentCommitID(path)
		if err != nil {
			panic(err)
		}

		fmt.Println(commitID)
	},
}

func init() {
	rootCmd.AddCommand(commitIDCmd)

	flags := commitIDCmd.Flags()
	flags.StringP("path", "p", ".", "path to git repository")
}
