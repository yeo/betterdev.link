package baja

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	// "text/template"

	"github.com/yeo/betterdev.link/baja/dao"
	"github.com/yeo/betterdev.link/baja/dts"
	"github.com/yeo/betterdev.link/baja/server"

	"crypto/sha256"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/jaytaylor/html2text"
)

func customEmail(doc *goquery.Document, issue, email string) (string, error) {
	// We just want to obfuscato email
	h := sha256.New()
	h.Write([]byte(email))

	doc.Find("a").Each(func(_ int, link *goquery.Selection) {
		href, ok := link.Attr("href")

		linkId := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(href))
		if ok {
			link.SetAttr("href", fmt.Sprintf("https://betterdev.link/links/%s?issue=%s&email=%x", linkId, issue, h.Sum(nil)))
		}
	})

	body, err := goquery.OuterHtml(doc.Selection)
	if err != nil {
		return "", err
	}

	//TODO: Refactor
	body = strings.Replace(body, "*|EMAIL_HASH|*", fmt.Sprintf("%x", h.Sum(nil)), -1)
	return body, nil
}

func sendTo(svc *ses.SES, issue, subject, email string) (*ses.SendEmailOutput, error) {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(fmt.Sprintf("./public/issues/%s/email/index.html", issue)) // Error handling elided for brevity.
	io.Copy(buf, f)                                                            // Error handling elided for brevity.
	f.Close()

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(buf.Bytes()))
	if err != nil {
		log.Fatal("Fail to parse email", err)
	}

	emailBody, err := customEmail(doc, issue, email)
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
				log.Println("Fail to send", email, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println("Fail to send", email, err.Error())
		}
		return nil, err
	}

	return result, nil
}

func Fanout(issueNumber string, mode string, confirm bool) {
	// Quick and dirty sent tracking
	// All of this crapy need to refactor
	config := server.LoadConfigFromEnv()
	db := dao.Connect(config.MongoURI)
	tracker := dts.TrackerService{db}
	// TODO: Close db

	log.Printf("Start to parse export contact from %s\n", mode)
	contacts, err := ioutil.ReadFile(fmt.Sprintf("./content/%s.csv", mode))
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	f, err := os.Stat(fmt.Sprintf("./content/issues/%s.yml", issueNumber))
	if err != nil {
		log.Fatal("Cannot open issue file")
	}
	issue, err := loadIssue(f)
	if err != nil {
		log.Fatal("Cannot parse issue content")
	}

	r := csv.NewReader(bytes.NewReader(contacts))
	subject := fmt.Sprintf("BetterDev #%s", issueNumber)
	if issue.Subject != "" {
		subject = fmt.Sprintf("%s - %s", subject, issue.Subject)
	}

	if mode == "dev" {
		subject = "[Test]" + subject
	}

	svc := ses.New(session.New())

	// Activity user
	file, err := os.Open(fmt.Sprintf("./content/%s", "no_activity.user"))
	var cleanUser map[string]bool
	cleanUser = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cleanUser[scanner.Text()] = true
	}
	log.Println(cleanUser)

	// Iterator throught the list and fanout email
	count := 0
	for {
		count = count + 1
		record, err := r.Read()
		if count <= 1 {
			log.Println("Skip header")
			continue
		}

		if err == io.EOF {
			break
		}
		log.Printf("Process at %d %v\n", count, record)

		if err != nil {
			log.Fatal(err)
		}

		if cleanUser[record[0]] == true {
			log.Println(record[0], "is ignored because user has not click any link in last 6 months")
			continue
		}

		if confirm == true {
			if tracker.AlreadySent(issueNumber, record[0]) == true {
				log.Println("Already fanout", issueNumber, "to", record)
				continue
			}

			result, err := sendTo(svc, issueNumber, subject, record[0])
			if err == nil {
				tracker.LogSent(issueNumber, record[0])

				log.Println("Send to", record[0], "succesfully", result)
			} else {
				log.Println("Fail to fanout", record[0], err)
			}
		} else {
			log.Println("DRY Fanout", issueNumber, subject, "to", record[0])
		}
	}
}
