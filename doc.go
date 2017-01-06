/*
Package firebase provides integration with Firebase Cloud Notification HTTP API https://firebase.google.com/docs/cloud-messaging/http-server-ref
You can send push notifications to Android and iOS devices via simple API.

Example:
	package main

	import (
		"fmt"
		"log"

		"github.com/maddevsio/fcm"
	)

	func main() {
		data := map[string]string{
			"msg": "Hello World1",
			"sum": "Happy Day",
		}
		c := fcm.NewFCM("serverKey")
		token := "token"
		response, err := c.Send(&fcm.Message{
			Data:             data,
			RegistrationIDs:  []string{token},
			ContentAvailable: true,
			Priority:         fcm.PriorityHigh,
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



If you want to send notification with Sound or Badge, then use:
	response, err := c.Send(&fcm.Message{
		Notification: &fcm.Notification{
			Title: "Hello",
			Body:  "World",
			Sound: "default",
			Badge: "3",
		},
	})

*/
package fcm
