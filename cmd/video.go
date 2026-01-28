package cmd

import (
	"context"
	"errors"
	"fmt"
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

		isValid := internals.CheckFileExtension(internals.VideoType,vidLocation)
		if !isValid{
			panic(errors.New("Not a valid video file"))
		}
		isValid, err := internals.CheckFileMimeType(internals.VideoType, vidLocation)
		if err!=nil || !isValid{
			panic(errors.New("Not a valid video file"))
		}
		fmt.Println("\u2713Validated the file");
		ctx,cancel := context.WithCancel(context.Background())
		defer cancel()
		internals.Wg.Add(1)
		go internals.StartWithContext(ctx,&internals.Wg,"Executing ffmpeg command to extract audio...")
		cmdToExec := exec.Command("ffmpeg","-hide_banner","-loglevel","quiet","-i",vidLocation,"-vn","-ac","1","-ar","16000","-c:a","libmp3lame","-b:a","64k","-y",internals.AudioName)

		cmdToExec.Stdin = os.Stdin
		cmdToExec.Stderr = os.Stderr

		err = cmdToExec.Run()

		if err!=nil{
			panic(err)
		}
		cancel()
		internals.Wg.Wait()
		fmt.Println("\u2713Audio file created")
		err = internals.ComposeSummary(internals.AudioName,internals.SummaryName)
		if err!=nil{
			panic(err)
		}
	},
}


func init(){
	rootCmd.AddCommand(videoCmd)
}