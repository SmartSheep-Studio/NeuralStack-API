package hooks

import (
	"errors"
	"reflect"
	"sync"
)

type Hook interface {
	Subscribe(id string, handler any) error
	Dispatch(id string, args ...any)
}

type AsyncHook struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncHook() *AsyncHook {
	return &AsyncHook{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (h *AsyncHook) Subscribe(id string, handler any) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return errors.New("handler must be a function")
	}

	handler, ok := h.handlers[id]
	if !ok {
		h.handlers[id] = []reflect.Value{}
	}
	h.handlers[id] = append(h.handlers[id], v)

	return nil
}

func (h *AsyncHook) Dispatch(id string, args ...any) {
	handlers, ok := h.handlers[id]
	if !ok {
		// If handler list is empty, skip dispatch
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for _, handler := range handlers {
		go handler.Call(params)
	}
}

func (h *AsyncHook) ListAllEvents() []string {
	keys := make([]string, len(h.handlers))
	index := 0
	for k := range h.handlers {
		keys[index] = k
		index++
	}

	return keys
}
