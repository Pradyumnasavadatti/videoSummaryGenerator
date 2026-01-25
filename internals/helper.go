package internals

import (
	"log"
	"os"
)

func ComposeSummary(audioFile,summaryFile string)error{
	log.Println("Call whisper api to get text from audio file")
		text, err := GetSpeechToText(audioFile)
		if err!=nil{
			return err
		}
		summary, err := CallOpenAi(text)
		err = os.WriteFile(summaryFile,[]byte(summary),0644)
		if err!=nil{
			return err
		}
		log.Printf("Successfully generated the transcripted into %v the audio",summaryFile)
		return nil
}