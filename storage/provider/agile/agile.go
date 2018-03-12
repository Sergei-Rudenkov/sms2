package agile

import (
	"time"
	"sync"
	"math/rand"
	"sms2/storage/provider"
)

//////////////////////////////////
//
// Agile - realisation based on groutines.
// Goroutine is starting right after Set operation is done and it's waiting until ttl time is over - then deleting the item.
// For reference see: Set method.
//
//////////////////////////////////

type entry struct {
	transactionID int64
	key           interface{}
	value         interface{}
}

type agileCache struct {
	items   map[interface{}]*entry
	lock    sync.RWMutex
}

// New creates a new Cache with cap entries
func New() provider.Cache {
	c := agileCache{}

	c.items = make(map[interface{}]*entry)
	return &c
}

func (c *agileCache) Keys() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keys := make([]interface{}, 0, len(c.items))
	for k, _ := range c.items {
		keys = append(keys, k)
	}
	return keys
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
		c.lock.Lock()
		defer c.lock.Unlock()
		e := c.getEntry(key)
		if e != nil && transactionID == e.transactionID {
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

func (c *agileCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.items)
}

func (c *agileCache) Cap() int {
	// There is no capacity limits for this type of Cache. Return MaxInt for consistency
	const MaxInt = 2147483647 // max int32
	return MaxInt
}

func (c *agileCache) Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items = make(map[interface{}]*entry)
}

func (c *agileCache) Del(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	if ent, ok := c.items[key]; ok {
		c.removeEntry(ent)
		return true
	}

	return false
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
	// must already have a write lock
	e.value = value
	e.transactionID = tID
}

func (c *agileCache) removeEntry(e *entry) {
	// must already have a write lock
	// delete the item from the map
	delete(c.items, e.key)
}
