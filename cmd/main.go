package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	sid   = os.Getenv("SERVICE_ID")
	token = os.Getenv("AUTH_TOKEN")
	accId = os.Getenv("ACCOUNT_ID")
	me    = os.Getenv("PHONE")
)

func main() {
	client := *twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: token,
		Username: accId,
	})

	SendSMS(client)
	// CheckOTPCode(client, "")
}

func SendSMS(client twilio.RestClient) {
	params := verify.CreateVerificationParams{}

	params.SetTo(me)
	params.SetChannel("sms")

	res, err := client.VerifyV2.CreateVerification(sid, &params)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("sent verification code to %s\n", *res.To)
}

func CheckOTPCode(client twilio.RestClient, code string) {
	params := verify.CreateVerificationCheckParams{}

	params.SetTo(me)
	params.SetCode(code)

	res, err := client.VerifyV2.CreateVerificationCheck(sid, &params)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("status: %s\n", *res.Status)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	sid = os.Getenv("SERVICE_ID")
	token = os.Getenv("AUTH_TOKEN")
	accId = os.Getenv("ACCOUNT_ID")
	me = os.Getenv("PHONE")
}
