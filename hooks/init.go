package hooks

var AppliedHooks = make(map[string][]Hook)

type Hook struct {
	ID      string `json:"id"`
	Scope   string `json:"scope"`
	Handler any    `json:"-"`
}

const (
	HookScopeProvider = "provider"
	HookScopeServer   = "server"
	HookScopeWebUI    = "webui"
)

func AddHook(scope string, id string, handler any) {
	if AppliedHooks[id] == nil {
		AppliedHooks[id] = []Hook{{id, scope, handler}}
	} else {
		AppliedHooks[id] = append(AppliedHooks[id], Hook{id, scope, handler})
	}
}

func AddHooks(scope string, id string, handlers ...any) {
	if AppliedHooks[id] == nil {
		AppliedHooks[id] = []Hook{}
	}

	for _, handler := range handlers {
		AppliedHooks[id] = append(AppliedHooks[id], Hook{id, scope, handler})
	}
}

func GetHookHandlers[HandlerT any](scope string, id string) []HandlerT {
	var handlers []HandlerT
	for _, handler := range AppliedHooks[id] {
		if handler.Scope == scope {
			handlers = append(handlers, handler.Handler.(HandlerT))
		}
	}
	return handlers
}
