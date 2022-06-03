package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/handlers"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
	"github.com/pushm0v/gorest/service"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	dokterRouter(api)
	pasienRouter(api)
	registrasiRouter(api)
	jenisPenyakitRouter(api)
	r.Use(LoggingMiddleware)
	return r
}

func dokterRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection()
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Dokter{})
	var dataRepo = repository.NewDokterRepository(dbConn)
	var dataService = service.NewDokterService(dataRepo)
	var dataHandler = handlers.NewDokterHandler(dataService)

	r.HandleFunc("/dokters/{id}", dataHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/dokters", dataHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/dokters/{id}", dataHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/dokters/{id}", dataHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", dataHandler.NotFound)
}

func pasienRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection()
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Pasien{})
	var dataRepo = repository.NewPasienRepository(dbConn)
	var dataService = service.NewPasienService(dataRepo)
	var dataHandler = handlers.NewPasienHandler(dataService)

	r.HandleFunc("/pasiens/{id}", dataHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/pasiens", dataHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/pasiens/{id}", dataHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/pasiens/{id}", dataHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", dataHandler.NotFound)
}

func registrasiRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection()
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.Registrasi{})
	var dataRepo = repository.NewRegistrasiRepository(dbConn)
	var dataService = service.NewRegistrasiService(dataRepo)
	var dataHandler = handlers.NewRegistrasiHandler(dataService)

	r.HandleFunc("/registrasi/{id}", dataHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/registrasi", dataHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/registrasi/{id}", dataHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/registrasi/{id}", dataHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", dataHandler.NotFound)
}

func jenisPenyakitRouter(r *mux.Router) {
	var dbConn, err = NewDBConnection()
	if err != nil {
		log.Fatalf("DB Connection error : %v", err)
	}
	dbConn.AutoMigrate(&model.JenisPenyakit{})
	var dataRepo = repository.NewJenisPenyakitRepository(dbConn)
	var dataService = service.NewJenisPenyakitService(dataRepo)
	var dataHandler = handlers.NewJenisPenyakitHandler(dataService)

	r.HandleFunc("/jenis-penyakit/{id}", dataHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/jenis-penyakit", dataHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/jenis-penyakit/{id}", dataHandler.Put).Methods(http.MethodPut)
	r.HandleFunc("/jenis-penyakit/{id}", dataHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/", dataHandler.NotFound)
}
