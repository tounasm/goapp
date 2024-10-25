package httpsrv

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_New(t *testing.T) {

	strChan := make(chan string)
	s := New(strChan)

	assert.Nil(t, s.server, "Expected server to be nil initially")
	assert.Empty(t, s.watchers, "Expected watchers to be empty initially")
	assert.NotNil(t, s.watchersLock, "Expected watchersLock to be initialized")
	assert.Empty(t, s.sessionStats, "Expected sessionStats to be empty initially")
}
func TestServer_Start(t *testing.T) {

	strChan := make(chan string)
	s := New(strChan)

	err := s.Start()
	assert.NoError(t, err, "Expected no error when starting the server")

	// Allow time for the server to start
	time.Sleep(1 * time.Second)

	testCases := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "Health Endpoint",
			url:            "http://localhost:8080/goapp/health",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "WebSocket Endpoint",
			url:            "http://localhost:8080/goapp/ws",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Home Endpoint",
			url:            "http://localhost:8080/goapp",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(tc.url)
			assert.NoError(t, err, "Failed to send request to %s", tc.url)
			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "Expected status code %d, but got %d", tc.expectedStatus, resp.StatusCode)
		})
	}

	s.Stop()
}
