package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	vidLocation string
	audioName string = "audio.wav"
	summaryName string = "summary.txt"
);


var rootCmd = &cobra.Command{
	Use: "vid",
	Short: "Video summarizer",
	Long: "This command will summarize the video which is powered by ai",

	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&vidLocation,"file", "f", "file location", "file location to get the video")
}