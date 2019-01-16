package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

type RedisSession struct {
	sessionId  string
	pool       *redis.Pool
	sessionMap map[string]interface{}
	r_wlock    *sync.RWMutex
	flag       int
}

func (r *RedisSession) loadFromRedis() error {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return err
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), r.sessionMap)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSession) Set(key string, value interface{}) error {
	r.r_wlock.Lock()
	defer r.r_wlock.Unlock()
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
}

func (r *RedisSession) Get(key string) (interface{}, error) {
	r.r_wlock.RLock()
	defer r.r_wlock.RUnlock()

	if r.flag == SessionFlagNone {
		//该session还没有加载，那么要从redis中加载数据
		err := r.loadFromRedis()
		if err != nil {
			return nil, nil
		}
	}
	result, ok := r.sessionMap[key]
	if !ok {
		err := ErrSessionNotExist
		return nil, err
	}

	return result, nil

}

func (r *RedisSession) Del(key string) error {}

func (r *RedisSession) Save() error {}
