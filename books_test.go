package main

import (
	"os"
	"testing"
)

func TestAddBook(t *testing.T) {

	oldSize := getSize()

	book := Book{
		Id:       1000,
		Title:    "test",
		Finished: false,
	}

	addBook(&book)

	newSize := getSize()

	if newSize != oldSize+1 {
		t.Errorf("Tamaño incorrecto después de agregar un libro. Tamaño actual: %d, Tamaño esperado: %d", newSize, oldSize+1)
	}
}

func TestCreateFile(t *testing.T) {

	testFilename := "test.txt"

	saveInTxt(testFilename)

	if _, err := os.Stat(testFilename); os.IsNotExist(err) {
		t.Errorf("No se creo el archivo correctamente")
	}

	os.Remove(testFilename)
}
