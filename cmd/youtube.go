package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"vidSummary/internals"

	"github.com/kkdai/youtube/v2"
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
		client := youtube.Client{}

		vid, err := client.GetVideo(vidLocation);

		if err!=nil{
			panic(err)
		}
		formats := vid.Formats.Type("audio");

		if len(formats) == 0{
			panic("No audio found for English language")
		}

		best := formats[0];
		for _,format := range(formats){
			if best.Bitrate < format.Bitrate{
				best = format
			}
		}

		stream, _, err := client.GetStream(vid, &best)
		if err!=nil{
			panic(err)
		}
		defer stream.Close()

		cancel()
		internals.Wg.Wait()
		fmt.Println("\u2713Saving youtube audio file")
		fptr ,err := os.Create(internals.AudioName)
		if err!=nil{
			panic(err)
		}
		defer fptr.Close()

		_, err = io.Copy(fptr, stream)

		if err!=nil{
			panic(err)
		}
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