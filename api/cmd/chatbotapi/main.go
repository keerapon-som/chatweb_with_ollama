package main

import (
	"easyaichat/cmd/chatbotapi/handlers"
	"easyaichat/internal/utils/ollama"
)

func main() {

	ollamaConfig := &ollama.ClientConfig{
		GatewayAddress: "localhost:11434",
		UseTls:         false,
	}

	ollama.InitClient(ollamaConfig)
	// inputz := ollama.GenerateAChatCompletionRequest{
	// 	Model: "llama3.2:1b",
	// 	Messages: []ollama.ChatMessage{
	// 		{
	// 			Role:    "user",
	// 			Content: "Hello, World!",
	// 		},
	// 	},
	// }
	// resp, err := ollama.GetClient().GenerateAChatCompletion(inputz).NonStreaming()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resp.Message)
	// ch := make(chan ollama.GenerateCompletionResponse)
	// promth := ollama.GenerateACompletionRequest{
	// 	Model:  "llama3.2:1b",
	// 	Prompt: "Hello, World!",
	// }
	// go aigenerate.NewAIGenerateService(ollama.GetClient()).GenerateACompletion(promth).Streaming(ch)
	// for x := range ch {
	// 	fmt.Println(x)
	// }

	app := handlers.CreateHandlers()

	app.Listen(":8080")

}
