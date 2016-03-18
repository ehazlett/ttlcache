package ttlcache

func (t *TTLCache) Get(key string) interface{} {
    t.lock.RLock()
	defer t.lock.RUnlock()
	if k, ok := t.data[key]; ok {
		return k.Value
	}

	return nil
}
