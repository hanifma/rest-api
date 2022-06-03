package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/service"
)

type JenisPenyakitHandler struct {
	custService service.JenisPenyakitService
}

func NewJenisPenyakitHandler(jenisPenyakitService service.JenisPenyakitService) *JenisPenyakitHandler {
	return &JenisPenyakitHandler{custService: jenisPenyakitService}
}

func (s *JenisPenyakitHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *JenisPenyakitHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	jenisPenyakit, err := s.custService.GetJenisPenyakit(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Get JenisPenyakit error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, jenisPenyakit)
}

func (s *JenisPenyakitHandler) Post(w http.ResponseWriter, r *http.Request) {

	var cust = &model.JenisPenyakit{}
	err := json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.CreateJenisPenyakit(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Create jenisPenyakit error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "jenisPenyakit created")
}

func (s *JenisPenyakitHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	var cust = &model.JenisPenyakit{}
	err = json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.UpdateJenisPenyakit(custID, cust)
	if err != nil {
		errMsg := fmt.Sprintf("Update jenisPenyakit error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "jenisPenyakit updated")
}

func (s *JenisPenyakitHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.DeleteJenisPenyakit(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Delete jenisPenyakit error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "jenisPenyakit deleted")
}

func (s *JenisPenyakitHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
