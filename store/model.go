package store

import "encoding/json"

// TODO, maybe binary formatted custom encoder and decoder for performance
// TODO, flexibility on value, slice of bytes instead of a fixed type

type CommandType string

const (
	SET    CommandType = "SET"
	GET    CommandType = "GET"
	DELETE CommandType = "DELETE"
)

type Command struct {
	CmdType CommandType `json:"type"`
	Key     string      `json:"key"`
	Value   string      `json:"value"`
}

func EncodeCommand(c Command) ([]byte, error) {
	return json.Marshal(c)
}

func decodeCommand(bytes []byte) (*Command, error) {
	cmd := Command{}
	err := json.Unmarshal(bytes, &cmd)
	if err != nil {
		return nil, err
	}
	return &cmd, nil
}
