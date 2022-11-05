package websocket

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Channel struct {
	Key          string      `json:"key"`
	MemberIDs    []uuid.UUID `json:"memberIDs"`
	LastUpdate   time.Time   `json:"lastUpdate"`
	MessageCount int         `json:"messageCount"`
}

func newChannel(key string) *Channel {
	return &Channel{Key: key, MemberIDs: []uuid.UUID{}, LastUpdate: time.Now()}
}

func (s *Service) Join(connID uuid.UUID, ch string) (bool, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return false, invalidConnection(connID)
	}
	if !slices.Contains(conn.Channels, ch) {
		conn.Channels = append(conn.Channels, ch)
	}

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	created := false
	curr, ok := s.channels[ch]
	if !ok {
		curr = newChannel(ch)
		s.channels[ch] = curr
		created = true
	}
	if !slices.Contains(curr.MemberIDs, connID) {
		curr.MemberIDs = append(curr.MemberIDs, connID)
	}
	return created, nil
}

func (s *Service) Leave(connID uuid.UUID, ch string) (bool, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return false, invalidConnection(connID)
	}
	conn.Channels = chanWithout(conn.Channels, ch)

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = newChannel(ch)
	}
	filtered := make([]uuid.UUID, 0)
	for _, i := range curr.MemberIDs {
		if i != connID {
			filtered = append(filtered, i)
		}
	}

	if len(filtered) == 0 {
		delete(s.channels, ch)
		return true, nil
	}

	s.channels[ch].MemberIDs = filtered
	return false, s.sendOnlineUpdate(ch, conn.ID, conn.ID, false)
}

func (s *Service) GetChannel(ch string) *Channel {
	ret, _ := s.channels[ch]
	return ret
}

func (s *Service) GetConnection(id uuid.UUID) *Connection {
	ret, _ := s.connections[id]
	return ret
}

func chanWithout(c []string, ch string) []string {
	ret := make([]string, 0, len(c))
	for _, x := range c {
		if x != ch {
			ret = append(ret, x)
		}
	}
	return ret
}
