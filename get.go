package ttlcache

func (t *TTLCache) Get(key string) interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if k, ok := t.data[key]; ok {
		return k.Value
	}

	return nil
}

type Tuple struct {
	Key string
	Val interface{}
}

func (t *TTLCache) Iter() <-chan Tuple {
	ch := make(chan Tuple)
	go func() {
		t.lock.RLock()
		for key, val := range t.data {
			ch <- Tuple{key, val.Value}
		}
		t.lock.RUnlock()
		close(ch)
	}()
	return ch
}

func (t *TTLCache) IterBuffered() <-chan Tuple {
	ch := make(chan Tuple, len(t.data))
	go func() {
		t.lock.RLock()
		for key, val := range t.data {
			ch <- Tuple{key, val.Value}
		}
		t.lock.RUnlock()
		close(ch)
	}()
	return ch
}
