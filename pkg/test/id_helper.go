package test

import "github.com/holocycle/holo-back/pkg/model"

type CustomIDGenerator struct {
	List  []string
	Index int
}

func NewIDGenerator(id ...string) model.IDGenerator {
	return &CustomIDGenerator{
		List:  id,
		Index: 0,
	}
}

func (g *CustomIDGenerator) New() string {
	if g.Index >= len(g.List) {
		return "run-out-of-id"
	}
	id := g.List[g.Index]
	g.Index++
	return id
}
