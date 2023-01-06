package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	DbUsername := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASS")
	DbAddress := os.Getenv("DB_ADDR")
	DbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", DbUsername, DbPassword, DbAddress, DbName))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT name, email, message FROM mailing`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	list := []Mail{}
	for rows.Next() {
		var mail Mail
		err = rows.Scan(&mail.name, &mail.email, &mail.message)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, mail)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Got mailing list from db")

	from := os.Getenv("FROM")
	passwd := os.Getenv("PASSWD")
	hostname := "smtp.yandex.ru"

	for _, mail := range list {
		log.Println("Send message to", mail.email)
		body, err := getBody(&mail, from)
		if err != nil {
			log.Fatal(err)
		}
		auth := smtp.PlainAuth("", from, passwd, hostname)
		fmt.Println(body)
		err = smtp.SendMail(hostname+":25", auth, from, []string{string(mail.email)}, []byte(body))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Done")
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

}

func getBody(m *Mail, from string) (string, error) {
	buf := new(bytes.Buffer)
	tmp, err := template.New("mail").Parse(mailTmpl)
	if err != nil {
		return "", err
	}
	if err := tmp.Execute(buf, struct {
		To   string
		Body string
		From string
	}{
		To:   m.email,
		Body: m.message,
		From: from,
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

const mailTmpl = `To: {{.To}}
From: {{.From}}
Subject: Проспал жизнь
Content-Type: text/html

<html>
	<body>
		<div>
			<table style="background-color: #FFE4C4">
				<tr>
					<td>
						<span style="background-color: #DEB887">
						<i> {{.Body}} </i>
						</span>
					</td>
				</tr>
			</table>
		</div>
	</body>
</html>
`

type Mail struct {
	name string
	email string
	message string
}
