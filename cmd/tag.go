package cmd

import (
	tagit "github.com/mrlyc/tagit/taggit"
	"github.com/mrlyc/tagit/taggit/taggers/local"
	"github.com/spf13/cobra"
)

// tagCmd create tag directly
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create tag directly",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		path, _ := flags.GetString("path")

		commitID, err := tagit.GetCurrentCommitID(path)
		if err != nil {
			panic(err)
		}

		tagger, err := local.NewTagger(path)
		if err != nil {
			panic(err)
		}

		name, _ := flags.GetString("name")
		message, _ := flags.GetString("message")
		err = tagit.CreateTag(tagger, commitID, name, message)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)

	flags := tagCmd.Flags()
	flags.StringP("path", "p", ".", "path to git repository")
	flags.StringP("name", "n", "", "tag name")
	flags.StringP("message", "m", "", "tag message")

	_ = rootCmd.MarkFlagRequired("name")
	_ = rootCmd.MarkFlagRequired("message")
}
