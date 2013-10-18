package etcd

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	c := NewClient(nil)

	result, err := c.Set("foo", "bar", 100)

	if err != nil {
		t.Fatal(err)
	}

	if result.Key != "/foo" || result.Value != "bar" || result.TTL != 100 {
		t.Fatalf("Set 1 failed with %s %s %v", result.Key, result.Value, result.TTL)
	}

	time.Sleep(time.Second)

	result, err = c.Set("foo", "bar", 100)

	if err != nil {
		t.Fatal(err)
	}

	if result.Key != "/foo" || result.Value != "bar" || result.PrevValue != "bar" || result.TTL != 100 {
		t.Fatalf("Set 2 failed with %s %s %v", result.Key, result.Value, result.TTL)
	}
}