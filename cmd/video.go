package cmd

import (
	"log"
	"os"
	"os/exec"
	"vidSummary/internals"

	"github.com/spf13/cobra"
)

var videoCmd = &cobra.Command{
	Use: "video",
	Short: "Summarize the video from the local",
	Long: "Takes video from the local and summarizes it",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Executing ffmpeg")
		cmdToExec := exec.Command("ffmpeg","-i",vidLocation,"-vn","-ac","1","-ar","16000","-c:a","libmp3lame","-b:a","64k","-y",audioName)

		cmdToExec.Stdin = os.Stdin
		cmdToExec.Stderr = os.Stderr

		err := cmdToExec.Run()

		if err!=nil{
			panic(err)
		}
		log.Println("Audio file created")
		err = internals.ComposeSummary(audioName,summaryName)
		if err!=nil{
			panic(err)
		}
	},
}


func init(){
	rootCmd.AddCommand(videoCmd)
}