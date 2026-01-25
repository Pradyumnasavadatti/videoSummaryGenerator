package internals

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func GetSpeechToText(fileName string) (string, error){
	log.Println("Inside speech to text") 
	filePtr,err := os.Open(fileName)
	if err!=nil{
		return "", errors.New("File not found")
	}
	defer filePtr.Close()
	filePtr.Seek(0, io.SeekStart)
	if apiKey, isPresent := os.LookupEnv("OPEN_AI_KEY"); isPresent{
		client := openai.NewClient(option.WithAPIKey(apiKey))
		response, err := client.Audio.Transcriptions.New(context.Background(),openai.AudioTranscriptionNewParams{
			File: openai.File(filePtr, fileName, "audio/wav"),
			Model: openai.AudioModelWhisper1,
		})
		if err!=nil{
			return "",err
		}
		log.Println("returning transcripts")
		return response.Text, nil
	}else{
		return "", errors.New("No api key found")
	}
}