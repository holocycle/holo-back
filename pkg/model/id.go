package model

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

func NewID() string {
	uuid, _ := uuid.New().MarshalBinary()
	return base58.Encode(uuid)
}
