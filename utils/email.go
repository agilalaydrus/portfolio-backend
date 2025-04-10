package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

func SendResendContact(name, email, phone, linkedin, message string) error {
	client := resty.New()

	from := os.Getenv("RESEND_FROM")
	to := os.Getenv("RESEND_TO")
	apiKey := os.Getenv("RESEND_API_KEY")

	body := fmt.Sprintf(`
		<b>Name:</b> %s<br>
		<b>Email:</b> %s<br>
		<b>Phone:</b> %s<br>
		<b>LinkedIn:</b> %s<br>
		<b>Message:</b><br>%s
	`, name, email, phone, linkedin, message)

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"from":    from,
			"to":      to,
			"subject": "New Contact Submission",
			"html":    body,
		}).
		Post("https://api.resend.com/emails")

	if err != nil {
		fmt.Println("❌ Resend error:", err)
		return err
	}

	fmt.Println("✅ Resend status:", resp.Status())
	return nil
}
