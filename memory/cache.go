package mem

import (
	"context"
	"sync"
)

type Parameters struct {
	User       string                 `json:"user"`
	UserId     string                 `json:"user_id"`
	UrlId      string                 `json:"url_id"`
	Info       map[string]interface{} `json:"info"`
	ProtocolId string                 `json:"protocol_id"`
	Domain     string                 `json:"domain"`
	Protocol   string                 `json:"protocol"`
}

type Memory struct {
	f     Function
	cache map[string]FunctionResult
	lock  sync.Mutex
}

type Function func(string) ([]byte, error)

type FunctionResult struct {
	value []byte
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[string]FunctionResult),
	}
}

func (m *Memory) Get(key string) ([]byte, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()
	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func (m *Memory) getEndId(ctx context.Context, n string) ([]byte, error) {
	return endId(ctx, n)
}

func (m *Memory) getUrl(ctx context.Context, n string) ([]byte, error) {
	return url_w(ctx, n)
}

type getSomethingAdapter struct {
	m   *Memory
	ctx context.Context
	boo bool
}

func NewGetSomethingAdapter(ctx context.Context, b bool) *getSomethingAdapter {
	return &getSomethingAdapter{
		ctx: ctx,
		boo: b,
	}
}

func (gs *getSomethingAdapter) GetSomething(n string) ([]byte, error) {
	if gs.boo {
		return gs.m.getUrl(gs.ctx, n)
	}
	return gs.m.getEndId(gs.ctx, n)
}
