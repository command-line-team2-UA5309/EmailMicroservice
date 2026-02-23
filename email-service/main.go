package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	godotenv.Load()
	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject: %s\n%s", req.Subject, req.Body))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{req.To}, message)
	if err != nil {
		http.Error(w, "fail to send email", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/send-email", sendEmailHandler)

	http.ListenAndServe(":8080", nil)
}
