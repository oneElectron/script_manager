package main

import (
	"os"
	"path"

	"github.com/mattn/go-isatty"
	"github.com/oneElectron/script_manager/internal/smgithub"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with GitHub",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		authKey := smgithub.AuthFlow(isatty.IsCygwinTerminal(os.Stdout.Fd()) || isatty.IsTerminal(os.Stdout.Fd()))

		viperConf.Set("GithubAuthToken", authKey.AuthKey)
		viperConf.Set("GithubUsername", authKey.Username)

		viperConf.WriteConfigAs(path.Join(DIRS.DataHome, "script_manager/sm.toml"))
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
