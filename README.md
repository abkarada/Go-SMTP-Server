# 📬 Go SMTP Server

A lightweight and fully functional **SMTP server** built using the [emersion/go-smtp](https://github.com/emersion/go-smtp) library in **Golang**.  
This server accepts incoming emails, logs sender, recipients, and message body, and is suitable for testing, local development, or building custom email processing logic.

---

## 🚀 Features

- 📥 Accepts and logs `MAIL FROM`, `RCPT TO`, and `DATA`
- 🔧 Clean implementation using Go interfaces
- 📡 Listens on custom port (`:1025`)
- 🔐 Supports optional plaintext authentication (for testing)
- 🧪 Ideal for development/testing environments

---

## 🧱 Tech Stack

- **Language**: Go (Golang)
- **SMTP Library**: [github.com/emersion/go-smtp](https://github.com/emersion/go-smtp)

---

## 📁 Project Structure

Go_Mail/
├── main.go # Main application with SMTP server implementation
├── go.mod # Module declaration
├── go.sum # Dependency hash tracking

---

## 🛠 How to Run

1. Make sure you have Go installed:  
   [https://go.dev/doc/install](https://go.dev/doc/install)

2. Clone the repository:

git clone https://github.com/yourusername/Go_Mail.git
cd Go_Mail

3.Initialize Modul

go mod init go_mail
go get github.com/emersion/go-smtp

4.Run the Server

go run main.go

By default, the server listens on:
localhost:1025



