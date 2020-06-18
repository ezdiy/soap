// Implements simple boilerplate for SOAP servers and clients.
package soap

import (
	"encoding/xml"
	"io"
	"net/http"
)

// Single handler for an RPC call
type Handler func(r *Request) error

// Represents whole RPC state lifetime of each request
type Request struct {
	Req  *http.Request // The underlying http request
	Name string        // The name of the RPC call
	Arg  Values        // Arguments k/v map
	Ret  Values        // Return values k/v map
}

// SOAP server is simply a map of handler methods
type Server struct {
	Handlers map[string]Handler
}

// Perform a single request/response round trip via provided map of handlers.
func (r *Request) Dispatch(rd io.Reader, wr io.Writer, handlers map[string]Handler) (err error) {
	x := &Envelope{}
	err = xml.NewDecoder(rd).Decode(&x)
	if err != nil {
		return
	}
	r.Name = x.Body.Item.XMLName.Local
	if h, ok := handlers[r.Name]; ok {
		r.Arg = x.Unmarshal()
		r.Ret = make(Values)
		err = h(r)
		if err == nil {
			x.Body.Item.XMLName.Local += r.Name + "Response"
			x.Marshal(r.Ret)
			err = xml.NewEncoder(wr).Encode(x)
		}
	}
	return err
}

// Wrapper for Dispatch() to make a generic http service.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := req.Dispatch(r.Body, w, s.Handlers)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
