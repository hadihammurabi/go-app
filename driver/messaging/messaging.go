package messaging

const (
	DriverRabbitMQ = "rabbitmq"
)

type Table map[string]any

func (t Table) Validate() error {
	return nil
}

type Message struct {
	Headers Table
	Tag     uint64
	Message []byte
}
