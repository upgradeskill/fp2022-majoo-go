package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "majoo",
	Short: `
     _____ ______    ________         ___   ________   ________
    |\   _ \  _   \ |\   __  \       |\  \ |\   __  \ |\   __  \
    \ \  \\\__\ \  \\ \  \|\  \      \ \  \\ \  \|\  \\ \  \|\  \
     \ \  \\|__| \  \\ \   __  \   __ \ \  \\ \  \\\  \\ \  \\\  \
      \ \  \    \ \  \\ \  \ \  \ |\  \\_\  \\ \  \\\  \\ \  \\\  \
       \ \__\    \ \__\\ \__\ \__\\ \________\\ \_______\\ \_______\
        \|__|     \|__| \|__|\|__| \|________| \|_______| \|_______|

    `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Run(apiCmd, []string{})
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
