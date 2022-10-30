package models

import (
	"errors"
	"strconv"
	"strings"
)

type Vote struct {
	ID       uint64 `json:"id,omitempty"` // omit empty faz com que o campo não seja passado para o JSON caso esteja vazio
	UserId   string `json:"userId,omitempty"`
	CryptoId string `json:"cryptoId,omitempty"`
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
	parsedUser, err := strconv.ParseUint(vote.UserId, 10, 64)
	if err != nil {
		return errors.New("userId must be a number")
	}

	parsedCrypto, err := strconv.ParseUint(vote.CryptoId, 10, 64)
	if err != nil {
		return errors.New("cryptoId must be a number")
	}

	if vote.UserId == "" {
		return errors.New("userId is required")
	}

	if parsedUser <= 0 {
		return errors.New("userId must be a valid id")
	}

	if vote.CryptoId == "" {
		return errors.New("cryptoId is required")
	}

	if parsedCrypto <= 0 {
		return errors.New("cryptoId must be a valid id")
	}

	return nil
}

func (vote *Vote) format() error {
	vote.UserId = strings.TrimSpace(vote.UserId)
	vote.CryptoId = strings.TrimSpace(vote.CryptoId)

	return nil
}
