package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, cse := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			cache := NewCache(interval) // *Cache

			cache.Add(cse.key, cse.val)

			val, ok := cache.Get(cse.key)
			if !ok {
				t.Fatalf("expected to find key %q", cse.key)
			}
			if string(val) != string(cse.val) {
				t.Fatalf("expected %q, got %q", cse.val, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Fatalf("expected to find key before reap")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Fatalf("expected key to be reaped after interval")
	}
}
