
package http

import "github.com/micro-plat/hydra/component"

type IReverse interface {
}

type Reverse struct {
	c component.IContainer
}

func NewReverse(c component.IContainer) *Reverse {
	return &Reverse{
		c: c,
	}
}
