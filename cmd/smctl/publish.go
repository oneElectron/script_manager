package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/oneElectron/script_manager/internal/scriptDB"
	"github.com/oneElectron/script_manager/internal/smgithub"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := scriptDB.FindDatabase()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		ctx := context.Background()
		initGithub()

		public, err := cmd.Flags().GetBool("public")
		if err != nil {
			slog.Error(err.Error())
			return
		}

		user, err := smgithub.GetUsername(ctx)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		name := args[0]
		script, err := db.FindLocalScript(name)

		desc := name + ".sm"
		contents, err := os.ReadFile(script.OsPath)

		if err != nil {
			slog.Error(err.Error())
			return
		}

		cmap := map[string]string{name:string(contents)}

		smgithub.CreateGist(ctx, desc , cmap,  public)
		db.ConvertLocalToOnline(name, "github.com", user)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	publishCmd.Flags().BoolP("public", "p", false, "Make the gist public")
}
