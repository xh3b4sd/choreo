package stream

type Connection interface {
	Close() error
	Read() ([]byte, error)
	Write([]byte) error
}
