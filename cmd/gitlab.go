package cmd

import (
	"fmt"
	"path"

	"github.com/mrlyc/tagit/tagit"
	"github.com/spf13/cobra"
)

// gitlabCmd create tag directly
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Create gitlab tag",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		projectPath, _ := flags.GetString("path")
		tagName, _ := flags.GetString("name")
		tagMessage, _ := flags.GetString("message")

		projectName, _ := flags.GetString("project")
		if projectName == "" {
			projectName = path.Base(projectPath)
		}

		manager, err := tagit.NewManager(projectPath)
		checkError(err)

		token, _ := flags.GetString("gitlab-token")
		url, _ := flags.GetString("gitlab-url")
		gitlab, err := tagit.NewGitlab(token, url)
		checkError(err)

		pid, err := gitlab.GetProjectIDByName(projectName, tagit.ProjectUrlDistance(projectName))
		checkError(err)

		fmt.Printf("Found project id: %d\n", pid)

		tagger := tagit.NewGitlabTagger(pid, gitlab)
		commitID, err := manager.TagHead(tagger, tagName, tagMessage)
		checkError(err)

		fmt.Printf("Tag name[%s] of project[%s] on commit[%s]\n", tagName, projectName, commitID)
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)

	flags := gitlabCmd.Flags()
	flags.StringP("path", "p", ".", "path to git repository")
	flags.StringP("name", "n", "", "tag name")
	flags.StringP("message", "m", "-", "tag message")
	flags.StringP("project", "P", "", "project name")
	flags.String("gitlab-url", "", "gitlab url")
	flags.String("gitlab-token", "", "gitlab token")

	_ = gitlabCmd.MarkFlagRequired("name")
	_ = gitlabCmd.MarkFlagRequired("gitlab-url")
	_ = gitlabCmd.MarkFlagRequired("gitlab-token")
}
