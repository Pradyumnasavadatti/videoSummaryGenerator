package cmd

import (
	"vidSummary/internals"

	"github.com/spf13/cobra"
)

var audioCmd = &cobra.Command{
	Use:"audio",
	Short:"Used to get summary of given audio file",
	Long:"Generates the summaury of the local audio file",

	Run: func(cmd *cobra.Command, args []string){
		err := internals.ComposeSummary(vidLocation,summaryName)
		if err!=nil{
			panic(err)
		}
	},
}

func init(){
	rootCmd.AddCommand(audioCmd)
}