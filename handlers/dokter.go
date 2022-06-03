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

type DokterHandler struct {
	custService service.DokterService
}

func NewDokterHandler(dokterService service.DokterService) *DokterHandler {
	return &DokterHandler{custService: dokterService}
}

func (s *DokterHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *DokterHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	dokter, err := s.custService.GetDokter(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Get Dokter error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, dokter)
}

func (s *DokterHandler) Post(w http.ResponseWriter, r *http.Request) {

	var cust = &model.Dokter{}
	err := json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.CreateDokter(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Create dokter error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, cust)
}

func (s *DokterHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	var cust = &model.Dokter{}
	err = json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.UpdateDokter(custID, cust)
	if err != nil {
		errMsg := fmt.Sprintf("Update dokter error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "dokter updated")
}

func (s *DokterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.DeleteDokter(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Delete dokter error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "dokter deleted")
}

func (s *DokterHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
