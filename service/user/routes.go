package user

import (
	"fmt"
	"net/http"

	"github.com/Luan-Pereira66/api-ecommerce/service/auth"
	"github.com/Luan-Pereira66/api-ecommerce/types"
	"github.com/Luan-Pereira66/api-ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}

}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/regster", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//obter carga útil JSON
	var payload types.RegisterUSerPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// verifique se o usuário existe
	_, err := h.store.GetUSerByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("O usuário %s ja existe", payload.Email))
		return
	}

	hashedSenha, err := auth.HashPassword(payload.Senha)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// se não, criamos o novo
	err = h.store.CreateUser(types.User{
		Nome:      payload.Nome,
		Sobrenome: payload.Sobrenome,
		Email:     payload.Email,
		Senha:     hashedSenha,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
