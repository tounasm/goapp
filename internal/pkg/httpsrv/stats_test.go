package httpsrv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_IncStats(t *testing.T) {
	s := Server{}

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
