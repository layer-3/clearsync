package quest

var registry = map[string]Handler{}

func RegisterHandler(key string, handler Handler) {
	registry[key] = handler
}

func GetHandler(key string) (Handler, bool) {
	handler, exists := registry[key]
	return handler, exists
}
