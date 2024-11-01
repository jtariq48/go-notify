package email_providers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"notify/config"
	"notify/models"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridProvider struct{}

// SendGrid processes the email notification using the SendGrid provider
func (s *SendGridProvider) Send(ctx context.Context, notification models.Notification) error {
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(notification.Payload), &payload)
	if err != nil {
		return err
	}

	log.Printf("SendGrid: Sending email to %s with subject %s", payload["to"], payload["subject"])

	// SendEmail()
	// Simulate sending email via SendGrid
	// Normally here you'd call SendGrid's API to send the email
	return nil
}

func SendEmail() {
	sg := sendgrid.NewSendClient(config.AppConfig.SendGridApiKey)
	from := mail.NewEmail("No Reply", config.AppConfig.FromEmail)
	subject := "Welcome To Aqary International and"
	to := mail.NewEmail("Junaid Tariq", "j.tariq@aqaryint.com")
	message := mail.NewV3MailInit(from, subject, to)

	// Create personalization and set dynamic template data
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("user_name", "Junaid Tariq")

	message.Personalizations = append(message.Personalizations, personalization)
	message.SetTemplateID("d-328983b366fe4d06a295e2df80b57471")

	response, err := sg.Send(message)
	if err != nil {
		log.Println(err)
	}
	log.Println(response.StatusCode)
}

// SendGrid processes the email notification using the SendGrid provider
// func (s *SendGridProvider) SendEmail(notification models.Notification) error {
// 	var payload map[string]interface{}
// 	err := json.Unmarshal([]byte(notification.Payload), &payload)
// 	if err != nil {
// 		return err
// 	}

// 	toEmail := payload["to"].(string)
// 	templateID := payload["templateid"].(string) // Assuming template_id is included in the payload

// 	// Ensure dynamic_data is a map of string key-value pairs
// 	dynamicData, ok := payload["data"].(map[string]interface{})
// 	if !ok {
// 		return err // Handle error for dynamic_data not being the expected type
// 	}

// 	// Convert dynamicData to map[string]string
// 	dynamicTemplateData := make(map[string]string)
// 	for key, value := range dynamicData {
// 		if strValue, ok := value.(string); ok {
// 			dynamicTemplateData[key] = strValue
// 		}
// 	}

// 	log.Printf("SendGrid: Sending email to %s using template %s", toEmail, templateID)

// 	// Create a new SendGrid client
// 	sendGridClient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

// 	// Create the email message
// 	from := mail.NewEmail("Your Name", "your-email@example.com") // Update sender's info
// 	to := mail.NewEmail("Recipient Name", toEmail)

// 	// Create a Personalization object to add dynamic data
// 	personalization := mail.NewPersonalization()
// 	personalization.AddTos(to)
// 	personalization.SetDynamicTemplateData("name", "Junaid")

// 	// Create the message and set the template ID
// 	message := mail.NewV3Mail(from, subject,  to,)
// 	message.SetTemplateID(templateID)
// 	message.AddPersonalization(personalization)

// 	// Send the email
// 	response, err := sendGridClient.Send(message)
// 	if err != nil {
// 		log.Printf("Failed to send email: %v", err)
// 		return err
// 	}

// 	log.Printf("Email sent successfully! Status Code: %d", response.StatusCode)
// 	return nil
// }

// func (s *SendGridProvider) SendEmail(notification models.Notification) error {
// 	var payload models.Email
// 	err := json.Unmarshal([]byte(notification.Payload), &payload)
// 	if err != nil {
// 		return err
// 	}

// 	m := mail.NewV3Mail()

// 	address := payload.
// 	name := "Example User"
// 	e := mail.NewEmail(name, address)
// 	m.SetFrom(e)

// 	m.SetTemplateID("d-c6dcf1f72bdd4beeb15a9aa6c72fcd2c")

// 	p := mail.NewPersonalization()
// 	tos := []*mail.Email{
// 		mail.NewEmail("Example User", "test1@example.com"),
// 	}
// 	p.AddTos(tos...)

