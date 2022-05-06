package gate

type Gate interface {
	Can() error
}
