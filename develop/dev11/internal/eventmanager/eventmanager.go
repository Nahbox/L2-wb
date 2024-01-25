package eventmanager

import (
	"encoding/json"
	"fmt"
	"time"
)

// EventManager ...
type EventManager struct {
	Users map[int]map[int]time.Time `json:"users"`
}

// DateInfo ...
type DateInfo struct {
	EventID int
	Date    time.Time
}

// NewEventManager ...
func NewEventManager() *EventManager {
	return &EventManager{
		Users: make(map[int]map[int]time.Time),
	}
}

// Add ...
func (s *EventManager) Add(userID, eventID int, date string) error {
	if _, ok := s.Users[userID]; !ok {
		s.Users[userID] = make(map[int]time.Time)
	}

	if _, ok := s.Users[userID][eventID]; !ok {
		dateTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return fmt.Errorf("date error")
		}
		s.Users[userID][eventID] = dateTime
		return nil
	}
	return fmt.Errorf("event id not unique")
}

// Delete ...
func (s *EventManager) Delete(userID, eventID int) error {
	if _, ok := s.Users[userID]; !ok {
		return fmt.Errorf("user id not found")
	}
	if _, ok := s.Users[userID][eventID]; !ok {
		return fmt.Errorf("event id not found")
	}

	delete(s.Users[userID], eventID)

	return nil
}

// Update ...
func (s *EventManager) Update(userID, eventID int, date string) error {
	if _, ok := s.Users[userID]; !ok {
		return fmt.Errorf("user id not found")
	}
	if _, ok := s.Users[userID][eventID]; !ok {
		return fmt.Errorf("event id not found")
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("date error")
	}
	s.Users[userID][eventID] = dateTime

	return nil
}

// ForDay ...
func (s *EventManager) ForDay(userID, eventDay int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, _, d := v.Date()
		if d == eventDay {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}

// ForWeek ...
func (s *EventManager) ForWeek(userID, eventWeek int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, w := v.ISOWeek()
		if w == eventWeek {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}

// ForMonth ...
func (s *EventManager) ForMonth(userID, eventMonth int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, m, _ := v.Date()
		if int(m) == eventMonth {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