// 	p.SetDynamicTemplateData("receipt", "true")
// 	p.SetDynamicTemplateData("total", "$ 239.85")

// 	items := []struct {
// 		text  string
// 		image string
// 		price string
// 	}{
// 		{"New Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/8dda1131320a6d978b515cc04ed479df259a458d5d45d58b6b381cae0bf9588113e80ef912f69e8c4cc1ef1a0297e8eefdb7b270064cc046b79a44e21b811802.png", "$ 79.95"},
// 		{"Old Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/3629f54390ead663d4eb7c53702e492de63299d7c5f7239efdc693b09b9b28c82c924225dcd8dcb65732d5ca7b7b753c5f17e056405bbd4596e4e63a96ae5018.png", "$ 79.95"},
// 		{"Blue Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/00731ed18eff0ad5da890d876c456c3124a4e44cb48196533e9b95fb2b959b7194c2dc7637b788341d1ff4f88d1dc88e23f7e3704726d313c57f350911dd2bd0.png", "$ 79.95"},
// 	}

// 	var itemList []map[string]string
// 	var item map[string]string
// 	for _, v := range items {
// 		item = make(map[string]string)
// 		item["text"] = v.text
// 		item["image"] = v.image
// 		item["price"] = v.price
// 		itemList = append(itemList, item)
// 	}
// 	p.SetDynamicTemplateData("items", itemList)

// 	p.SetDynamicTemplateData("name", "Sample Name")
// 	p.SetDynamicTemplateData("address01", "1234 Fake St.")
// 	p.SetDynamicTemplateData("address02", "Apt. 123")
// 	p.SetDynamicTemplateData("city", "Place")
// 	p.SetDynamicTemplateData("state", "CO")
// 	p.SetDynamicTemplateData("zip", "80202")

// 	m.AddPersonalizations(p)

// 	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
// 	request.Method = "POST"
// 	var Body = mail.GetRequestBody(m)
// 	request.Body = Body
// 	response, err := sendgrid.API(request)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(response.StatusCode)
// 		fmt.Println(response.Body)
// 		fmt.Println(response.Headers)
// 	}
// }

func sendGridEmail() {
	m := mail.NewV3Mail()

	address := "test@example.com"
	name := "Example User"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.SetTemplateID("d-c6dcf1f72bdd4beeb15a9aa6c72fcd2c")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Example User", "test1@example.com"),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("receipt", "true")
	p.SetDynamicTemplateData("total", "$ 239.85")

	items := []struct {
		text  string
		image string
		price string
	}{
		{"New Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/8dda1131320a6d978b515cc04ed479df259a458d5d45d58b6b381cae0bf9588113e80ef912f69e8c4cc1ef1a0297e8eefdb7b270064cc046b79a44e21b811802.png", "$ 79.95"},
		{"Old Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/3629f54390ead663d4eb7c53702e492de63299d7c5f7239efdc693b09b9b28c82c924225dcd8dcb65732d5ca7b7b753c5f17e056405bbd4596e4e63a96ae5018.png", "$ 79.95"},
		{"Blue Line Sneakers", "https://marketing-image-production.s3.amazonaws.com/uploads/00731ed18eff0ad5da890d876c456c3124a4e44cb48196533e9b95fb2b959b7194c2dc7637b788341d1ff4f88d1dc88e23f7e3704726d313c57f350911dd2bd0.png", "$ 79.95"},
	}

	var itemList []map[string]string
	var item map[string]string
	for _, v := range items {
		item = make(map[string]string)
		item["text"] = v.text
		item["image"] = v.image
		item["price"] = v.price
		itemList = append(itemList, item)
	}
	p.SetDynamicTemplateData("items", itemList)

	p.SetDynamicTemplateData("name", "Sample Name")
	p.SetDynamicTemplateData("address01", "1234 Fake St.")
	p.SetDynamicTemplateData("address02", "Apt. 123")
	p.SetDynamicTemplateData("city", "Place")
	p.SetDynamicTemplateData("state", "CO")
	p.SetDynamicTemplateData("zip", "80202")

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
