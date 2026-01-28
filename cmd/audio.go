package cmd

import (
	"errors"
	"fmt"
	"vidSummary/internals"

	"github.com/spf13/cobra"
)

var audioCmd = &cobra.Command{
	Use:"audio",
	Short:"Used to get summary of given audio file",
	Long:"Generates the summaury of the local audio file",

	Run: func(cmd *cobra.Command, args []string){
		isValid := internals.CheckFileExtension(internals.AudioType,vidLocation)
		if !isValid{
			panic(errors.New("Not a valid audio file"))
		}
		isValid, err := internals.CheckFileMimeType(internals.AudioType, vidLocation)
		if err!=nil || !isValid{
			panic(errors.New("Not a valid audio file"))
		}
		fmt.Println("\u2713Validated the file");
		err = internals.ComposeSummary(vidLocation,internals.SummaryName)
		if err!=nil{
			panic(err)
		}
	},
}

func init(){
	rootCmd.AddCommand(audioCmd)
}