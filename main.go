package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// define a porta
const PORT string = ":8080"

// mensagem para enviar como resposta json
type Message struct {
	Msg string
}

// resposta no formato json
func jsonMessageByte(msg string) []byte {
	errrMessage := Message{msg}
	byteContent, _ := json.Marshal(errrMessage)
	return byteContent
}

// função principal começa aqui
func main() {

	// http://localhost:8080
	http.HandleFunc("/", handleGetCars)

	// http://localhost:8080/car?id=1
	http.HandleFunc("/car", handleGetCarById)

	// http://localhost:8080/add
	http.HandleFunc("/add", handleAddCar)

	// http://localhost:8080/update
	http.HandleFunc("/update", handleUpdateCar)

	// http://localhost:8080/delete?id=1
	http.HandleFunc("/delete", handleDeleteCar)

	fmt.Printf("App está ouvindo em %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	// para o app se houver algum erro ao iniciar o servidor
	if err != nil {
		log.Fatal(err)
	}
}

// função handleGetCars para obter todos os carros
func handleGetCars(w http.ResponseWriter, r *http.Request) {
	cars, err := getCars()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	byteContent, _ := json.Marshal(cars)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteContent)
}

// função handleGetCarById para obter carro pelo ID
func handleGetCarById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	car, err := getCarById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	byteContent, _ := json.Marshal(car)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteContent)
}

// função handleAddCar para adicionar carros
func handleAddCar(w http.ResponseWriter, r *http.Request) {
	var cars []Cars
	err := json.NewDecoder(r.Body).Decode(&cars)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	err = addCars(cars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonMessageByte("Carros adicionados com sucesso"))
}

// função handleUpdateCar para atualizar carro
func handleUpdateCar(w http.ResponseWriter, r *http.Request) {
	var car Cars
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	err = updateCar(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMessageByte("Carro atualizado com sucesso"))
}

// função handleDeleteCar para deletar carro
func handleDeleteCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := deleteCar(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonMessageByte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMessageByte("Carro deletado com sucesso"))
}
