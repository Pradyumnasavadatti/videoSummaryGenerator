package internals

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func GetSpeechToText(fileName string) (string, error){
	fmt.Println("Inside speech to text")
	filePtr,err := os.Open(fileName)
	if err!=nil{
		return "", errors.New("File not found")
	}
	defer filePtr.Close()
	ctx2, cancel2 := context.WithCancel(context.Background())	
	wg.Add(1)
	go StartWithContext(ctx2,&wg,"Calling Wisper for audio transcripts")
	defer cancel2()
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
		cancel2()
		wg.Wait()
		fmt.Println("\u2713Received the transcripts")
		return response.Text, nil
	}else{
		return "", errors.New("No api key found")
	}
}