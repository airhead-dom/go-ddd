package mapper

import "gopkg.in/jeevatkm/go-model.v1"

type GoModelMapper struct {
}

func NewGoModelMapper() Mapper {
	return &GoModelMapper{}
}

func (g *GoModelMapper) Map(dest interface{}, src interface{}) []error {
	errs := model.Copy(dest, src)
	return errs
}
