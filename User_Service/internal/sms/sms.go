package sms

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
)

func SendSMS(verificationCode string) {
	sms := mail.NewMessage()
	sms.SetHeader("From", "asilbekxolamtov2002@gmail.com")
	sms.SetHeader("To", "besthacker8163264@gmail.com")
	sms.SetHeader("Subject", "Your verification code:")
	sms.SetBody("text/plain", verificationCode)

	d := mail.NewDialer("smtp.gmail.com", 587, "asilbekxolmatov2002@gmail.com", "ylsq ffrt nrhc dmnb")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(sms); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Verification code has been sent, check your email")
}
