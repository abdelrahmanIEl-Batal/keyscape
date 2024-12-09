package store

import (
	"fmt"
	util "keyscape/utils"
	"sync"
)

type Store struct {
	data *sync.Map
}

func NewStore() *Store {
	return &Store{
		data: &sync.Map{},
	}
}

func (s *Store) put(key, value string) {
	s.data.Store(key, value)
}

func (s *Store) get(key string) ([]byte, error) {
	val, exists := s.data.Load(key)
	if !exists {
		return nil, ErrKeyNotFound
	}
	return []byte(val.(string)), nil
}

func (s *Store) delete(key string) {
	s.data.Delete(key)
}

func (s *Store) Apply(bytes []byte) ([]byte, error) {
	command, err := decodeCommand(bytes)
	if err != nil {
		util.DebugError(err, "Apply:decodeCommand")
		return nil, ErrInvalidCommand
	}
	switch command.CmdType {
	case GET:
		return s.get(command.Key)
	case SET:
		s.put(command.Key, command.Value)
		return nil, nil
	case DELETE:
		s.delete(command.Key)
	default:
		return nil, fmt.Errorf("invalid command, type: %s", command.CmdType)
	}
	return nil, nil
}
