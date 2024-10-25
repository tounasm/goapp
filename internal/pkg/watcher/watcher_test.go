package watcher

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWatcher(t *testing.T) {
	w := New()

	assert.NotEmpty(t, w.GetWatcherId(), "Expected non-empty watcher ID")
	uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{12}$`)
	assert.True(t, uuidRegex.MatchString(w.GetWatcherId()), "Expected watcher ID to be a valid UUID")

	w.Start()
	defer w.Stop()

	testStr := "test_hex_value"
	w.Send(testStr)

	select {
	case counter := <-w.Recv():
		assert.Equal(t, testStr, counter.Value, "Expected counter value to match")
		assert.Equal(t, 1, counter.Iteration, "Expected counter iteration to be 1")
	case <-time.After(2 * time.Second):
		t.Error("Timeout waiting for counter update")
	}

	w.ResetCounter()

	select {
	case counter := <-w.Recv():
		assert.Equal(t, 0, counter.Iteration, "Expected counter iteration to be 0 after reset")
		assert.Equal(t, "", counter.Value, "Expected counter value to be empty after reset")
	case <-time.After(2 * time.Second):
		t.Error("Timeout waiting for counter reset")
	}
}
