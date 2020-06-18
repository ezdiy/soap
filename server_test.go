package soap

import (
	"net/http"
	"testing"
	"time"
)

func TestServer_ServeHTTP(t *testing.T) {
	s := &Server{}
	s.Handlers = map[string]Handler{
		"TestRPCCall": func(r *Request) error {
			r.Ret["Ret"] = r.Arg["Arg"] + "123"
			return nil
		},
	}
	go http.ListenAndServe("127.0.0.1:19000", s)
	time.Sleep(100 * time.Millisecond)
	c := NewClient("http://127.0.0.1:19000", "")
	got, _ := c.Call("TestRPCCall", Values{"Arg":"testarg"})
	if got["Ret"] != "testarg123" {
		t.Fatal("failed round trip")
	}
}
