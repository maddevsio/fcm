package fcm

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendTopic(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handleTopic))
	fcmServerURL = srv.URL

	defer srv.Close()

	c := NewFCM("key")

	data := map[string]string{
		"first":  "value",
		"second": "value",
	}

	res, err := c.Send(&Message{
		Data: data,
		To:   "/topics/topicName",
	})
	if err != nil {
		t.Error("Response Error : ", err)
	}
	if res == nil {
		t.Error("Res is nil")
	}
}

func TestSendMessageCanSendToMultipleRegIDs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handleRegs))
	fcmServerURL = srv.URL

	defer srv.Close()

	c := NewFCM("key")

	data := map[string]string{
		"msg": "Hello World",
		"sum": "Happy Day",
	}

	ids := []string{
		"token0",
		"token1",
		"token2",
	}

	res, err := c.Send(&Message{
		Data:            data,
		RegistrationIDs: ids,
	})
	if err != nil {
		t.Error("Response Error : ", err)
	}
	if res == nil {
		t.Error("Res is nil")
	}

	if res.Success != 2 || res.Fail != 1 {
		t.Error("Parsing Success or Fail error")
	}
}

func handleTopic(w http.ResponseWriter, r *http.Request) {
	result := `{"message_id":6985435902064854329}`

	fmt.Fprintln(w, result)
}

func handleRegs(w http.ResponseWriter, r *http.Request) {
	result := `{"multicast_id":1003859738309903334,"success":2,"failure":1,"canonical_ids":0,"results":[{"message_id":"0:1448128667408487%ecaaa23db3fd7efd"},{"message_id":"0:1468135657607438%ecafacddf9ff8ead"},{"error":"InvalidRegistration"}]}`
	fmt.Fprintln(w, result)

}
