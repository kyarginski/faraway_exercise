package reader

type WisdomReader interface {
	ReadOne() (string, error)
}
