package ollama

var client Client
var configcli ClientConfig

func GetClient() Client {
	if client == nil {
		client = InitClient(&configcli)
	}

	return client
}

func InitClient(config *ClientConfig) Client {
	configcli = *config
	cli, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	return cli
}
