package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addNewContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Success added contact: %v\n", name)
}

func findContactByName(name string) *Contact {
	if index, exists := contactIndexByName[name]; exists {
		return &contactList[index]
	} else {
		return nil
	}
}

func ListContacts() {
	fmt.Println("--- Listing Contacts ---")
	if len(contactList) == 0 {
		fmt.Println("No contact found")
	}

	for index, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", index+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	addNewContact("Tomi", "tomi@exmaple.com", "123456789")
	addNewContact("Mandala", "mandala@exmaple.com", "0808080808")
	addNewContact("Putra", "putra@exmaple.com", "5555555555")
	addNewContact("Tomi", "tomtom@exmaple.com", "343433434343") // duplicate name

	ListContacts()

	var keyword string = "joko"
	result := findContactByName(keyword)
	if result == nil {
		fmt.Printf("Find contact by keyword %s not found", keyword)
	} else {
		fmt.Printf("Keyword is %s - contact found. Name is %s, email: %s\n", keyword, result.Name, result.Email)
	}
}
