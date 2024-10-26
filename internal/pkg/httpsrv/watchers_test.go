package httpsrv

import (
	"testing"

	"goapp/internal/pkg/watcher"

	"github.com/stretchr/testify/assert"
)

func TestAddWatcher(t *testing.T) {
	strChan := make(chan string)
	server := New(strChan)
	w := &watcher.Watcher{}

	server.addWatcher(w)

	_, ok := server.watchers[w.GetWatcherId()]
	assert.True(t, ok, "Expected watcher to be added, but it was not")
}
func TestRemoveWatcher(t *testing.T) {
	strChan := make(chan string)
	server := New(strChan)
	w := &watcher.Watcher{}
	server.watchers = make(map[string]*watcher.Watcher)
	server.watchers[w.GetWatcherId()] = w

	server.removeWatcher(w)

	_, ok := server.watchers[w.GetWatcherId()]
	assert.False(t, ok, "Expected watcher to be removed, but it was still present")
}
