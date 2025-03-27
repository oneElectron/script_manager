package main

import (
	"fmt"
	"log/slog"

	"github.com/oneElectron/script_manager/internal/scriptDB"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a script",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := scriptDB.FindDatabase()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		from := args[0]
		to := args[1]

		err = db.RenameScript(from, to)

		if err != nil {
			fmt.Println(err.Error())
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
