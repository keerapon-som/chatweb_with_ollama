package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClientImpl struct {
	GatewayAddress string
	TTlhttp        string
}

type ClientConfig struct {
	GatewayAddress string
	UseTls         bool
}

func NewClient(config *ClientConfig) (Client, error) {
	protocol := "http"
	if config.UseTls {
		protocol = "https"
	}

	// var response ListRunningModelsResponse
	resp, err := http.Get(protocol + "://" + config.GatewayAddress)
	if err != nil {
		return &ClientImpl{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return &ClientImpl{}, err
		}

		return &ClientImpl{}, fmt.Errorf("error connecting to gateway: %s with status code %s", string(bodyBytes), resp.Status)
	}

	return &ClientImpl{
		GatewayAddress: config.GatewayAddress,
		TTlhttp:        protocol,
	}, nil
}

type doGenerateCompletion struct {
	url string
	req GenerateACompletionRequest
}

func (c *ClientImpl) GenerateACompletion(input GenerateACompletionRequest) GenerateCompletionType {

	return &doGenerateCompletion{
		url: fmt.Sprintf("%s://%s/api/generate", c.TTlhttp, c.GatewayAddress),
		req: input,
	}
}

func (g *doGenerateCompletion) NonStreaming() (response *GenerateCompletionResponse, err error) {
	stream := false
	g.req.Stream = &stream
	requestBody, _ := json.Marshal(g.req)

	resp, err := http.Post(g.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decoder := json.NewDecoder(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (g *doGenerateCompletion) Streaming(progress chan<- GenerateCompletionResponse) {
	stream := true
	g.req.Stream = &stream
	requestBody, _ := json.Marshal(g.req)

	resp, err := http.Post(g.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		close(progress)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var response GenerateCompletionResponse

		err := decoder.Decode(&response)
		if err != nil {
			fmt.Println("Got error while decoding response")
			close(progress)
			return
		}
		progress <- response
	}
	close(progress)
}

type doGenerateChatCompletion struct {
	url string
	req GenerateAChatCompletionRequest
}

func (c *ClientImpl) GenerateAChatCompletion(input GenerateAChatCompletionRequest) GenerateChatCompletionType {

	return &doGenerateChatCompletion{
		url: fmt.Sprintf("%s://%s/api/chat", c.TTlhttp, c.GatewayAddress),
		req: input,
	}
}

func (g *doGenerateChatCompletion) NonStreaming() (response *GenerateChatCompletionResponse, err error) {
	stream := false
	g.req.Stream = &stream
	requestBody, _ := json.Marshal(g.req)

	resp, err := http.Post(g.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decoder := json.NewDecoder(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (g *doGenerateChatCompletion) Streaming(progress chan<- GenerateChatCompletionResponse) {
	stream := true
	g.req.Stream = &stream
	requestBody, _ := json.Marshal(g.req)

	resp, err := http.Post(g.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		close(progress)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var response GenerateChatCompletionResponse
		err := decoder.Decode(&response)
		if err != nil {
			fmt.Println("Got error while decoding response")
			close(progress)
			return
		}
		progress <- response
	}
	close(progress)
}

func (c *ClientImpl) ListRunningModels() (ListRunningModelsResponse, error) {
	var response ListRunningModelsResponse
	resp, err := http.Get(c.TTlhttp + "://" + c.GatewayAddress + "/api/ps")
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func (c *ClientImpl) ListLocalModels() (ListLocalModelsResponse, error) {
	var response ListLocalModelsResponse
	resp, err := http.Get(c.TTlhttp + "://" + c.GatewayAddress + "/api/tags")
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func (c *ClientImpl) DeleteAModel(input DeleteAModelRequest) error {
	requestBody, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", c.TTlhttp+"://"+c.GatewayAddress+"/api/delete", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting model %s", input.Model)
	}

	return nil
}

type PullModelImpl struct {
	Url        string
	InputModel PullaModelRequest
}

func (c *ClientImpl) PullaModel(input PullaModelRequest) PullModelRequestStep1 {

	return &PullModelImpl{
		Url:        fmt.Sprintf("%s://%s/api/pull", c.TTlhttp, c.GatewayAddress),
		InputModel: input,
	}
}

func (p *PullModelImpl) Open() PullModelOpenStep1Response {
	p.InputModel.Stream = false
	requestBody, _ := json.Marshal(p.InputModel)

	resp, err := http.Post(p.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return PullModelOpenStep1Response{}
	}
	defer resp.Body.Close()

	var response PullModelOpenStep1Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return PullModelOpenStep1Response{}
	}
	return response
}

func (p *PullModelImpl) OpenStream(progress chan<- PullModelOpenStreamResponse) {
	p.InputModel.Stream = true
	requestBody, _ := json.Marshal(p.InputModel)

	resp, err := http.Post(p.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		close(progress)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var response PullModelOpenStreamResponse
		err := decoder.Decode(&response)
		if err != nil {
			fmt.Println("Got error while decoding response")
			close(progress)
			return
		}
		progress <- response
	}
	close(progress)
}

func (c *ClientImpl) ShowModelInformation(input ShowModelInformationRequest) ([]byte, error) {

	requestBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.TTlhttp+"://"+c.GatewayAddress+"/api/show", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//print string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (c *ClientImpl) Version() (VersionResp, error) {
	var response VersionResp
	resp, err := http.Get(c.TTlhttp + "://" + c.GatewayAddress + "/api/version")
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
