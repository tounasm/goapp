package httpsrv

import "log"

type sessionStats struct {
	id   string
	sent int
}

func (w *sessionStats) print() {
	log.Printf("session %s has received %d messages\n", w.id, w.sent)
}

func (w *sessionStats) inc() {
	w.sent++
}

func (s *Server) incStats(id string) {
	// Find and increment.
	if !s.incrementExistingSessionStats(id) {
		// Not found, add new
		s.addNewSessionStats(id)
	}
}

func (s *Server) incrementExistingSessionStats(id string) bool {
	s.sessionLock.RLock()
	defer s.sessionLock.RUnlock()
	for i, ws := range s.sessionStats {
		if ws.id == id {
			s.sessionStats[i].inc()
			return true
		}
	}
	return false
}

func (s *Server) addNewSessionStats(id string) {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	s.sessionStats = append(s.sessionStats, sessionStats{id: id, sent: 1})
}

func (s *Server) removeStats(i int) {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	if i < 0 || i >= len(s.sessionStats) {
		log.Printf("Index %d is out of bounds, no element removed from sessionStats", i)
		return
	}
	s.sessionStats = append(s.sessionStats[:i], s.sessionStats[i+1:]...)
}
