package messageQueueCreator

import "encoding/json"

type MessageQueue struct {
	Body    interface{} `json:"body"`
	Pattern string      `json:"pattern"`
	Age     string      `json:"age"`
	Data    interface{} `json:"data"`
}

func NewMessageQueue(bodyM interface{}, pattern string, age string, data interface{}) *MessageQueue {
	return &MessageQueue{
		bodyM, pattern, age, data,
	}
}

func (m *MessageQueue) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}
	return bytes, err
}
