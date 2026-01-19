package main

import "fmt"

/*
Pengertian Interface di Go-Lang merupakan kontrak prilaku.
Interface tidak peduli pada isi struct-nya, hanya peduli pada method-nya.
*/

/*
Interface Notifier
Siapapun yang memiliki method Send(message string) string
secara otomatis dianggap sebagai Notifier.
*/
type Notifier interface {
	Send(message string) string
}

type Email struct {
	Address string
}

// Implementasi interface Notifier pd method Send untuk parameter Email
func (e Email) Send(message string) string {
	return "Email sent to " + e.Address + " with message " + message
}

type SMS struct {
	PhoneNumber string
}

// Implementasi interface Notifier pd method Send untuk parameter SMS
func (s SMS) Send(message string) string {
	return "SMS sent to " + s.PhoneNumber + " with message " + message
}

/*
NotifyAll
Func ini menerima parameter n Notifier dan message string.
Func ini akan mengirimkan message ke n.
*/
func NotifyAll(n Notifier, message string) {
	result := n.Send(message)
	fmt.Println(result)
}

func main() {
	email := Email{Address: "example@gmail.com"}
	sms := SMS{PhoneNumber: "081234567890"}

	// fmt.Println(email.Send("Hello, Tomi"))
	// fmt.Println(sms.Send("Lorem ipsum dolor sit amet"))

	NotifyAll(email, "Hello, Tomi")
	NotifyAll(sms, "Lorem ipsum dolor sit amet")
}
