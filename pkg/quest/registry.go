package quest

var registry = map[handlerKey]Handler{}

type handlerKey struct {
	key string
	id  string
}

func RegisterHandler(key string, questID string, handler Handler) {
	registry[handlerKey{key: key, id: questID}] = handler
}

func GetHandler(key string, questID string) (Handler, bool) {
	handler, exists := registry[handlerKey{key: key, id: questID}]
	return handler, exists
}
