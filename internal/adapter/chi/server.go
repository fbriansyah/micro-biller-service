package chi

import (
	"log"
	"net/http"

	"github.com/fbriansyah/micro-biller-service/internal/port"
)

type ChiAdapter struct {
	billerService port.BillerServicePort
	serverAddress string
}

func NewChiAdapter(billerService port.BillerServicePort) *ChiAdapter {

	return &ChiAdapter{
		billerService: billerService,
	}
}

func (adapter *ChiAdapter) Run() {
	srv := &http.Server{
		Addr:    adapter.serverAddress,
		Handler: adapter.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
