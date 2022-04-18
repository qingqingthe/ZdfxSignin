package forum

type Sign interface {
	Name() string

	BasicUrl() string

	Cookie() string

	Sign() (<-chan string, bool)
}
