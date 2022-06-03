package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/service"
)

type PasienHandler struct {
	custService service.PasienService
}

func NewPasienHandler(pasienService service.PasienService) *PasienHandler {
	return &PasienHandler{custService: pasienService}
}

func (s *PasienHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *PasienHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID := vars["id"]
	pasien, err := s.custService.GetPasien(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Get Pasien error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, pasien)
}

func (s *PasienHandler) Post(w http.ResponseWriter, r *http.Request) {

	var cust = &model.Pasien{}
	err := json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.CreatePasien(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Create pasien error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, cust)
}

func (s *PasienHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID := vars["id"]

	var cust = &model.Pasien{}
	err := json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.UpdatePasien(custID, cust)
	if err != nil {
		errMsg := fmt.Sprintf("Update pasien error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "pasien updated")
}

func (s *PasienHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID := vars["id"]

	err := s.custService.DeletePasien(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Delete pasien error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "pasien deleted")
}

func (s *PasienHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
