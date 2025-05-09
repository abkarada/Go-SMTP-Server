package main

import (
	"io"
	"log"

	"github.com/emersion/go-smtp"
)

/*
SMTP server setup in Go using the emersion/go-smtp library.

1. Define a Backend and Session structure
2. Backend handles authentication (Login or AnonymousLogin)
3. Session handles the actual email transaction:
   - MAIL FROM
   - RCPT TO
   - DATA
   - LOGOUT
*/

/*
Backend is the core authentication handler. When a client connects,
the server will first interact with this backend to authenticate.

- Login() is called on AUTH command.
- AnonymousLogin() is used if no AUTH is required.
*/

type Backend struct{}

func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	log.Println("Login attempt from:", state.RemoteAddr)
	// TODO: Add actual authentication logic
	return &Session{}, nil
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	log.Println("Anonymous login from:", state.RemoteAddr)
	return &Session{}, nil
}

/*
Session holds the state of an individual SMTP mail transaction.

- Mail()    → Called when sender is specified
- Rcpt()    → Called for each recipient
- Data()    → Called when the message body is received
- Reset()   → Called on RSET command
- Logout()  → Called on QUIT command
*/

type Session struct{}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Sender:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Recipient:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	log.Println("Message body:", string(body))
	return nil
}

func (s *Session) Reset() {
	log.Println("Session reset")
}

func (s *Session) Logout() error {
	log.Println("Session closed")
	return nil
}

func main() {
	backend := &Backend{}

	server := smtp.NewServer(backend)
	server.Addr = ":1025"
	server.Domain = "localhost"     // Must be a valid hostname (not purely numeric)
	server.AllowInsecureAuth = true // Allow plaintext login for testing

	server.MaxMessageBytes = 1024 * 100 // Max size: 100 KB
	server.MaxRecipients = 100          // Max recipients per message

	log.Println("Starting SMTP server at", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}

// Note: This is a basic SMTP server setup. In production, ensure to use TLS and proper authentication methods.
// Also, consider implementing proper error handling and logging.
// For testing, you can use telnet or a mail client to connect to the server and send emails.
// To test, run the server and use a telnet client to connect:
// telnet localhost 1025
// Then, you can issue SMTP commands like:
// HELO localhost
// MAIL FROM:<sender@example.com>
