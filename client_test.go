package soap

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestClient_Call(t *testing.T) {
	//ipc := "http://192.168.1.1:1900/ipc"
	ipc := "http://192.168.1.1:1900/upnp/control/WANIPConn1"
	ipc = "http://192.168.1.1:1900/upnp/control/WANCommonIFC1"
	schema := "urn:schemas-upnp-org:control"
	c := NewClient(ipc, schema)
	if false {
		r, _ := c.Encode("QueryStateVariable", Values{"varName": "PortMappingNumberOfEntries"})
		rr, _ := ioutil.ReadAll(r)
		log.Println("\n" + string(rr))
	}
	got, err := c.Call("QueryStateVariable", Values{"varName": "PortMappingNumberOfEntries"})
	log.Println(got, err)
}
