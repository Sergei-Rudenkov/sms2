package unit

import (
	"testing"
	"sms2/storage/provider/fixed"
	"time"
)

func TestFixedSet(t *testing.T){
	c := fixed.New(10, fixed.WithTTL(ttl * time.Second))
	c.Set("ff", "42", 0)
	value, success := c.Get("ff")
	assertEqual(t, success, true, "Get operation didn't succeed")
	assertEqual(t, value, "42", "Get operation: value unequal to expected")
}

func TestFixedRemove(t *testing.T) {
	const key = "test"
	c := fixed.New(10, fixed.WithTTL(ttl * time.Second))
	c.Set(key, "42", 0)
	value, success := c.Get(key)
	assertEqual(t, success, true, "Get operation didn't succeed")
	assertEqual(t, value, "42", "Get operation: value unequal to expected")
	isRemoved := c.Del(key)
	_, getSuccess := c.Get(key)
	assertEqual(t, isRemoved, true, "Remove operation didn't succeed")
	assertEqual(t, getSuccess, false, "Remove operation didn't succeed")
}

func TestFixedSetExpiration(t *testing.T) {
	c := fixed.New(10, fixed.WithTTL(ttl * time.Second))
	c.Set("t", "42", 0)
	// wait until ttl is over
	time.Sleep(expires)

	_, success := c.Get("t")
	assertEqual(t, success, false, "ttl didn't erase the record")
}

func TestFixedKeys(t *testing.T) {
	c := fixed.New(10, fixed.WithTTL(ttl * time.Second))
	c.Set("t1", "42", 0)
	c.Set("t2", "42", 0)

	keys := c.Keys()
	expected := make([]interface{}, 0, 2)
	expected = append(expected, "t1")
	expected = append(expected, "t2")
	assertEqual(t, len(keys), len(expected), "")
}
