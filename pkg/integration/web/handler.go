package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/co0p/go-tls-watch/pkg/domain"
	"github.com/co0p/go-tls-watch/pkg/usecases"
)

type Handler struct {
	ValidateUsecase *usecases.ValidateUsecase
}

type WebsiteInfoResponse struct {
	Website string    `json:"website"`
	Valid   bool      `json:"valid"`
	Expires time.Time `json:"expires"`
}

func (h *Handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// for now, we only allow post request
		if r.Method != http.MethodPost {
			log.Println("method not allowed")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// get the payload
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed reading request", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		website := string(requestData)

		// now do the validation
		cert, err := h.ValidateUsecase.Validate(website)
		if err != nil {
			log.Println("failed validating website", err)
			cert = domain.Certificate{
				Origin: website,
			}
		}

		response := WebsiteInfoResponse{
			Website: cert.Origin,
			Valid:   cert.IsValid(),
			Expires: cert.NotAfter,
		}

		js, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}
