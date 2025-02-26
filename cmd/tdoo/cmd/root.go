package cmd

import (
	"frec.kr/tdoo/cmd/tdoo/cmd/run"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tdoo",
	Short: "CLI tool for managing Tdoo project",
	Long:  "A CLI tool for managing Tdoo, including generating Ent code, compiling Protobuf files, and more.",
}

func init() {
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.PersistentFlags().String("config", "./config.yaml", "config yaml file")
}

// Execute runs the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
