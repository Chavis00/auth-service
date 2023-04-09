package pkg

import (
	"fmt"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"authservice/helpers"
)

func SendGridMail(name, email, subject, fileName, token string) (*rest.Response, error) {
	from := mail.NewEmail("Chavis", "chavis.delcourt@gmail.com")
	to := mail.NewEmail(name, email)
	subjectMail := subject
	template := helpers.ParseHtml(fileName, map[string]string{
		"to":    email,
		"token": token,
	})

	message := mail.NewSingleEmail(from, subjectMail, to, "", template)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return response, err
}
