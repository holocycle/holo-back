package model

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

type IDGenerator interface {
	New() string
}

var DefaultIDGenerator IDGenerator = &RandomIDGenerator{}
var currentGenerator IDGenerator = DefaultIDGenerator

func GetIDGenerator() IDGenerator {
	return currentGenerator
}

func SetIDGenerator(g IDGenerator) {
	currentGenerator = g
}

type RandomIDGenerator struct {
}

func (g *RandomIDGenerator) New() string {
	uuid, _ := uuid.New().MarshalBinary()
	return base58.Encode(uuid)
}
