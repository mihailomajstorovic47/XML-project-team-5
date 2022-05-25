package application

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	Email    string
	Password string
	Url      string
}

func NewEmailService(email, password, host, port string) *EmailService {
	return &EmailService{
		Email:    email,
		Password: password,
		Url:      fmt.Sprintf("http://%s:%s", host, port),
	}
}

func (service *EmailService) SendEmail(sendTo, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", service.Email)
	msg.SetHeader("To", sendTo)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	n := gomail.NewDialer("smtp.office365.com", 587, service.Email, service.Password)
	err := n.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	return err
}

func (service *EmailService) SendVerificationEmail(sendTo, username, verificationCode string) error {

	href := fmt.Sprintf("%s/verify/%s/%s", service.Url, username, verificationCode)
	body := `<div style="background: linear-gradient(#00FF00 , #48c6ef); height: 320px; width: 500px; font-family: Montserrat, sans-serif; text-align: center; align-items: center; color: white; margin: 10px; padding: 4px; -webkit-box-shadow: 0px 7px 12px -6px rgba(0,0,0,0.62); box-shadow: 0px 7px 12px -6px rgba(0,0,0,0.62); border-radius: 10px;">
			<h1>Welcome to Dislink</h1>
			<h3 style="font-weight: normal;">Verify <b>` + username + `</b> account</h3>

			<a href="` + href + `" style="box-shadow: -2px 7px 4px 0px #004cff18;
			background-color:#ffffff;
			border-radius:10px;
			border-width: 0;
			display:inline-block;
			cursor:pointer;
			color:#48c5ef;
			font-family:Arial;
			font-size:20px;
			font-weight:bold;
			padding:14px 50px;
			text-decoration:none;
			text-shadow:0px 2px 10px #48c5ef;">Verify</a>
		</div>`

	return service.SendEmail(sendTo, "Dislinkt Verification", body)
}

func (service *EmailService) SendRecoveryEmail(sendTo, username, recoveryCode string) error {
	body := `    
		<div style="background: linear-gradient(#00FF00 , #48c6ef); height: 320px; width: 500px; font-family: Montserrat, sans-serif; text-align: center; align-items: center; color: white; margin: 10px; padding: 4px; -webkit-box-shadow: 0px 7px 12px -6px rgba(0,0,0,0.62); box-shadow: 0px 7px 12px -6px rgba(0,0,0,0.62); border-radius: 10px;">
        <h1>Dislinkt Recovery</h1>
        <h4 style="font-weight: normal;">Recover your <b>` + username + `</b> account with code</h4> 

        <div style="box-shadow: -2px 7px 4px 0px #004cff18;
        background-color:#ffffff;
        border-radius:10px;
        border-width: 0;
        display:inline-block;
        cursor:auto;
        color:#48c5ef;
        font-family:Arial;
        font-size:20px;
        font-weight:bold;
        padding:14px 50px;
        text-decoration:none;
        text-shadow:0px 2px 25px #48c5ef;">CODE: ` + recoveryCode + `</div>

    </div>`
	return service.SendEmail(sendTo, "Dislinkt Recovery", body)
}