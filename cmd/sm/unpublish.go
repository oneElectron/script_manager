package main

import (
	"context"
	"log/slog"

	"github.com/google/go-github/v70/github"
	"github.com/oneElectron/script_manager/internal/scriptDB"
	"github.com/oneElectron/script_manager/internal/smgithub"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Unpublish a script",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		db, err := scriptDB.FindDatabase()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		initGithub()

		name := args[0]

		list, err := smgithub.ListGists(ctx)
		if err != nil {
			println(err.Error())
			return
		}

		var gist *github.Gist = nil
		for _, item := range list {
			if *item.Description == name + ".sm" {
				gist = item
				break
			}
		}

		if gist == nil {
			return
		}


		user, err := smgithub.GetUsername(ctx)
		if err != nil {
			println(err.Error())
			return
		}

		smgithub.RemoveGist(ctx, *gist.ID)

		db.ConvertOnlineToLocal("github.com", user, name)
	},
}

func init() {
	rootCmd.AddCommand(unpublishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("public", "p", false, "Make the gist public")
}
