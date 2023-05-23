package gmap

import (
	"math/rand"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	m := New[string, int]()
	m.Store("foo", 42)
	value, loaded := m.Load("foo")
	if !loaded {
		t.Fatal("expected loaded=true, but got false")
	}
	if value != 42 {
		t.Fatalf("expected value=42, but got %d", value)
	}
	m.Range(func(key string, value int) bool {
		if key != "foo" {
			t.Fatalf(`expected key="foo", but got %s`, key)
		}
		if value != 42 {
			t.Fatalf(`expected value=42, but got %d`, value)
		}
		return true
	})
	m.Delete("foo")
	value, loaded = m.Load("foo")
	if loaded {
		t.Fatal("expected loaded=false, but got true")
	}
	if value != 0 {
		t.Fatalf("expected value=0, but got %d", value)
	}
}

var benchmarkSyncMapResult int64

func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		var res int64
		key := randKey()
		for pb.Next() {
			m.Store(key, rand.Int63())
			v, _ := m.Load(key)
			m.Delete(key)
			res = v.(int64)
		}
		benchmarkSyncMapResult = res
	})
}

var benchmarkMapResult int64

func BenchmarkMap(b *testing.B) {
	m := New[string, int64]()
	b.RunParallel(func(pb *testing.PB) {
		var res int64
		key := randKey()
		for pb.Next() {
			m.Store(key, rand.Int63())
			res, _ = m.Load(key)
			m.Delete(key)
		}
		benchmarkMapResult = res
	})
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randKey() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
