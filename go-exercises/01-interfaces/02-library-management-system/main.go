package main

import "fmt"

type Readable interface {
	Read()
	GetInfo()
}

type ReadingMaterials interface {
	Title() string
	Author() string
	Genre() string
	Publisher() string
	Date() string
	IssueNumber() int
}

type MaterialAttr struct {
	title       string
	author      string
	genre       string
	publisher   string
	date        string
	issueNumber int
}

func (m MaterialAttr) Title() string {
	return m.title
}

func (m MaterialAttr) Author() string {
	return m.author
}

func (m MaterialAttr) Genre() string {
	return m.genre
}

func (m MaterialAttr) Publisher() string {
	return m.publisher
}

func (m MaterialAttr) Date() string {
	return m.date
}

func (m MaterialAttr) IssueNumber() int {
	return m.issueNumber
}

type ReadingMaterialType struct {
	ReadingMaterials
}

type Book struct {
	ReadingMaterialType
}

func (rm *ReadingMaterialType) GetInfo() {
	fmt.Printf(`Info:
   Title: %s
   Author: %s
   Genre: %s
   Publisher: %s
   Date: %s
   Issue Number: %d \n`, rm.Title(), rm.Author(), rm.Genre(), rm.Publisher(), rm.Date(), rm.IssueNumber())
}

func (rm *ReadingMaterialType) Read() {
	fmt.Printf("Currently Reading: %s \n", rm.Title())
}

func main() {
	library := []Readable{
		NewBook("Atomic Habits", "James Clear", "Self Improvement", "Shroff Publications", "2024-10-06", 2),
	}

	for _, l := range library {
		l.Read()
		l.GetInfo()
	}
}

func NewBook(title, author, genre, publisher, date string, issueNumber int) *Book {
	return &Book{ReadingMaterialType: ReadingMaterialType{MaterialAttr{title, author, genre, publisher, date, issueNumber}}}
}
