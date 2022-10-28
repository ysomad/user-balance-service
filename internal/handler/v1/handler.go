package v1

import (
	"github.com/ysomad/avito-internship-task/internal"
)

var _ ServerInterface = &handler{}

type handler struct {
	log internal.Logger
}

func NewHandler(l internal.Logger) *handler {
	return &handler{
		log: l,
	}
}
