package main

import (
	"fmt"
	"log/slog"

	"github.com/oneElectron/script_manager/internal/scriptDB"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your scripts",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := scriptDB.FindDatabase()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		list, err := db.ListScripts()
		if err != nil {
			print(err)
			return
		}

		for _, item := range list {
			if item.OnlinePath() == "" {
				fmt.Printf("%s (local)\n", item.Name)
			} else {
				fmt.Printf("%s (%s)\n", item.Name, item.OnlinePath())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
