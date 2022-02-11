package ttlcache

import "time"

func (t *TTLCache) Get(key string) *KV {
	t.lock.Lock()
	defer t.lock.Unlock()

	if k, ok := t.data[key]; ok {
		return &KV{
			Key:   key,
			Value: k.Value,
			TTL:   t.ttl - (time.Since(k.updated)),
		}
	}

	return nil
}

func (t *TTLCache) GetAll() []*KV {
	t.lock.Lock()
	defer t.lock.Unlock()

	keys := []*KV{}
	for k, v := range t.data {
		keys = append(keys, &KV{
			Key:   k,
			Value: v.Value,
			TTL:   t.ttl - (time.Since(v.updated)),
		})
	}
	return keys
}
