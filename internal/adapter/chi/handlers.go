package chi

import (
	"errors"
	"net/http"

	"github.com/fbriansyah/micro-biller-service/internal/application"
)

func (adapter *ChiAdapter) Inquiry(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		BillNumber string `json:"bill_number"`
	}

	err := adapter.readJSON(w, r, &requestPayload)
	if err != nil {
		adapter.errorJSON(w, errors.New("invalid param"), http.StatusBadRequest)
		return
	}

	bill, err := adapter.billerService.Inquiry(requestPayload.BillNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			adapter.errorJSON(w, errors.New("cannot find bill number"), http.StatusNotFound)
			return
		}

		if err == application.ErrorBillAlreadyPaid {
			adapter.errorJSON(w, err, http.StatusNotFound)
			return
		}

		adapter.errorJSON(w, err, http.StatusInternalServerError)

		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    bill,
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}
