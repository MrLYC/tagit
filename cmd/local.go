package cmd

import (
	"fmt"

	"github.com/mrlyc/tagit/tagit"
	"github.com/spf13/cobra"
)

// localCmd create tag directly
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Create local tag",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		projectPath, _ := flags.GetString("path")
		tagName, _ := flags.GetString("name")
		tagMessage, _ := flags.GetString("message")

		repo, err := tagit.NewRepository(projectPath)
		checkError(err)

		commitID, err := tagit.TagProject(repo, repo, tagName, tagMessage)
		checkError(err)

		fmt.Printf("Tag name[%s] of project[%s] on commit[%s]\n", tagName, projectPath, commitID)
	},
}

func init() {
	rootCmd.AddCommand(localCmd)

	flags := localCmd.Flags()
	flags.StringP("path", "p", ".", "path to git repository")
	flags.StringP("name", "n", "", "tag name")
	flags.StringP("message", "m", "-", "tag message")

	_ = localCmd.MarkFlagRequired("name")
}
