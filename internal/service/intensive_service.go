package service

import (
	"errors"
	"sort"
	"sync"
	"time"

	"intensive-bot/internal/domain"
)

var ErrIntensiveNotFound = errors.New("intensive not found")

type IntensiveService struct {
	mu         sync.RWMutex
	intensives map[int64]domain.Intensive
	nextID     int64
}

func NewIntensiveService() *IntensiveService {
	return &IntensiveService{intensives: make(map[int64]domain.Intensive), nextID: 1}
}

func (s *IntensiveService) ListOpen() ([]domain.Intensive, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]domain.Intensive, 0)
	for _, i := range s.intensives {
		if i.IsOpen {
			out = append(out, i)
		}
	}
	sort.Slice(out, func(a, b int) bool { return out[a].StartsAt.Before(out[b].StartsAt) })
	return out, nil
}

func (s *IntensiveService) ListAll() []domain.Intensive {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Intensive, 0, len(s.intensives))
	for _, i := range s.intensives {
		out = append(out, i)
	}
	sort.Slice(out, func(a, b int) bool { return out[a].ID < out[b].ID })
	return out
}

func (s *IntensiveService) GetByID(id int64) (domain.Intensive, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	i, ok := s.intensives[id]
	if !ok {
		return domain.Intensive{}, ErrIntensiveNotFound
	}
	return i, nil
}

func (s *IntensiveService) Create(i domain.Intensive) domain.Intensive {
	s.mu.Lock()
	defer s.mu.Unlock()
	i.ID = s.nextID
	s.nextID++
	if i.StartsAt.IsZero() {
		i.StartsAt = time.Now().Add(24 * time.Hour)
	}
	i.IsOpen = true
	s.intensives[i.ID] = i
	return i
}

func (s *IntensiveService) Update(id int64, i domain.Intensive) (domain.Intensive, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.intensives[id]
	if !ok {
		return domain.Intensive{}, ErrIntensiveNotFound
	}
	i.ID = id
	if i.StartsAt.IsZero() {
		i.StartsAt = cur.StartsAt
	}
	s.intensives[id] = i
	return i, nil
}

func (s *IntensiveService) SetOpen(id int64, isOpen bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	i, ok := s.intensives[id]
	if !ok {
		return ErrIntensiveNotFound
	}
	i.IsOpen = isOpen
	s.intensives[id] = i
	return nil
}
