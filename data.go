package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{
	{Id: 1, Title: "1984", Finished: true},
	{Id: 2, Title: "To Kill a Mockingbird", Finished: false},
	{Id: 3, Title: "The Great Gatsby", Finished: true},
	{Id: 4, Title: "Pride and Prejudice", Finished: false},
	{Id: 5, Title: "The Catcher in the Rye", Finished: true},
	{Id: 6, Title: "Animal Farm", Finished: false},
	{Id: 7, Title: "Lord of the Flies", Finished: true},
	{Id: 8, Title: "Brave New World", Finished: false},
	{Id: 9, Title: "Fahrenheit 451", Finished: true},
	{Id: 10, Title: "The Hobbit", Finished: false},
}

func showAll() {

	for _, book := range books {
		book.showBook()
	}
}

func saveInTxt(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}

	defer file.Close()

	escritor := bufio.NewWriter(file)

	for _, book := range books {
		_, err = escritor.WriteString("Book: " + book.Title + " ID: " + strconv.Itoa(book.Id) + "\n")
	}

	if err != nil {
		fmt.Println("Error trying to write in the file", err)
		return
	}

	err = escritor.Flush()
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

}

// metodo asociado "Receiver"
func (book *Book) showBook() {
	fmt.Println("Title: ", book.Title)
	fmt.Println("Id: ", book.Id)
}

func getSize() int {
	return len(books)
}

func getLast() *Book {
	return books[len(books)-1]
}

func addBook(book *Book) (int, Book) {
	books = append(books, book)
	size := getSize()
	return size, *getLast()
}

func findBook(id int) (int, *Book) {
	index := -1
	var book *Book

	for i, b := range books {
		if b.Id == id {
			index = i
			book = b
			break
		}
	}

	return index, book
}

func finishBook(id int) {
	i, book := findBook(id)

	if i < 0 {
		return
	}

	book.Finished = true
	books[i] = book
	fmt.Printf("Finished book %s \n", book.Title)
}
