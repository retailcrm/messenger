package messenger

// Action is used to determine what kind of message a webhook event is.
type Action int

const (
	// UnknownAction means that the event was not able to be classified.
	UnknownAction Action = iota - 1
	// TextAction means that the event was a text message (May contain attachments).
	TextAction
	// DeliveryAction means that the event was advising of a successful delivery to a
	// previous recipient.
	DeliveryAction
	// ReadAction means that the event was a previous recipient reading their respective
	// messages.
	ReadAction
	// PostBackAction represents post call back.
	PostBackAction
	// OptInAction represents opting in through the Send to Messenger button.
	OptInAction
	// ReferralAction represents ?ref parameter in m.me URLs.
	ReferralAction
	// AccountLinkingAction means that the event concerns changes in account linking
	// status.
	AccountLinkingAction
)

// SenderAction is used to send a specific action (event) to the Facebook.
// The result of sending said action is supposed to give more interactivity to the bot.
type SenderAction string

const (
	// MarkSeen marks message as seen.
	MarkSeen SenderAction = "MARK_SEEN"
	// TypingOn turns on "Bot is typing..." indicator.
	TypingOn SenderAction = "TYPING_ON"
	// TypingOff turns off typing indicator.
	TypingOff SenderAction = "TYPING_OFF"
	// React to the message.
	React SenderAction = "REACT"
	// Unreact to the message (remove reaction).
	Unreact SenderAction = "UNREACT"
)
