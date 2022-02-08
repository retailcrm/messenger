package messenger

import (
	"encoding/json"
	"time"
)

// Message represents a Facebook messenger message.
type Message struct {
	// Sender is who the message was sent from.
	Sender Sender `json:"-"`
	// Recipient is who the message was sent to.
	Recipient Recipient `json:"-"`
	// Time is when the message was sent.
	Time time.Time `json:"-"`
	// Message is mine
	IsEcho bool `json:"is_echo,omitempty"`
	// Mid is the ID of the message.
	Metadata string `json:"metadata"`
	// Mid is the ID of the message.
	Mid string `json:"mid"`
	// Seq is order the message was sent in relation to other messages.
	Seq int `json:"seq"`
	// StickerID is the ID of the sticker user sent.
	StickerID int `json:"sticker_id"`
	// Text is the textual contents of the message.
	Text string `json:"text"`
	// Attachments is the information about the attachments which were sent
	// with the message.
	Attachments []Attachment `json:"attachments"`
	// Selected quick reply
	QuickReply *QuickReply `json:"quick_reply,omitempty"`
	// Entities for NLP
	// https://developers.facebook.com/docs/messenger-platform/built-in-nlp/
	NLP json.RawMessage `json:"nlp"`
	// Read Instagram message data to which this reply was sent to.
	Read *IGMessageRead `json:"read,omitempty"`
	// Reaction represents reaction to Instagram message.
	Reaction *IGMessageReaction `json:"reaction,omitempty"`
	// Referral with Instagram product data.
	Referral *IGMessageReferral `json:"referral,omitempty"`
	// IsUnsupported is being sent if Instagram message is not supported.
	IsUnsupported bool `json:"is_unsupported,omitempty"`
	// IsDeleted is being sent if message was deleted.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// ReplyTo the Instagram story or to the message.
	ReplyTo *IGReplyTo `json:"reply_to"`
}

// Delivery represents a the event fired when Facebook delivers a message to the
// recipient.
type Delivery struct {
	// Mids are the IDs of the messages which were read.
	Mids []string `json:"mids"`
	// RawWatermark is the timestamp of when the delivery was.
	RawWatermark int64 `json:"watermark"`
	// Seq is the sequence the message was sent in.
	Seq int `json:"seq"`
}

// Read represents a the event fired when a message is read by the
// recipient.
type Read struct {
	// RawWatermark is the timestamp before which all messages have been read
	// by the user
	RawWatermark int64 `json:"watermark"`
	// Seq is the sequence the message was sent in.
	Seq int `json:"seq"`
	// MID is id of message
	MID string `json:"mid"`
}

// IGMessageRead represents data with the read message ID. Present in the Instagram webhook.
type IGMessageRead struct {
	// Mid is a message ID.
	Mid string `json:"mid"`
}

// IGMessageReaction represents reaction to the Instagram message.
type IGMessageReaction struct {
	// Mid is a message ID.
	Mid string `json:"mid"`
	// Action can be {react|unreact}
	Action string `json:"action"`
	// Reaction is a reaction name. Optional.
	Reaction string `json:"reaction,omitempty"`
	// Emoji is optional.
	Emoji string `json:"emoji,omitempty"`
}

// IGMessageProduct represents Instagram product.
type IGMessageProduct struct {
	// ID of the product.
	ID string `json:"id,omitempty"`
}

// IGMessageReferral represents Instagram message referral with product ID.
type IGMessageReferral struct {
	// Product data.
	Product IGMessageProduct `json:"product,omitempty"`
}

// IGPostback represents Instagram postback webhook data.
type IGPostback struct {
	// Selected icebreaker question or title for the CTA (Generic Template)
	Title string `json:"title,omitempty"`
	// Payload is user defined payload.
	Payload string `json:"payload"`
}

// IGReplyTo represents data of the thing to what reply has been sent.
type IGReplyTo struct {
	// Mid is a message ID to which reply was sent.
	Mid string `json:"mid"`
	// Story data.
	Story *IGReplyToStory `json:"story,omitempty"`
}

// IGReplyToStory is a story data to which reply has been sent.
type IGReplyToStory struct {
	// URL of the story.
	URL string `json:"url,omitempty"`
	// ID of the story.
	ID string `json:"id,omitempty"`
}

// PostBack represents postback callback.
type PostBack struct {
	// Sender is who the message was sent from.
	Sender Sender `json:"-"`
	// Recipient is who the message was sent to.
	Recipient Recipient `json:"-"`
	// Time is when the message was sent.
	Time time.Time `json:"-"`
	// PostBack ID
	Payload string `json:"payload"`
	// Optional referral info
	Referral Referral `json:"referral"`
	// Title for the CTA that was clicked on
	Title string `json:"title"`
	// Message ID
	Mid string `json:"mid"`
}

type AccountLinking struct {
	// Sender is who the message was sent from.
	Sender Sender `json:"-"`
	// Recipient is who the message was sent to.
	Recipient Recipient `json:"-"`
	// Time is when the message was sent.
	Time time.Time `json:"-"`
	// Status represents the new account linking status.
	Status string `json:"status"`
	// AuthorizationCode is a pass-through code set during the linking process.
	AuthorizationCode string `json:"authorization_code"`
}

// Watermark is the RawWatermark timestamp rendered as a time.Time.
func (d Delivery) Watermark() time.Time {
	return time.Unix(d.RawWatermark/int64(time.Microsecond), 0)
}

// Watermark is the RawWatermark timestamp rendered as a time.Time.
func (r Read) Watermark() time.Time {
	return time.Unix(r.RawWatermark/int64(time.Microsecond), 0)
}

// GetNLP simply unmarshals the NLP entities to the given struct and returns
// an error if it's not possible.
func (m *Message) GetNLP(i interface{}) error {
	return json.Unmarshal(m.NLP, &i)
}
