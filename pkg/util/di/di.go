package di

const (
	DI_CONFIG = iota
	DI_REPOSITORY
	DI_SERVICE
	DI_DELIVERY_REST
	DI_DELIVERY_MQ
)

type IOC map[int]interface{}
