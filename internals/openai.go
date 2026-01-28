package internals

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func CallOpenAi(transcriptedText string)(string,error){
	ctx, cancel := context.WithCancel(context.Background())	
	defer cancel()
	Wg.Add(1)
	go StartWithContext(ctx,&Wg,"Getting summary from openai")
	if apiKey,isPresent := os.LookupEnv("OPEN_AI_KEY"); isPresent{
		client := openai.NewClient(option.WithAPIKey(apiKey))
		prompt:= `
			Summarize the following text into:
			1. A concise paragraph overview
			2. Key points as bullet points

			Text:
			"""
			`+transcriptedText+`
			"""
		`
		res, err := client.Chat.Completions.New(context.Background(),openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
			Model:openai.ChatModelGPT3_5Turbo,
		})
		if err!=nil{
			panic(err)
		}
		cancel()
		Wg.Wait()
		fmt.Println("\u2713Generated the summary")
		return res.Choices[0].Message.Content, nil;
	}else{
		return "",errors.New("No Open AI api key found")
	}
}