package user

import (
	"net/http"

	"github.com/Luan-Pereira66/api-ecommerce/types"
	"github.com/Luan-Pereira66/api-ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}

}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/regster", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUSerPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exists

	// if it doesnt we create the new user

}
