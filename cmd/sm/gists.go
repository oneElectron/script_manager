package main

import (
	"context"
	"fmt"

	"github.com/oneElectron/script_manager/internal/smgithub"

	"github.com/spf13/cobra"
)

// gistsCmd represents the gists command
var gistsCmd = &cobra.Command{
	Use:   "gists",
	Short: "List your gists",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		initGithub()

		ctx := context.Background()

		gists, err := smgithub.ListGists(ctx)
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
	// rootCmd.AddCommand(gistsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gistsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gistsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
