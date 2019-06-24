package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	cache := NewCache(100)
	cache.Put("test", "abcde")
	val := cache.Get("test")

	if val != "abcde" {
		t.Fail()
	}
}

func TestGetNoKey(t *testing.T) {
	cache := NewCache(100)
	cache.Put("test", "abcde")
	val := cache.Get("test2")

	if val != "" {
		t.Fail()
	}
}

func TestSpaceRemaining(t *testing.T) {
	cache := NewCache(100)
	cache.Put("test", "aaaaaa")

	if cache.SizeRemaining() != 99 {
		t.Fail()
	}
}

func TestCacheEvictionNoUpdates(t *testing.T) {
	cache := NewCache(5)

	for i := 0; i < 10; i++ {
		cache.Put(string(i), string(i))
	}

	if cache.Order.Length != 5 {
		t.Errorf("Expected length to be 5")
	}

	for i := 0; i < 5; i++ {
		if _, exists := cache.Cache[string(i)]; exists {
			t.Errorf("Expected %d to be evicted", i)
		}
	}

	for i := 5; i < 10; i++ {
		if _, exists := cache.Cache[string(i)]; !exists {
			t.Errorf("Expected %d to be cached", i)
		}
	}
}

func TestLeastRecentlyUsed(t *testing.T) {
	cache := NewCache(3)

	cache.Put("one", "one")
	cache.Put("two", "two")
	cache.Put("three", "three")

	cache.Get("one")

	if cache.Order.Head.Key != "two" {
		t.Errorf("Expected least recently used to be two")
	}

	if cache.Order.RemoveTail().Key != "one" {
		t.Fail()
	}
}
