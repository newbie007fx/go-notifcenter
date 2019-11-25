package notif

import (
	"net/http"
	"time"

	"github.com/imroc/req"
)

type action func() *Response

type NotifCenter struct {
	timeout *time.Duration
	retry   *Retry
	url     string
}

func CreateNewClient(url string) *NotifCenter {
	return &NotifCenter{url: url}
}

func (nc *NotifCenter) SetTimeout(timeout time.Duration) {
	nc.timeout = &timeout
}

func (nc *NotifCenter) SetRetry(count uint, backoff Backoff) {
	retry := Retry{
		Count:   count,
		Backoff: backoff,
	}
	nc.retry = &retry
}

func (nc *NotifCenter) SendAPIReq(body NotifBody) *Response {
	if nc.timeout != nil {
		req.SetTimeout(*nc.timeout)
	}
	action := func() *Response {
		var resp *http.Response
		r, err := req.Post(nc.url, req.BodyJSON(&body))
		if err == nil {
			resp = r.Response()
		}
		return &Response{Error: err, resp: resp}
	}

	return nc.do(action)
}

func (nc *NotifCenter) do(act action) *Response {
	var resp *Response
	retry := nc.retry
	if retry != nil {
		for i := 0; i < int(nc.retry.Count); i++ {
			resp = act()
			if resp.resp != nil && resp.resp.StatusCode < http.StatusInternalServerError {
				break
			}
			time.Sleep(retry.Backoff.next(i))
		}
		return resp
	}
	return act()
}
