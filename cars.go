package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cars struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Brand   string `json:"brand"`
	Mileage string `json:"mileage"`
	Price   string `json:"price"`
}

// Função para obter todos os carros do arquivo cars.json
func getCars() ([]Cars, error) {
	cars := []Cars{}
	carsByte, err := os.ReadFile("cars.json")
	if err != nil {
		return cars, err
	}
	err = json.Unmarshal(carsByte, &cars)
	if err != nil {
		return cars, err
	}
	return cars, nil
}

// Função para obter um carro pelo ID
func getCarById(id string) (Cars, error) {
	cars, err := getCars()
	if err != nil {
		return Cars{}, err
	}
	for _, car := range cars {
		if car.Id == id {
			return car, nil
		}
	}
	return Cars{}, fmt.Errorf("carro com ID %s não encontrado", id)
}

// Função para salvar a lista de carros no arquivo cars.json
func saveCars(cars []Cars) error {
	carsByte, err := json.Marshal(cars)
	if err != nil {
		return err
	}
	err = os.WriteFile("cars.json", carsByte, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Função para adicionar novos carros à lista existente
func addCars(car []Cars) error {
	cars, err := getCars()
	if err != nil {
		return err
	}
	cars = append(cars, car...)
	err = saveCars(cars)
	if err != nil {
		return err
	}
	return nil
}

// Função para atualizar um carro existente
func updateCar(car Cars) error {
	cars, err := getCars()
	if err != nil {
		return err
	}
	for i, c := range cars {
		if c.Id == car.Id {
			cars[i] = car
			break
		}
	}
	err = saveCars(cars)
	if err != nil {
		return err
	}
	return nil
}

// Função para deletar um carro pelo ID
func deleteCar(id string) error {
	cars, err := getCars()
	if err != nil {
		return err
	}
	found := false
	for i, car := range cars {
		if car.Id == id {
			cars = append(cars[:i], cars[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("carro com ID %s não encontrado", id)
	}
	err = saveCars(cars)
	if err != nil {
		return err
	}
	return nil
}
