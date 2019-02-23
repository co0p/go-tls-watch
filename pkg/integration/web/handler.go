package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/co0p/go-tls-watch/pkg/usecases"
)

type Handler struct {
	ValidateUsecase *usecases.ValidateUsecase
}

type WebsiteInfoResponse struct {
	Website string `json:"website"`
	Valid   bool   `json:"valid"`
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
		info, err := h.ValidateUsecase.Validate(website)
		if err != nil {
			log.Println("failed validating website", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := WebsiteInfoResponse{
			Website: website,
			Valid:   info.IsValid(),
		}

		log.Println("returning validationInfo", response)
		js, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}
