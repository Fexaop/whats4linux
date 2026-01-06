package query

type MessageType uint8

const (
	MessageTypeText MessageType = iota + 1
	MessageTypeImage
	MessageTypeVideo
	MessageTypeAudio
	MessageTypeDocument
	MessageTypeSticker
	MessageTypeContact
)

type MediaType uint8

const (
	MediaTypeNone MediaType = iota
	MediaTypeImage
	MediaTypeVideo
	MediaTypeAudio
	MediaTypeDocument
	MediaTypeSticker
)
