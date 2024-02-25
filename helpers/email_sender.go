package helpers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"

	"github.com/akhil-is-watching/techletics_alumni_reg/config"
	qrcode "github.com/skip2/go-qrcode"
)

// func SendMail(to string, content string) error {

// 	config, err := config.LoadConfig()

// 	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpHost)

// 	client, err := smtp.Dial(fmt.Sprintf("%s:%s", config.SmtpHost, config.SmtpPort))
// 	if err != nil {
// 		fmt.Println("Error connecting to the SMTP server:", err)
// 		return err
// 	}
// 	defer client.Close()

// 	if err := client.StartTLS(&tls.Config{ServerName: config.SmtpHost}); err != nil {
// 		return fmt.Errorf("failed to start TLS: %w", err)
// 	}

// 	if err := client.Auth(auth); err != nil {
// 		fmt.Println("Error authenticating:", err)
// 		return err
// 	}

// 	if err := client.Mail(config.SmtpUsername); err != nil {
// 		fmt.Println("Error setting sender:", err)
// 		return err
// 	}
// 	if err := client.Rcpt(to); err != nil {
// 		fmt.Println("Error setting recipient:", err)
// 		return err
// 	}

// 	wc, err := client.Data()
// 	if err != nil {
// 		fmt.Println("Error getting data:", err)
// 		return err
// 	}
// 	_, err = fmt.Fprintf(wc, content)
// 	if err != nil {
// 		fmt.Println("Error writing message:", err)
// 		return err
// 	}
// 	err = wc.Close()
// 	if err != nil {
// 		fmt.Println("Error closing data:", err)
// 		return err
// 	}

// 	fmt.Println("Email sent successfully!")
// 	return nil
// }

// func SendMail(to string, content string) error {
// 	config, _ := config.LoadConfig()

// 	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", config.SmtpUsername, to, "Techletics Alumni Registration", "Welcome to techletics alumni registration")
// 	// Authenticate with the SMTP server
// 	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpHost)

// 	// Connect to the SMTP server
// 	conn, err := smtp.Dial(fmt.Sprintf("%s:%s", config.SmtpHost, config.SmtpPort))
// 	if err != nil {
// 		fmt.Println("Error connecting to the server:", err)
// 		return err
// 	}

// 	if err := conn.StartTLS(&tls.Config{ServerName: config.SmtpHost}); err != nil {
// 		return fmt.Errorf("failed to start TLS: %w", err)
// 	}

// 	if err := conn.Auth(auth); err != nil {
// 		fmt.Println("Error authenticating:", err)
// 		return err
// 	}

// 	// Send the email
// 	if err := conn.Mail(config.SmtpUsername); err != nil {
// 		fmt.Println("Error sending MAIL command:", err)
// 		return err
// 	}

// 	if err := conn.Rcpt(to); err != nil {
// 		fmt.Println("Error sending RCPT command:", err)
// 		return err
// 	}
// 	w, err := conn.Data()
// 	if err != nil {
// 		fmt.Println("Error sending DATA command:", err)
// 		return err
// 	}
// 	_, err = w.Write([]byte(message))
// 	if err != nil {
// 		fmt.Println("Error writing message body:", err)
// 		return err
// 	}
// 	err = w.Close()
// 	if err != nil {
// 		fmt.Println("Error closing connection:", err)
// 		return err
// 	}

// 	// Close the connection
// 	conn.Quit()
// 	return nil
// }

func SendMail(to string, qrCodeData string) error {
	// Load configuration
	config, _ := config.LoadConfig()
	png, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	qrFileName := UIDGen().GenerateID("RQR")
	GetS3Uploader().UploadBytes(qrFileName, png)

	// Load the email template
	templateData := struct {
		QRCode string
		Host   string
	}{
		QRCode: fmt.Sprintf("https://techleticsassetbucket.s3.ap-south-1.amazonaws.com/%s", qrFileName),
		Host:   "Christ College of Engineering",
	}

	// Parse the email template
	emailTemplate := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Welcome to Techletics 2024</title>
    </head>
    <body>
        <h2>Welcome to Techletics 2024 at {{.Host}}!</h2>
        <p>Dear Alumni,</p>
        <p>We are excited to invite you to Techletics 2024, our annual techfest at {{.Host}}.</p>
        <p>This year, Techletics promises to be bigger and better with a lineup of exciting events, workshops, and guest speakers.</p>
        <p>We look forward to your participation and contribution to make Techletics 2024 a grand success!</p>
        <p>Attached below is your personalized QR code for easy check-in at the event:</p>
        <!-- QR code image goes here -->
        <img src="{{.QRCode}}" alt="Techletics 2024 QR Code"> 
		<div style="min-height:100px;background-color:black; padding:10px">
			<img src="https://techleticsassetbucket.s3.ap-south-1.amazonaws.com/levantate_banner.jpeg" alt="Levantate Banner" style="height:100px">
		</div>
    </body>
    </html>
    `

	// Execute the template
	var tpl bytes.Buffer
	t := template.Must(template.New("email").Parse(emailTemplate))
	if err := t.Execute(&tpl, templateData); err != nil {
		return err
	}
	emailContent := tpl.String()

	// Compose the email message
	message := "From: " + config.SmtpUsername + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + "Invitation to Techletics 2024" + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + emailContent

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpHost)

	// Connect to the SMTP server
	conn, err := smtp.Dial(fmt.Sprintf("%s:%s", config.SmtpHost, config.SmtpPort))
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return err
	}

	if err := conn.StartTLS(&tls.Config{ServerName: config.SmtpHost}); err != nil {
		return fmt.Errorf("failed to start TLS: %w", err)
	}

	if err := conn.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		return err
	}

	// Set sender and recipient
	fromAddr := mail.Address{Name: "", Address: config.SmtpUsername}
	toAddr := mail.Address{Name: "", Address: to}

	// Send the email
	if err := conn.Mail(fromAddr.Address); err != nil {
		fmt.Println("Error sending MAIL command:", err)
		return err
	}

	if err := conn.Rcpt(toAddr.Address); err != nil {
		fmt.Println("Error sending RCPT command:", err)
		return err
	}

	// Send email body
	wc, err := conn.Data()
	if err != nil {
		fmt.Println("Error sending DATA command:", err)
		return err
	}
	defer wc.Close()

	buf := bytes.NewBufferString(message)
	if _, err := wc.Write(buf.Bytes()); err != nil {
		fmt.Println("Error writing message body:", err)
		return err
	}

	// Close the connection
	conn.Quit()
	return nil
}
