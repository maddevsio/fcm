# Firebase Cloud Notifications Client

[![Go Report Card](https://goreportcard.com/badge/github.com/maddevsio/fcm)](https://goreportcard.com/report/github.com/maddevsio/fcm)
[![Build Status](https://travis-ci.org/maddevsio/fcm.svg)](https://travis-ci.org/maddevsio/fcm.svg)[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
[![](https://godoc.org/github.com/maddevsio/fcm?status.svg)](https://godoc.org/github.com/maddevsio/fcm)
[![Coverage Status](https://coveralls.io/repos/github/maddevsio/fcm/badge.svg?branch=master)](https://coveralls.io/github/maddevsio/fcm?branch=master)

Firebase Cloud Messaging for application servers implemented using the Go programming language.
It's designed for simple push notification sending via HTTP API

# Getting started

To install fcm, use go get:

```
go get gopkg.in/maddevsio/fcm.v1
```

Import fcm with the following:

```
import "gopkg.in/maddevsio/fcm.v1"
```

# Sample usage

```
package main

import (
	"fmt"
	"log"

	"gopkg.in/maddevsio/fcm.v1"
)

func main() {
	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	c := fcm.NewFCM("serverKey")
	token := "token"
	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
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

```

More examples can be found in /_examples/ directory

# License

MIT License

Copyright (c) 2017 Mad Devs Developers

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
