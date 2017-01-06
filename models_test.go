package fcm

import "testing"

func TestResponse_GetRetryAfterTime(t *testing.T) {
	r := &Response{
		RetryAfter: "120s",
	}
	_, err := r.GetRetryAfterTime()
	if err != nil {
		t.Error(err)
	}

}
