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

		manager, err := tagit.NewManager(projectPath)
		checkError(err)

		repo := manager.Repository()

		commitID, err := manager.TagHead(repo, tagName, tagMessage)
		checkError(err)

		fmt.Printf("Tag name[%s] of project[%s] on commit[%s]\n", tagName, projectPath, commitID)

		pushTo, _ := flags.GetString("push-to")
		if pushTo != "" {
			err = repo.PushTo(pushTo)
			checkError(err)

			fmt.Printf("Pushed to remote[%s]\n", pushTo)
		}
	},
}

func init() {
	rootCmd.AddCommand(localCmd)

	flags := localCmd.Flags()
	flags.StringP("path", "p", ".", "path to git repository")
	flags.StringP("name", "n", "", "tag name")
	flags.StringP("message", "m", "-", "tag message")
	flags.String("push-to", "", "remote name")

	_ = localCmd.MarkFlagRequired("name")
}
