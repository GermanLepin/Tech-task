package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/validate"
)

func UpBalance(w http.ResponseWriter, r *http.Request) {
	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	corectAmount := validate.AmountValidate(w, r, amount)
	if corectAmount < 0.01 {
		return
	}

	instance.UpBalanceDB(ctx, w, id, corectAmount)
	jsonenc.JSONUpBalance(w, id, corectAmount)
}
