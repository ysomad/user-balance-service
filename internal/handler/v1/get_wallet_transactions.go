package v1

import (
	"net/http"

	"github.com/google/uuid"
)

func (h *handler) GetWalletTransactions(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
}
