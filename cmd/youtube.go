package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"vidSummary/internals"

	"github.com/spf13/cobra"
)

var youtubeCmd = &cobra.Command{
	Use: "youtube",
	Short: "To generate summary of youtube video",
	Long: "By given youtube video id or link, summary of the video will be generated",
	Run: func(cmd *cobra.Command, args []string) {
		ctx,cancel := context.WithCancel(context.Background())
		defer cancel()
		internals.Wg.Add(1)
		go internals.StartWithContext(ctx,&internals.Wg,"Getting youtube video")
		ytCmd := exec.Command("yt-dlp",
			"-x",                           
			"--audio-format", "m4a",        
			"--audio-quality", "0",         
			"-o", internals.AudioName,     
			vidLocation,
			"--quiet",
		)

		ytCmd.Stdout = os.Stdout
		ytCmd.Stderr = os.Stderr

		err := ytCmd.Run()
		if err != nil {
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
	rootCmd.AddCommand(youtubeCmd)
}