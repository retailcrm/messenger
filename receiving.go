package messenger

import "time"

// Receive is the format in which webhook events are sent.
type Receive struct {
	// Object should always be `page`. (I don't quite understand why)
	Object string `json:"object"`
	// Entry is all of the different messenger types which were
	// sent in this event.
	Entry []Entry `json:"entry"`
}

// Entry is a batch of events which were sent in this webhook trigger.
type Entry struct {
	// ID is the ID of the batch.
	ID int64 `json:"id,string"`
	// Time is when the batch was sent.
	Time int64 `json:"time"`
	// Messaging is the events that were sent in this Entry
	Messaging []MessageInfo `json:"messaging"`
}

// MessageInfo is an event that is fired by the webhook.
type MessageInfo struct {
	// Sender is who the event was sent from.
	Sender Sender `json:"sender"`
	// Recipient is who the event was sent to.
	Recipient Recipient `json:"recipient"`
	// Timestamp is the true time the event was triggered.
	Timestamp int64 `json:"timestamp"`
	// Message is the contents of a message if it is a MessageAction.
	// Nil if it is not a MessageAction.
	Message *Message `json:"message"`
	// Delivery is the contents of a message if it is a DeliveryAction.
	// Nil if it is not a DeliveryAction.
	Delivery *Delivery `json:"delivery"`
	// Reaction represents reaction to Instagram message.
	Reaction *IGMessageReaction `json:"reaction,omitempty"`

	PostBack *PostBack `json:"postback"`

	Read *Read `json:"read"`

	OptIn *OptIn `json:"optin"`

	ReferralMessage *ReferralMessage `json:"referral"`

	AccountLinking *AccountLinking `json:"account_linking"`
}

type OptIn struct {
	// Sender is the sender of the message
	Sender Sender `json:"-"`
	// Recipient is who the message was sent to.
	Recipient Recipient `json:"-"`
	// Time is when the message was sent.
	Time time.Time `json:"-"`
	// Ref is the reference as given
	Ref string `json:"ref"`
}

// ReferralMessage represents referral endpoint.
type ReferralMessage struct {
	*Referral

	// Sender is the sender of the message
	Sender Sender `json:"-"`
	// Recipient is who the message was sent to.
	Recipient Recipient `json:"-"`
	// Time is when the message was sent.
	Time time.Time `json:"-"`
}

// Referral represents referral info.
type Referral struct {
	// Data originally passed in the ref param
	Ref string `json:"ref"`
	// Source type
	Source string `json:"source"`
	// The identifier dor the referral
	Type string `json:"type"`
	// ID of the ad
	AdID string `json:"ad_id,omitempty"`
	// The data containing information about the CTM ad, the user initiated the thread from.
	AdsContextData AdsContextData `json:"ads_context_data,omitempty"`
	// URI of the site from which the message was sent to the Facebook chat plugin.
	RefererURI string `json:"referer_uri,omitempty"`
}

// AdsContextData represents data containing information about the CTM ad, the user initiated the thread from.
type AdsContextData struct {
	// Title of the Ad
	AdTitle string `json:"ad_title"`
	// Url of the image from the Ad the user is interested
	PhotoURL string `json:"photo_url,omitempty"`
	// Thumbnail url of the video from the ad
	VideoURL string `json:"video_url,omitempty"`
	// ID of the post
	PostID string `json:"post_id"`
	// Product ID from the Ad the user is interested
	ProductID string `json:"product_id,omitempty"`
}

// Sender is who the message was sent from.
type Sender struct {
	ID int64 `json:"id,string"`
}

// Recipient is who the message was sent to.
type Recipient struct {
	ID        int64  `json:"id,string,omitempty"`
	PostID    string `json:"post_id,omitempty"`
	CommentID string `json:"comment_id,omitempty"`
}

// Attachment is a file which used in a message.
type Attachment struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
	// Type is what type the message is. (image, video, audio or location)
	Type string `json:"type"`
	// Payload is the information for the file which was sent in the attachment.
	Payload Payload `json:"payload"`
}

// QuickReply is a file which used in a message.
type QuickReply struct {
	// ContentType is the type of reply
	ContentType string `json:"content_type,omitempty"`
	// Title is the reply title
	Title string `json:"title,omitempty"`
	// Payload is the  reply information
	Payload string `json:"payload"`
}

// Payload is the information on where an attachment is.
type Payload struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	// Coordinates is Lat/Long pair of location pin
	Coordinates  *Coordinates `json:"coordinates,omitempty"`
	TemplateType string       `json:"template_type,omitempty"`
	Buttons      []Button     `json:"buttons,omitempty"`
}

type Button struct {
	Type                string `json:"type,omitempty"`
	Title               string `json:"title,omitempty"`
	Payload             string `json:"payload,omitempty"`
	URL                 string `json:"url,omitempty"`
	WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
	FallbackURL         string `json:"fallback_url,omitempty"`
	WebviewShareButton  string `json:"webview_share_button,omitempty"`
}

// Coordinates is a pair of latitude and longitude.
type Coordinates struct {
	// Lat is latitude
	Lat float64 `json:"lat"`
	// Long is longitude
	Long float64 `json:"long"`
}
