//go:generate go run github.com/getstream/easyjson/easyjson -pkg -all
package stream_chat

import (
	"time"
)

// for interfaces type matching
var (
	_ StreamClient  = (*Client)(nil)
	_ StreamChannel = (*Channel)(nil)
)

type StreamClient interface {
	// app.go
	GetApp() (*AppResponseFields, error)
	UpdateAppSettings(settings *AppSettings) error

	// device.go
	AddDevice(device *Device) error
	DeleteDevice(userID string, deviceID string) error
	GetDevices(userId string) (devices []*Device, err error)

	// channel.go
	CreateChannel(chanType string, chanID string, userID string, data map[string]interface{}) (*Channel, error)

	// channel_type.go
	CreateChannelType(chType *ChannelType) (*ChannelType, error)
	DeleteChannelType(chType string) error
	GetChannelType(chanType string) (ct *ChannelType, err error)
	ListChannelTypes() (map[string]*ChannelType, error)
	UpdateChannelType(name string, options map[string]interface{}) error

	// client.go
	CreateToken(userID string, expire time.Time) ([]byte, error)

	// message.go
	DeleteMessage(msgID string) error
	GetMessage(msgID string) (*Message, error)
	MarkAllRead(userID string) error
	UpdateMessage(msg *Message, msgID string) (*Message, error)
	FlagMessage(msgID string) error
	UnflagMessage(msgID string) error

	// query.go
	QueryUsers(q *QueryOption, sort ...*SortOption) ([]*User, error)
	QueryChannels(q *QueryOption, sort ...*SortOption) ([]*Channel, error)

	// user.go
	BanUser(targetID string, userID string, options map[string]interface{}) error
	DeactivateUser(targetID string, options map[string]interface{}) error
	ReactivateUser(targetID string, options map[string]interface{}) error
	DeleteUser(targetID string, options map[string][]string) error
	ExportUser(targetID string, options map[string][]string) (user *User, err error)
	FlagUser(targetID string, options map[string]interface{}) error
	MuteUser(targetID string, userID string) error
	MuteUsers(targetIDs []string, userID string) error
	UnBanUser(targetID string, options map[string]string) error
	UnFlagUser(targetID string, options map[string]interface{}) error
	UnmuteUser(targetID string, userID string) error
	UnmuteUsers(targetIDs []string, userID string) error
	UpdateUser(user *User) (*User, error)
	UpdateUsers(users ...*User) (map[string]*User, error)
}

type StreamChannel interface {
	// channel.go
	AddMembers(userIDs ...string) error
	AddModerators(userIDs ...string) error
	BanUser(targetID string, userID string, options map[string]interface{}) error
	Delete() error
	DemoteModerators(userIDs ...string) error
	MarkRead(userID string, options map[string]interface{}) error
	RemoveMembers(userIDs ...string) error
	Truncate() error
	UnBanUser(targetID string, options map[string]string) error
	Update(options map[string]interface{}, message string) error

	// event.go
	SendEvent(event *Event, userID string) error

	// message.go
	SendMessage(message *Message, userID string) (*Message, error)
	GetReplies(parentID string, options map[string][]string) (replies []*Message, err error)

	// reaction.go
	DeleteReaction(messageID string, reactionType string, userID string) (*Message, error)
	GetReactions(messageID string, options map[string][]string) ([]*Reaction, error)
	SendReaction(reaction *Reaction, messageID string, userID string) (*Message, error)
}
