package models

import (
	"errors"
)

type Vote struct {
	ID       uint64 `json:"id,omitempty"` // omit empty faz com que o campo não seja passado para o JSON caso esteja vazio
	UserId   uint64 `json:"userId,omitempty"`
	CryptoId uint64 `json:"cryptoId,omitempty"`
}

type Ranking struct {
	CryptoName  string `json:"name,omitempty"`
	CryptoVotes uint64 `json:"votes,omitempty"`
}

func (vote *Vote) Prepare() error {
	if err := vote.validate(); err != nil {
		return err
	}

	if err := vote.format(); err != nil {
		return err
	}

	return nil
}

func (vote *Vote) validate() error {
	if vote.UserId <= 0 {
		return errors.New("userId is required")
	}
	if vote.CryptoId <= 0 {
		return errors.New("cryptoId is required")
	}
	return nil
}

func (vote *Vote) format() error {
	vote.UserId = uint64(vote.UserId)
	vote.CryptoId = uint64(vote.CryptoId)

	return nil
}
