package brato

import (
	"encoding/json"
	"io"
)

type Config struct {
	Instruments []*Instrument `json:"instruments"`
}

type Brato struct {
	*InstrumentsService
}

func New() *Brato {
	return &Brato{}
}

func Read(r io.Reader) (*Config, error) {
	var c Config
	err := json.NewDecoder(r).Decode(&c)
	return &c, err
}

func (b *Brato) Sync(c *Config) error {
	for _, i := range c.Instruments {
		if _, err := b.instrumentsService().Create(i); err != nil {
			return err
		}
	}

	return nil
}

func (b *Brato) instrumentsService() *InstrumentsService {
	if b.InstrumentsService == nil {
		return DefaultInstrumentsService
	}

	return b.InstrumentsService
}
