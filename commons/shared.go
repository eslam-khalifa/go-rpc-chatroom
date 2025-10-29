package commons

type Message struct {
	Name string
	Text string
}

type Reply struct {
	Messages []string
}

func GetServerAddress() string {
	return "0.0.0.0:9999"
}