package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	timeout        = 20
	ch             = make(chan bool)
	otp     string = GenerateRandomNumber()
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func GenerateRandomNumber() string {
	rand.Seed(time.Now().UnixMilli())
	var out strings.Builder
	for i := 0; i < 4; i++ { //if you want change the max number just change 4 to another number
		out.WriteString(strconv.Itoa(rand.Intn(9)))
	}

	return out.String()
}
func main() {
	// Email as data sender.
	from := "your email"
	password := "your password email"

	// Receiver email address.
	to := []string{
		"receiver-email",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(otp)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Println("Input OTP Code before 20 seconds: ")
	fmt.Scan(&input)

	if input == otp {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}
