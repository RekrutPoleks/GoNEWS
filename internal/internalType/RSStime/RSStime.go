package RSStime

import (
	"encoding/xml"
	"time"
)

type RSStime interface {
	UnmarshalJSON(b []byte) (err error)
	MarshalJSON() ([]byte, error)
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error)
	MarshalXML(e *xml.Encoder, start xml.StartElement)
	ConvertToTime() time.Time
}
