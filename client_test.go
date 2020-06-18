package soap

import (
	"log"
	"testing"
)

func TestClient_Call(t *testing.T) {
	ipc := "http://192.168.1.1:1900/ipc"
	schema := "urn:schemas-upnp-org:control"
	c := NewClient(ipc, schema)
	got, err := c.Call("QueryStateVariable", Values{"varName": "PortMappingNumberOfEntries"})
	if err != nil {
		t.Fatal(err)
	}
	log.Println(got["return"])
}
