package httpsrv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_IncStats(t *testing.T) {
	strChan := make(chan string, 100)
	s := New(strChan)

	testCases := []struct {
		name           string
		sessionID      string
		expectedSent   int
		expectedLength int
		increments     int
	}{
		{
			name:           "Session1 increment 5 times",
			sessionID:      "session1",
			expectedSent:   5,
			expectedLength: 1,
			increments:     5,
		},
		{
			name:           "Session2 increment 2 times",
			sessionID:      "session2",
			expectedSent:   2,
			expectedLength: 2,
			increments:     2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			for i := 0; i < tc.increments; i++ {
				s.incStats(tc.sessionID)
			}

			assert.Equal(t, tc.expectedLength, len(s.sessionStats), "Expected sessionStats length to be %d, but got %d", tc.expectedLength, len(s.sessionStats))

			// Find the correct session by sessionID and verify the 'sent' count
			found := false
			for _, stat := range s.sessionStats {
				if stat.id == tc.sessionID {
					found = true
					assert.Equal(t, tc.expectedSent, stat.sent, "Expected sent count for session %s to be %d, but got %d", tc.sessionID, tc.expectedSent, stat.sent)
					break
				}
			}

			// Assert that the session was found
			assert.True(t, found, "Expected session with ID %s to be present, but it wasn't", tc.sessionID)
		})
	}
}

func TestServer_IncrementExistingSessionStats(t *testing.T) {

	strChan := make(chan string)
	server := New(strChan)
	server.sessionStats = []sessionStats{
		{id: "session1", sent: 5},
		{id: "session2", sent: 10},
	}

	testCases := []struct {
		name         string
		id           string
		sessionFound bool
	}{
		{
			name:         "Existing session ID",
			id:           "session1",
			sessionFound: true,
		},
		{
			name:         "Non-existing session ID",
			id:           "session3",
			sessionFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := server.incrementExistingSessionStats(tc.id)
			assert.Equal(t, tc.sessionFound, result)
		})
	}
}

func TestServer_RemoveStats(t *testing.T) {

	strChan := make(chan string)
	server := New(strChan)

	server.sessionStats = []sessionStats{{id: "session1", sent: 1}, {id: "session2", sent: 2}, {id: "session3", sent: 3}}

	tests := []struct {
		name          string
		indexToRemove int
		expectedStats []sessionStats
	}{
		{
			name:          "Remove middle element",
			indexToRemove: 1,
			expectedStats: []sessionStats{{id: "session1", sent: 1}, {id: "session3", sent: 3}},
		},
		{
			name:          "Remove first element",
			indexToRemove: 0,
			expectedStats: []sessionStats{{id: "session3", sent: 3}},
		},
		{
			name:          "No change when index is out of range",
			indexToRemove: 5,                                         // Out of range
			expectedStats: []sessionStats{{id: "session3", sent: 3}}, // Expect no change
		},
		{
			name:          "Remove last element",
			indexToRemove: 0,
			expectedStats: []sessionStats{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server.removeStats(tt.indexToRemove)
			assert.Equal(t, tt.expectedStats, server.sessionStats, "Expected sessionStats to match the expected state after removal")
		})
	}
}
