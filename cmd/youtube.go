package cmd

import (
	"io"
	"log"
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
		log.Println("Getting youtube video")
		client := youtube.Client{}

		vid, err := client.GetVideo(vidLocation);

		if err!=nil{
			panic(err)
		}
		formats := vid.Formats.Type("audio").Language("English (US) original");

		if len(formats) == 0{
			panic("No audio found for English language")
		}

		stream, _, err := client.GetStream(vid, &formats[0])
		if err!=nil{
			panic(err)
		}
		defer stream.Close()

		log.Println("Saving youtube audio file")

		fptr ,err := os.Create(audioName)
		if err!=nil{
			panic(err)
		}
		defer fptr.Close()

		_, err = io.Copy(fptr, stream)

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
	rootCmd.AddCommand(youtubeCmd)
}