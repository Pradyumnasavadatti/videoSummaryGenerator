package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var vidLocation string
var rootCmd = &cobra.Command{
	Use: "vid",
	Short: "Video summarizer",
	Long: "This command will summarize the video which is powered by ai",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&vidLocation,"file", "f", "file location", "file location to get the video/audio")
}