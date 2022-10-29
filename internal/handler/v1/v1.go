package v1

import (
	"context"
	"io"

	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/domain"
)

var _ ServerInterface = &handler{}

type validate interface {
	Into(r io.Reader, dest any) error
	Translate(error) map[string]string
}

type walletService interface {
	AddFunds(ctx context.Context, userID, amount string) (domain.Wallet, error)
}

type handler struct {
	log      internal.Logger
	validate validate
	wallet   walletService
}

func NewHandler(l internal.Logger, v validate, w walletService) *handler {
	return &handler{
		log:      l,
		validate: v,
		wallet:   w,
	}
}
