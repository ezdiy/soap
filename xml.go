package soap

import "encoding/xml"

type Values = map[string]string

// Individual values listed as call arguments, or return values in response
type Value struct {
	XMLName xml.Name
	Content string `xml:",innerxml"`
}

// The entire SOAP body, both request and response
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body struct {
		Item struct {
			XMLName xml.Name
			Content string `xml:",innerxml"`
			Values []Value
		} `xml:",any"`
	}
}

// Unmarshal the request/response into a map of values
func (b *Envelope) Unmarshal() (ret Values) {
	ret = Values{}
	for _, i := range b.Body.Item.Values {
		ret[i.XMLName.Local] = i.Content
	}
	return
}

// Marshal map of values into a request/response
func (b *Envelope) Marshal(values Values) {
	var vl []Value
	for k, v := range values {
		 vl = append(vl, Value{
		 	XMLName: xml.Name{
		 		Local: k,
			},
			Content: v,
		 })
	}
	b.Body.Item.Values = vl
}
