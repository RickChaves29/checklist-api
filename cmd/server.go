package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RickChaves29/checklist-api/internal/config"
	"github.com/RickChaves29/checklist-api/internal/dtos"
	"github.com/RickChaves29/checklist-api/internal/repositories"
	"github.com/RickChaves29/checklist-api/internal/usecases"
	"github.com/gorilla/mux"
)

func init() {
	db, err := config.Connection()
	if err != nil {
		println("error in connection")
		log.Fatal(err.Error())
	}
	_, err = db.Exec("DROP TABLE IF EXISTS tasks; CREATE TABLE tasks ( id SERIAL PRIMARY KEY, title VARCHAR(150) NOT NULL, description TEXT, done BOOLEAN DEFAULT false);")
	if err != nil {
		println("error when try exec query")
		log.Fatal(err.Error())
	}
}

func main() {
	router := mux.NewRouter()
	db, err := config.Connection()
	repo := repositories.NewRepository(db)
	uc := usecases.TaskUsecase{TaskRepository: repo}
	if err != nil {
		log.Fatal(err.Error())
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/tasks", http.StatusMovedPermanently)
	})
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		var data dtos.CreateTaskDTO
		json.NewDecoder(r.Body).Decode(&data)
		err := uc.Save(data)
		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}).Methods("POST")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		tasks, err := uc.FindAll()
		json.NewEncoder(w).Encode(tasks)
		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}).Methods("GET")
	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		stringID := vars["id"]
		id, errToParse := strconv.Atoi(stringID)
		if errToParse != nil {
			log.Fatal(errToParse)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		task, err := uc.FindById(id)
		json.NewEncoder(w).Encode(&task)
		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}).Methods("GET")
	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		var data dtos.UpdateTaskDTO
		vars := mux.Vars(r)
		stringID := vars["id"]
		id, errToParse := strconv.Atoi(stringID)
		if errToParse != nil {
			log.Fatal(errToParse)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewDecoder(r.Body).Decode(&data)
		println(data.Done)
		err := uc.Update(id, data.Done)
		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}).Methods("PUT")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		stringID := vars["id"]
		id, errToParse := strconv.Atoi(stringID)
		if errToParse != nil {
			log.Fatal(errToParse.Error())
			http.Error(w, errToParse.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err := uc.Delete(id)

		if err != nil {
			log.Fatal(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}).Methods("DELETE")

	println("task save in database")
	println("server is run in http://localhost:3000/tasks")
	http.ListenAndServe(":3000", router)

}
