package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)


func main() {
	from := os.Getenv("FROM")
	passwd := os.Getenv("PASSWD")
	hostname := "smtp.yandex.ru"
	auth := smtp.PlainAuth("", from, passwd, hostname)

	input, _ := ioutil.ReadAll(os.Stdin)
	lines := strings.Split(string(input), "\n")
	to := lines[0]
	subject := lines[1]
	body := strings.Join(lines[2:], "\n")

	msg := fmt.Sprintf(`From: %s
To: %s
Subject: %s
Content-Type: text/plain

%s
`, from, to, subject, body)
	fmt.Println(string(msg))

	err := smtp.SendMail(hostname+":25", auth, from, []string{string(to)}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

