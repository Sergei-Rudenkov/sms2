package provider

import (
	"time"
	"sync"
	"math/rand"
)

type Cache interface {
	// Set a key with value to the cache. Returns true if an item was
	// evicted.
	Set(key, value interface{}) bool

	// Get an item from the cache by key. Returns the value if it exists,
	// and a bool stating whether or not it existed.
	Get(key interface{}) (interface{}, bool)

	// Keys returns a slice of all the keys in the cache
	Keys() []interface{}

	// Len returns the number of items present in the cache
	Len() int

	// Cap returns the total number of items the cache can retain
	Cap() int

	// Purge removes all items from the cache
	Purge()

	// Del deletes an item from the cache by key. Returns if an item was
	// actually deleted.
	Del(key interface{}) bool
}

type entry struct {
	transactionID int64
	key           interface{}
	value         interface{}
}

type agileCache struct {
	cap     int
	items   map[interface{}]*entry
	lock    sync.RWMutex
	NoReset bool
}

func (c *agileCache) Set(key, value interface{}, ttl time.Duration) bool {
	var updated bool
	c.lock.Lock()
	defer c.lock.Unlock()

	transactionID := rand.Int63()
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.updateEntry(ent, value, transactionID)
		updated = true
	} else {
		c.insertEntry(key, value, transactionID)
		updated = false
	}

	go func() {
		time.Sleep(ttl)
		e := c.getEntry(key)
		if transactionID == e.transactionID {
			// delete the item from the map
			c.removeEntry(e)
		}
	}()
	return updated
}

func (c *agileCache) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if ent, ok := c.items[key]; ok {
		return ent.value, true
	}
	return nil, false
}

func (c *agileCache) getEntry (key interface{}) *entry {
	ent, _ := c.items[key]
	return ent
}

func (c *agileCache) insertEntry(key, value interface{}, tID int64) *entry {
	// must already have a write lock
	ent := &entry{
		key:   key,
		value: value,
		transactionID: tID,
	}
	c.items[key] = ent

	return ent
}

func (c *agileCache) updateEntry(e *entry, value interface{}, tID int64) {
	e.value = value
	e.transactionID = tID
}

func (c *agileCache) removeEntry(e *entry) {
	// must already have a write lock

	// delete the item from the map
	delete(c.items, e.key)
}
