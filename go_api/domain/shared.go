package domain

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
    return nil
}
