package aigenerate_hdl

import (
	"bufio"
	"easyaichat/internal/utils/ollama"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type AiGenerateHandler struct {
}

func (h *AiGenerateHandler) Init(root fiber.Router) {
	root.Post("/generate", h.AiGenerate)
	root.Post("/chat", h.ChatGenerate)
}

type AiGenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type AiChatGenerateRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

func (h *AiGenerateHandler) AiGenerate(c *fiber.Ctx) error {

	var req AiGenerateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	ch := make(chan ollama.GenerateCompletionResponse)
	input := ollama.GenerateACompletionRequest{
		Model:  req.Model,
		Prompt: req.Prompt,
	}

	go ollama.GetClient().GenerateACompletion(input).Streaming(ch)
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for response := range ch {
			if err := json.NewEncoder(w).Encode(response); err != nil {
				return
			}
			w.Flush()
		}
	})

	return nil
}

func (h *AiGenerateHandler) ChatGenerate(c *fiber.Ctx) error {
	var req AiChatGenerateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var ChatMessages []ollama.ChatMessage

	for _, message := range req.Messages {
		chatMessage := ollama.ChatMessage{
			Role:    message.Role,
			Content: message.Content,
		}
		ChatMessages = append(ChatMessages, chatMessage)
	}
	input := ollama.GenerateAChatCompletionRequest{
		Model:    req.Model,
		Messages: ChatMessages,
	}

	ch := make(chan ollama.GenerateChatCompletionResponse)
	go ollama.GetClient().GenerateAChatCompletion(input).Streaming(ch)
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for response := range ch {
			if err := json.NewEncoder(w).Encode(response); err != nil {
				return
			}
			w.Flush()
		}
	})

	return nil
}
