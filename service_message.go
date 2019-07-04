package teamcitymsg

type ServiceMessage interface {
	MessageType() string
	Attributes() map[string]string
	Single() bool
	String() string
}
