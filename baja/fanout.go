package baja

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	// "strings"
	// "text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/jaytaylor/html2text"
)

func customEmail(doc *goquery.Document, email string) (string, error) {
	doc.Find("a").Each(func(_ int, link *goquery.Selection) {
		href, ok := link.Attr("href")

		linkId := base64.StdEncoding.EncodeToString([]byte(href))
		if ok {
			link.SetAttr("href", fmt.Sprintf("https://betterdev.link/links/%s?email=%s", linkId, email))
		}
	})

	return goquery.OuterHtml(doc.Selection)
}

func sendTo(svc *ses.SES, doc *goquery.Document, subject, email string) (*ses.SendEmailOutput, error) {
	emailBody, err := customEmail(doc, email)
	if err != nil {
		log.Fatal("Fail to build custom email")
	}

	emailText, err := html2text.FromString(emailBody, html2text.Options{PrettyTables: true})
	if err != nil {
		log.Fatal("Cannot convert HTML email to Text")
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			//CcAddresses: []*string{
			//	aws.String("vinh@noty.im"),
			//},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(emailBody),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(emailText),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		//ReturnPath:    aws.String(""),
		//ReturnPathArn: aws.String(""),
		Source: aws.String("vinh@yeo.space"),
		//SourceArn:     aws.String(""),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil
}

func Fanout(issue string, mode string) {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(fmt.Sprintf("./public/issues/%s/email/index.html", issue)) // Error handling elided for brevity.
	io.Copy(buf, f)                                                            // Error handling elided for brevity.
	f.Close()

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(buf.Bytes()))
	if err != nil {
		log.Fatal("Fail to parse email", err)
	}

	log.Printf("Start to parse export contact from %s\n", mode)
	contacts, err := ioutil.ReadFile(fmt.Sprintf("./content/%s.csv", mode))
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	r := csv.NewReader(bytes.NewReader(contacts))
	subject := fmt.Sprintf("Better Dev Link #%s", issue)
	if mode == "dev" {
		subject = "[Test]" + subject
	}

	svc := ses.New(session.New())
	count := 0
	for {
		record, err := r.Read()
		if count == 0 {
			log.Println("Skip header")
			count += 1
			continue
		}

		if err == io.EOF {
			break
		}
		log.Printf("Process at %s %v\n", count, record)

		if err != nil {
			log.Fatal(err)
		}

		result, err := sendTo(svc, doc, subject, record[0])
		log.Println(result)
	}
}
