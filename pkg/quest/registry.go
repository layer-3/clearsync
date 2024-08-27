package quest

var registry = map[string]Handler{}

func RegisterHandler(key string, questID string, handler Handler) {
	registry[key+"_"+questID] = handler
}

func GetHandler(key string, questID string) (Handler, bool) {
	handler, exists := registry[key+"_"+questID]
	return handler, exists
}
