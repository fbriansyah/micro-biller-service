package chi

import (
	"log"
	"net/http"

	"github.com/fbriansyah/micro-biller-service/internal/port"
)

type ChiAdapter struct {
	billerService port.BillerServicePort
}

type ChiAdapterConfig struct {
	ServerAddress string
}

func NewChiAdapter(billerService port.BillerServicePort) *ChiAdapter {

	return &ChiAdapter{
		billerService: billerService,
	}
}

func (adapter *ChiAdapter) Run(config ChiAdapterConfig) {
	srv := &http.Server{
		Addr:    config.ServerAddress,
		Handler: adapter.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
