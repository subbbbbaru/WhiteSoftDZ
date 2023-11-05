package record

type DataLoader interface {
	FromJson([]byte) error
}
