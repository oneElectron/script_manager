package main

import (
	"context"
	"fmt"

	github "github.com/oneElectron/script_manager/internal/github_connection"

	"github.com/spf13/cobra"
)

// gistsCmd represents the gists command
var gistsCmd = &cobra.Command{
	Use:   "gists",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		authKey := viperConf.GetString("GithubAuthToken")
		if authKey == "PLACEHOLDER" {
			println("You are not authenticated, please run smctl auth")
			return
		}

		ctx := context.Background()
		gists, err := github.ListGists(ctx)
		if err != nil {
			println(err.Error())
			return
		}

		for _, gist := range gists {
			fmt.Println(*gist.Description)
		}


	},
}

func init() {
	rootCmd.AddCommand(gistsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gistsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gistsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
