package email

import (
	"crypto/tls"
	"log"
	"net"
	"net/smtp"
	"time"
)

type EmailControllerConfig struct {
	UserEmail    string
	UserPassword string
	Recepients   []string
	SmtpPort     string
	SmtpHost     string
}

type EmailController struct {
	Config EmailControllerConfig
	auth   smtp.Auth
	conn   *smtp.Client
}

type Message struct {
	To      []string
	Subject string
	Body    string
}

// Инициализация SMTP-контроллера
func Init(c EmailControllerConfig) (*EmailController, error) {
	// Авторизация
	auth := smtp.PlainAuth("", c.UserEmail, c.UserPassword, c.SmtpHost)

	// Подключение к SMTP
	smtpServer := c.SmtpHost + ":" + c.SmtpPort
	dialer := &net.Dialer{Timeout: 10 * time.Second}

	conn, err := dialer.Dial("tcp", smtpServer)
	if err != nil {
		log.Panicln("Не удалось подключится к SMTP-серверу: ", err)
	}

	smtpClient, err := smtp.NewClient(conn, c.SmtpHost)
	if err != nil {
		log.Panicln("Не удалось подключится к SMTP-серверу: ", err)
	}

	// Запуск TLS
	if err = smtpClient.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
		log.Panicln("Ошибка при TLS-подключении: ", err)
	}

	// Авторизация
	if err = smtpClient.Auth(auth); err != nil {
		log.Panicln("Аутентификация провалилась:", err)
	}

	return &EmailController{
		Config: c,
		auth:   auth,
		conn:   smtpClient,
	}, nil
}

func (c *EmailController) SendMessage(m Message) {
	subject := "Subject: " + m.Subject + "\n"
	message := []byte(subject + "\n" + m.Body)

	// Устанавливаем отправителя
	if err := c.conn.Mail(c.Config.UserEmail); err != nil {
		log.Println("Ошибка при установке отправителя: ", c.Config.UserEmail, err)
		return
	}

	// Устанавливаем получателей
	for _, recipient := range m.To {
		if err := c.conn.Rcpt(recipient); err != nil {
			log.Println("Ошибка при установке получателя: ", recipient, err)
			return
		}
	}

	// Получаем SMTP-writer
	writer, err := c.conn.Data()
	if err != nil {
		log.Println("Ошибка при получении SMTP-writer: ", err)
		return
	}

	// Записываем сообщение
	_, err = writer.Write(message)
	if err != nil {
		log.Println("Ошибка при записи сообщения: ", string(message), err)
		return
	}

	// Закрываем SMTP-writer
	err = writer.Close()
	if err != nil {
		log.Println("Ошибка при закрытии SMTP-writer: ", err)
		return
	}
}
