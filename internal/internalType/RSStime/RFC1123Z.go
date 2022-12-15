package RSStime

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeRFC1123Z time.Time

func (d *TimeRFC1123Z) UnmarshalJSON(b []byte) (err error) {
	sDate := strings.Trim(string(b), `"`)
	td, err := time.Parse(time.RFC1123Z, sDate)
	*d = TimeRFC1123Z(td)
	if err != nil {
		return
	}
	return
}

func (d *TimeRFC1123Z) MarshalJSON() ([]byte, error) {
	if y := time.Time(*d).Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(fmt.Sprintf("%s", time.Time(*d).Format(time.RFC1123))), nil
}

func (dt *TimeRFC1123Z) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t string
	if err := d.DecodeElement(&t, &start); err != nil {
		return err
	}
	td, err := time.Parse(time.RFC1123Z, t)
	*dt = TimeRFC1123Z(td)
	if err != nil {
		return
	}
	return
}

func (d *TimeRFC1123Z) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if y := time.Time(*d).Year(); y < 0 || y >= 10000 {
		return errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	return e.EncodeElement([]byte(fmt.Sprintf("%s", time.Time(*d).Format(time.RFC1123Z))), start)
}

func (d *TimeRFC1123Z) ConvertToTime() time.Time {
	return time.Time(*d)
}
