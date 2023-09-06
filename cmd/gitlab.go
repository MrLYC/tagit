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

		pushTo, _ := flags.GetString("push-to")
		repo := manager.Repository()
		remoteUrl, err := repo.GetRemoteUrl(pushTo)
		checkError(err)

		projectID, _ := flags.GetInt("project-id")
		if projectID == 0 {
			projectInfo, err := gitlab.GetProjectByName(projectName, tagit.ProjectUrlDistance(remoteUrl))
			checkError(err)

			projectID = projectInfo.ID
			fmt.Printf("Found project[%s] id=%d\n", projectInfo.NameWithNamespace, projectID)
		}

		tagger := tagit.NewGitlabTagger(projectID, gitlab)
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
	flags.String("push-to", "", "remote name")
	flags.String("project", "", "gitlab project name")
	flags.Int("project-id", 0, "gitlab project id")
	flags.String("gitlab-url", "", "gitlab url")
	flags.String("gitlab-token", "", "gitlab token")

	_ = gitlabCmd.MarkFlagRequired("name")
	_ = gitlabCmd.MarkFlagRequired("gitlab-url")
	_ = gitlabCmd.MarkFlagRequired("gitlab-token")
}
