package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-mqadmin",
	Short: "a tool to manage rocketmq written in go",
	Long:  "a tool to manage rocketmq written in go",
}

var nameSrvAddrs []string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&nameSrvAddrs, "nameSrvAddrs", "n", []string{}, "NameServer endpoint in format ip:port")
}
