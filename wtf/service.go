package wtf

import (
	log "github.com/sirupsen/logrus" // Alias logrus as log
)

// Service represents an API service that tracks the
// WTF level of users. It's methods conform to the `net/rpc`
// standard such that they can be call directly via a JSON-RPC
// server.
type Service struct {
	// Maps username -> WTF level
	levels map[string]float64
}

// New creates a new Service object.
func New() *Service {
	return &Service{
		levels: make(map[string]float64), // allocate memory for the map
	}
}

// SetLevelRequest is used as the first input parameter to `Service.SetLevel()`.
type SetLevelRequest struct {
	Username string
	Level    float64
}

// SetLevel records a user's WTF level.
// SetLevel does not have a return value,
// so the second argument should always be nil
func (s *Service) SetLevel(args SetLevelRequest, _ *int) error {
	log.Printf("Setting user %s's WTF level to %v", args.Username, args.Level)
	s.levels[args.Username] = args.Level
	return nil
}

// RemoveUser deletes a user's WTF level.
func (s *Service) RemoveUser(username string, _ *int) error {
	log.Printf("Removing user %s's WTF level", username)
	delete(s.levels, username)
	return nil
}

// Avg calcuates the mean average of all users current WTF levels.
func (s *Service) Avg(_ *int, result *float64) error {
	log.Print("Calculating avg WTF level")
	avg := 0.0
	for _, level := range s.levels {
		avg += level
	}
	avg /= float64(len(s.levels))
	*result = avg
	return nil
}
