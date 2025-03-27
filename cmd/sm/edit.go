/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"io/fs"
	"log/slog"
	"os"
	"path"

	"github.com/oneElectron/script_manager/internal/edit"
	"github.com/oneElectron/script_manager/internal/scriptDB"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a script",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		db, err := scriptDB.FindDatabase()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		err = os.Mkdir(path.Join(db.LocalRoot()), fs.ModeDir|0o755)

		err = edit.Editor(path.Join(db.LocalRoot(), args[0]))
		if err != nil {
			print(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
