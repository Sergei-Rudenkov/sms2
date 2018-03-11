package unit

import (
	"testing"
	"time"
	"sms2/storage/provider/agile"
)

const ttl time.Duration = 2 // in seconds
const expires time.Duration = 3 * time.Second

func TestAgileSet(t *testing.T) {
	c := agile.New()
	c.Set("first", "42", ttl)
	value, success := c.Get("first")
	assertEqual(t, success, true, "Get operation didn't succeed")
	assertEqual(t, value, "42", "Get operation: value unequal to expected")
}

func TestAgileRemove(t *testing.T) {
	c := agile.New()
	const key string = "test"
	c.Set(key, "42", ttl)
	value, success := c.Get(key)
	assertEqual(t, success, true, "Get operation didn't succeed")
	assertEqual(t, value, "42", "Get operation: value unequal to expected")
	isRemoved := c.Del(key)
	_, getSuccess := c.Get(key)
	assertEqual(t, isRemoved, true, "Remove operation didn't succeed")
	assertEqual(t, getSuccess, false, "Remove operation didn't succeed")
}

func TestAgileSetExpiration(t *testing.T) {
	c := agile.New()
	c.Set("t", "42", ttl)
	// wait until ttl is over
	time.Sleep(expires)

	_, success := c.Get("t")
	assertEqual(t, success, false, "ttl didn't erase the record")
}

func TestAgileKeys(t *testing.T) {
	c := agile.New()
	c.Set("t1", "42", ttl)
	c.Set("t2", "42", ttl)

	keys := c.Keys()
	expected := make([]interface{}, 0, 2)
	expected = append(expected, "t1")
	expected = append(expected, "t2")
	assertEqual(t, len(keys), len(expected), "")
}
