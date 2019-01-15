package session

import "sync"

type MemorySession struct {
	data    map[string]interface{}
	id      string
	r_wlock sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		id:   id,
		data: make(map[string]interface{}, 8),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.r_wlock.Lock()
	defer m.r_wlock.Unlock()

	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.r_wlock.RLock()
	defer m.r_wlock.RUnlock()

	value, ok := m.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.r_wlock.Lock()
	defer m.r_wlock.Unlock()

	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
