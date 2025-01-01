package aigenerate

import "easyaichat/internal/utils/ollama"

type AIGenerateCompletionService interface {
	GenerateACompletion(prompt ollama.GenerateACompletionRequest) GeneralCompletion
	GenerateAChatCompletion(prompt ollama.GenerateAChatCompletionRequest) ChatCompletion
}

type ChatCompletion interface {
	Streaming(progress chan<- ollama.GenerateChatCompletionResponse)
	NonStreaming() (*ollama.GenerateChatCompletionResponse, error)
}

type GeneralCompletion interface {
	Streaming(progress chan<- ollama.GenerateCompletionResponse)
	NonStreaming() (*ollama.GenerateCompletionResponse, error)
}
