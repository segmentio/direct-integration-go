package integration

import "time"

// Message contains fields common to all messaages.
// https://segment.com/docs/spec/common/
type Message struct {
	Type        string                 `json:"type"`
	UserID      string                 `json:"user_id"`
	AnonymousID string                 `json:"anonymous_id"`
	Context     map[string]interface{} `json:"context"`
	MessageID   string                 `json:"message_id"`
	ReceivedAt  time.Time              `json:"received_at"`
	SentAt      time.Time              `json:"sent_at"`
	Timestamp   time.Time              `json:"timestamp"`
	Version     string                 `json:"version"`
}

// TrackMessage contains fields for a track message.
// https://segment.com/docs/spec/track/
type TrackMessage struct {
	Event      string                 `json:"event"`
	Properties map[string]interface{} `json:"properties"`
	Message
}

// PageMessage contains fields for a page message.
// https://segment.com/docs/spec/page/
type PageMessage struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
	Message
}

// ScreenMessage contains fields for a screen message.
// https://segment.com/docs/spec/screen/
type ScreenMessage struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
	Message
}

// IdentifyMessage contains fields for an identify message.
// https://segment.com/docs/spec/identify/
type IdentifyMessage struct {
	Traits map[string]interface{} `json:"traits"`
	Message
}

// AliasMessage contains fields for an alias message.
// https://segment.com/docs/spec/alias/
type AliasMessage struct {
	PreviousID string `json:"previous_id"`
	Message
}

// GroupMessage contains fields for a group message.
// https://segment.com/docs/spec/group/
type GroupMessage struct {
	GroupID string                 `json:"group_id"`
	Traits  map[string]interface{} `json:"traits"`
	Message
}
