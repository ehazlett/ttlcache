package ttlcache

import (
	"strconv"
	"testing"
	"time"
)

func TestSetSimple(t *testing.T) {
	c, err := NewTTLCache(time.Millisecond * 1000)
	if err != nil {
		t.Fatal(err)
	}

	k := "testkey"
	v := "testval"

	if err := c.Set(k, v); err != nil {
		t.Fatal(err)
	}
}

func TestSetGetSimple(t *testing.T) {
	c, err := NewTTLCache(time.Millisecond * 1000)
	if err != nil {
		t.Fatal(err)
	}

	k := "testkey"
	v := "testval"

	if err := c.Set(k, v); err != nil {
		t.Fatal(err)
	}

	r := c.Get(k)
	if r.(string) != v {
		t.Fatalf("expected value %s; received %s", v, r)
	}
}

func TestSetInvalidate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	ttl := time.Millisecond * 1000

	c, err := NewTTLCache(ttl)
	if err != nil {
		t.Fatal(err)
	}

	k := "testkey"
	v := "testval"

	if err := c.Set(k, v); err != nil {
		t.Fatal(err)
	}

	// wait to check
	time.Sleep(ttl - 250)

	r := c.Get(k)
	if r.(string) != v {
		t.Fatalf("expected value %s; received %s", v, r)
	}

	// wait to invalidate
	time.Sleep(ttl + 250)

	// confirm key is gone
	r = c.Get(k)
	if r != nil {
		t.Fatalf("expected nil value; received %s", r)
	}

}

func TestSetUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	ttl := time.Millisecond * 2000

	c, err := NewTTLCache(ttl)
	if err != nil {
		t.Fatal(err)
	}

	k := "testkey"
	v := "testval"

	if err := c.Set(k, v); err != nil {
		t.Fatal(err)
	}

	r := c.Get(k)
	if r.(string) != v {
		t.Fatalf("expected value %s; received %s", v, r)
	}

	// wait to invalidate
	time.Sleep(time.Millisecond * 1900)

	nv := "newval"

	if err := c.Set(k, nv); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1500)

	r = c.Get(k)
	if r.(string) != nv {
		t.Fatalf("expected value %s; received %s", nv, r)
	}

}

type Animal struct {
	name string
}

func TestIterator(t *testing.T) {
	ttl := time.Second * 2000

	c, err := NewTTLCache(ttl)
	if err != nil {
		t.Fatal(err)
	}

	// Insert 100 elements.
	for i := 0; i < 100; i++ {
		c.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}

	counter := 0
	// Iterate over elements.
	for item := range c.Iter() {
		val := item.Val

		if val == nil {
			t.Error("Expecting an object.")
		}
		counter++
	}

	if counter != 100 {
		t.Error("We should have counted 100 elements.")
	}
}

func TestBufferedIterator(t *testing.T) {
	ttl := time.Second * 2000

	c, err := NewTTLCache(ttl)
	if err != nil {
		t.Fatal(err)
	}

	// Insert 100 elements
	for i := 0; i < 100; i++ {
		c.Set(strconv.Itoa(i), Animal{strconv.Itoa(i)})
	}

	counter := 0
	// Iterate over elements
	for item := range c.IterBuffered() {
		val := item.Val

		if val == nil {
			t.Error("Expecting an object.")
		}
		counter++
	}

	if counter != 100 {
		t.Error("We should have counted 100 elements.")
	}
}
