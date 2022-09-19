package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"steamsale/entity"
	"steamsale/service"
	"strconv"
)

type Handler struct {
	user service.UserService
	item service.ItemService
}

func newHandler(userService service.UserService, itemService service.ItemService) *Handler {
	return &Handler{
		user: userService,
		item: itemService,
	}
}

func (hdr *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	_, err = hdr.user.RegisterUser(user)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
	}

	response := fmt.Sprintf("User %s created successfully.", user.Nickname)

	sendResponse(w, http.StatusCreated, response)
}

func (hdr *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	var user entity.UserForFind

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	result, err := hdr.user.FindUser(user)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
	}

	sendResponse(w, result.ID, result.Nickname, result.Status)
}

func (hdr *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	err = hdr.user.EditUser(user)
}

func (hdr *Handler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	QID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(QID)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
	}

	err = hdr.user.RemoveUser(user, id)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
	}

	sendResponse(w, http.StatusOK, "Removed user successfully")
}

func (hdr *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	item, err = hdr.item.AddItem(item)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}

	sendResponse(w, http.StatusCreated, item)
}

func (hdr *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Item

	err := json.NewDecoder(r.Body).Decode(&item.ID)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
	}
}
