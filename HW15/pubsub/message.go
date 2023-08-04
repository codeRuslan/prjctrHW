package pubsub

type Message struct {
	fileName string
	action   string
}

func NewMessage(fileName string, action string) (*Message) {
	return &Message{
		fileName: fileName,
		action:   action,
	}
}

func (m *Message) GetFileName() string {
	return m.fileName
}

func (m *Message) GetAction() string {
	return m.action
}