package rabbitmq

type rabbitMQTestService struct{}

func InitTest() {
	rabbitMQ = &rabbitMQTestService{}
}

func (r *rabbitMQTestService) ExchangeDeclare(name, kind string) error {
	return nil
}

func (r *rabbitMQTestService) Publish(exchange, key string, message []byte) error {
	return nil
}

func (r *rabbitMQTestService) Close() {}
