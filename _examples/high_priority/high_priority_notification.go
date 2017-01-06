package main

import (
	"fmt"
	"log"

	"github.com/maddevsio/fcm"
)

func main() {
	data := map[string]string{
		"first":  "World",
		"second": "Hello",
	}
	c := fcm.NewFCM("serverKey")
	token := "token"
	response, err := c.Send(&fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityNormal,
		Notification: &fcm.Notification{
			Title: "Hello",
			Body:  "World",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
}
