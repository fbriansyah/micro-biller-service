package chi

import (
	"errors"
	"net/http"
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
		adapter.errorJSON(w, err, http.StatusNotFound)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    bill,
	}

	adapter.writeJSON(w, http.StatusOK, payload)
}
