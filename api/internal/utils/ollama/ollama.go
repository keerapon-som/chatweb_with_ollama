package ollama

import "time"

type GenerateACompletionRequest struct {
	Model     string    `json:"model"`      // required
	Prompt    string    `json:"prompt"`     // optional
	Suffix    string    `json:"suffix"`     // optional
	Images    *[]string `json:"images"`     // optional, base64-encoded images
	Format    *string   `json:"format"`     // optional, can be json or a JSON schema
	Options   *string   `json:"options"`    // optional, additional model parameters
	System    *string   `json:"system"`     // optional, overrides what is defined in the Modelfile
	Template  *string   `json:"template"`   // optional, overrides what is defined in the Modelfile
	Stream    *bool     `json:"stream"`     // optional, if false the response will be returned as a single response object
	Raw       *bool     `json:"raw"`        // optional, if true no formatting will be applied to the prompt
	KeepAlive *string   `json:"keep_alive"` // optional, controls how long the model will stay loaded into memory
	Context   *string   `json:"context"`    // deprecated, the context parameter returned from a previous request
}

type GenerateCompletionResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

type GenerateCompletionType interface {
	Streaming(progress chan<- GenerateCompletionResponse)
	NonStreaming() (*GenerateCompletionResponse, error)
}

type GenerateAChatCompletionRequest struct {
	Model     string             `json:"model"`
	Messages  []ChatMessage      `json:"messages"`
	Tools     []string           `json:"tools,omitempty"`
	Format    *string            `json:"format,omitempty"`
	Options   *map[string]string `json:"options,omitempty"`
	Stream    *bool              `json:"stream,omitempty"`
	KeepAlive *string            `json:"keep_alive,omitempty"`
}

type ChatMessage struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Images    *[]string `json:"images,omitempty"`
	ToolCalls *[]string `json:"tool_calls,omitempty"`
}

type GenerateChatCompletionType interface {
	Streaming(progress chan<- GenerateChatCompletionResponse)
	NonStreaming() (*GenerateChatCompletionResponse, error)
}

type ListLocalModelsResponse struct {
	Models []struct {
		Name       string `json:"name"`
		ModifiedAt string `json:"modified_at"`
		Size       int64  `json:"size"`
		Digest     string `json:"digest"`
		Details    struct {
			Format            string      `json:"format"`
			Family            string      `json:"family"`
			Families          interface{} `json:"families"`
			ParameterSize     string      `json:"parameter_size"`
			QuantizationLevel string      `json:"quantization_level"`
		} `json:"details"`
	} `json:"models"`
}

type PullaModelRequest struct {
	Model    string `json:"model"`
	Insecure bool   `json:"insecure"`
	Stream   bool   `json:"stream"`
}

type PullModelOpenStep1Response struct {
	Status string `json:"status"`
}

type PullModelOpenStreamResponse struct {
	Status    string `json:"status"`
	Digest    string `json:"digest"`
	Total     int    `json:"total"`
	Completed int    `json:"completed"`
}

type ListRunningModelsResponse struct {
	Models []struct {
		Name    string `json:"name"`
		Model   string `json:"model"`
		Size    int64  `json:"size"`
		Digest  string `json:"digest"`
		Details struct {
			ParentModel       string   `json:"parent_model"`
			Format            string   `json:"format"`
			Family            string   `json:"family"`
			Families          []string `json:"families"`
			ParameterSize     string   `json:"parameter_size"`
			QuantizationLevel string   `json:"quantization_level"`
		} `json:"details"`
		ExpiresAt string `json:"expires_at"`
		SizeVram  int64  `json:"size_vram"`
	} `json:"models"`
}

type PullModelRequestStep1 interface {
	Open() PullModelOpenStep1Response
	OpenStream(progress chan<- PullModelOpenStreamResponse)
}

type ShowModelInformationRequest struct {
	Model   string `json:"model"`
	Verbose bool   `json:"verbose"`
}

type DeleteAModelRequest struct {
	Model string `json:"model"`
}

type VersionResp struct {
	Version string `json:"version"`
}

type GenerateChatCompletionResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   *struct {
		Role    string       `json:"role"`
		Content string       `json:"content"`
		Images  *interface{} `json:"images"`
	} `json:"message"`
	Done               bool   `json:"done"`
	TotalDuration      *int64 `json:"total_duration"`
	LoadDuration       *int   `json:"load_duration"`
	PromptEvalCount    *int   `json:"prompt_eval_count"`
	PromptEvalDuration *int   `json:"prompt_eval_duration"`
	EvalCount          *int   `json:"eval_count"`
	EvalDuration       *int64 `json:"eval_duration"`
}

type Client interface {
	GenerateACompletion(input GenerateACompletionRequest) GenerateCompletionType
	GenerateAChatCompletion(input GenerateAChatCompletionRequest) GenerateChatCompletionType
	ListLocalModels() (ListLocalModelsResponse, error)
	ShowModelInformation(ShowModelInformationRequest) ([]byte, error)
	DeleteAModel(input DeleteAModelRequest) error
	PullaModel(input PullaModelRequest) PullModelRequestStep1
	ListRunningModels() (ListRunningModelsResponse, error)
	Version() (VersionResp, error)
}
