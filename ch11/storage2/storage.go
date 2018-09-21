package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 { return 1000000000 * 0.98 }

const sender = "user@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
const template = `Внимание, вы использовали %d байт хранилища,
                  %d%% вашей квоты.`

var notifyUser = func(username string, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("Сбой smtp.SendMail(%s): %s", username, err)
	}

}

func checkQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
