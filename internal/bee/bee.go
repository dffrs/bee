package internal

type Bee struct{}

func NewBee() *Bee {
	return &Bee{}
}

func (b *Bee) Run() error {
	return nil
}
