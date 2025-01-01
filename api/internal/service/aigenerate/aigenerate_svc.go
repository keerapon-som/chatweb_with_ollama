package aigenerate

import "easyaichat/internal/utils/ollama"

type aiGeneration struct {
	ollamaClient ollama.Client
}

func NewAIGenerateService(ollamaClient ollama.Client) AIGenerateCompletionService {
	return &aiGeneration{
		ollamaClient: ollamaClient,
	}
}

type generalCompletionSetup struct {
	ollamaClient ollama.Client
	promth       ollama.GenerateACompletionRequest
}

func (a *aiGeneration) GenerateACompletion(prompt ollama.GenerateACompletionRequest) GeneralCompletion {
	return &generalCompletionSetup{
		ollamaClient: a.ollamaClient,
		promth:       prompt,
	}
}

func (o *generalCompletionSetup) Streaming(progress chan<- ollama.GenerateCompletionResponse) {
	o.ollamaClient.GenerateACompletion(o.promth).Streaming(progress)
}

func (o *generalCompletionSetup) NonStreaming() (resp *ollama.GenerateCompletionResponse, err error) {
	return o.ollamaClient.GenerateACompletion(o.promth).NonStreaming()
}

type chatCompletionSetup struct {
	ollamaClient ollama.Client
	promth       ollama.GenerateAChatCompletionRequest
}

func (a *aiGeneration) GenerateAChatCompletion(prompt ollama.GenerateAChatCompletionRequest) ChatCompletion {
	return &chatCompletionSetup{
		ollamaClient: a.ollamaClient,
		promth:       prompt,
	}
}

func (chat *chatCompletionSetup) Streaming(progress chan<- ollama.GenerateChatCompletionResponse) {
	chat.ollamaClient.GenerateAChatCompletion(chat.promth).Streaming(progress)
}

func (chat *chatCompletionSetup) NonStreaming() (resp *ollama.GenerateChatCompletionResponse, err error) {
	return chat.ollamaClient.GenerateAChatCompletion(chat.promth).NonStreaming()
}
