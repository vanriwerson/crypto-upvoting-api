package models

import (
	"errors"
	"strings"
)

type Crypto struct {
	ID   uint64 `json:"id,omitempty"` // omit empty faz com que o campo não seja passado para o JSON caso esteja vazio
	Name string `json:"name,omitempty"`
}

func (crypto *Crypto) Prepare() error {
	if err := crypto.validate(); err != nil {
		return err
	}

	if err := crypto.format(); err != nil {
		return err
	}

	return nil
}

func (crypto *Crypto) validate() error {
	if crypto.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (crypto *Crypto) format() error {
	crypto.Name = strings.TrimSpace(crypto.Name)

	return nil
}
