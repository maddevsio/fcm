// Copyright (c) 2017 Mad Devs Developers All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package firebase provides integration with Firebase Cloud Notification HTTP API https://firebase.google.com/docs/cloud-messaging/http-server-ref
You can send push notifications to Android and iOS devices via simple API.

Example:
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	c := firebase.NewFCM("serverKey")
	response, err := c.Send(&firebase.Message{
		Data:             data,
		RegistrationIds:  []string{token},
		ContentAvailable: true,
		Priority:         firebase.PriorityHigh,
		Notification: &firebase.Notification{
			Title: "Hello",
			Body:  "World",
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.Canonical_ids)
	fmt.Println("Topic MsgId   :", response.MsgId)

If you want to send notification with Sound or Badge, then use:
	response, err := c.Send(&firebase.Message{
		Notification: &firebase.Notification{
			Title: "Hello",
			Body:  "World",
			Sound: "default",
			Badge: "3",
		},
	})

*/
package fcm
