package soap

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestClient_Call(t *testing.T) {
	c := NewClient("http://192.168.1.1:1900/ipc", "urn:schemas-upnp-org:service:WANIPConnection:1")
	if false {
		r, _ := c.Encode("QueryStateVariable", Values{"varName": "PortMappingNumberOfEntries"})
		rr, _ := ioutil.ReadAll(r)
		log.Println("\n" + string(rr))
	}
	got, _ := c.Call("QueryStateVariable", Values{"varName": "PortMappingNumberOfEntries"})
	log.Println(got)
}
