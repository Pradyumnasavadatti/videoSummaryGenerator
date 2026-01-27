package internals

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func ComposeSummary(audioFile,summaryFile string)error{
		text, err := GetSpeechToText(audioFile)
		if err!=nil{
			return err
		}
		summary, err := CallOpenAi(text)
		err = os.WriteFile(summaryFile,[]byte(summary),0644)
		if err!=nil{
			return err
		}
		fmt.Printf("\u2713Successfully generated the summary in %v",summaryFile)
		return nil
}