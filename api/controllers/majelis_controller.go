package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"penyuluhan2/api/models"
	"penyuluhan2/api/responses"
	"penyuluhan2/api/utils/formaterror"

	"github.com/gorilla/mux"
)

func (server *Server) CreateMajelis(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	majelis := models.Majelis{}
	err = json.Unmarshal(body, &majelis)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	majelis.Prepare()
	majelisCreated, err := majelis.SaveMajelis(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, majelisCreated.ID))
	responses.JSON(w, http.StatusCreated, majelisCreated)
}

func (server *Server) GetMajeliss(w http.ResponseWriter, r *http.Request) {

	majelis := models.Majelis{}

	majeliss, err := majelis.FindAllMajelis(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, majeliss)
}

func (server *Server) GetMajelis(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	majelis := models.Majelis{}
	majelisGotten, err := majelis.FindMajelisByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, majelisGotten)
}

// func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("update")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, updatedUser)
// }

// func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	user := models.User{}

// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != 0 && tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	_, err = user.DeleteAUser(server.DB, uint32(uid))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
// 	responses.JSON(w, http.StatusNoContent, "")
// }
