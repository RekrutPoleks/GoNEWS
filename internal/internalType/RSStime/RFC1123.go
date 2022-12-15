package RSStime

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeRFC1123 time.Time

func (d *TimeRFC1123) UnmarshalJSON(b []byte) (err error) {
	sDate := strings.Trim(string(b), `"`)
	td, err := time.Parse(time.RFC1123, sDate)
	*d = TimeRFC1123(td)
	if err != nil {
		return
	}
	return
}

func (d *TimeRFC1123) MarshalJSON() ([]byte, error) {
	if y := time.Time(*d).Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(fmt.Sprintf("%s", time.Time(*d).Format(time.RFC1123))), nil
}

func (dt *TimeRFC1123) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {

	var t string
	if err := d.DecodeElement(&t, &start); err != nil {
		return err
	}
	td, err := time.Parse(time.RFC1123, t)
	*dt = TimeRFC1123(td)
	if err != nil {
		return
	}
	return
}

func (d *TimeRFC1123) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if y := time.Time(*d).Year(); y < 0 || y >= 10000 {
		return errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	s := time.Time(*d).Format(time.RFC1123)
	return e.EncodeElement([]byte(fmt.Sprintf("%s", s)), start)
}

func (d *TimeRFC1123) ConvertToTime() time.Time {
	return time.Time(*d)
}
