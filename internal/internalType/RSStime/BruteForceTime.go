package RSStime

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"
)

type BruteForceTime struct {
	layout string
	time.Time
}

func NewBruteForceTime(layout string) *BruteForceTime {
	return &BruteForceTime{layout: layout}
}

func (bt *BruteForceTime) UnmarshalJSON(b []byte) (err error) {
	lyaouts := []string{
		`Mon, 2 Jan 2006 15:04:05 -0700`,
		"Mon Jan _2 15:04:05 2006",
		"Mon Jan _2 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"02 Jan 06 15:04 MST",
		"02 Jan 06 15:04 -0700", // RFC822 with numeric zone
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700", // RFC1123 with numeric zone
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"3:04PM",
		// Handy time stamps.
		"Jan _2 15:04:05",
		"Jan _2 15:04:05.000",
		"Jan _2 15:04:05.000000",
		"Jan _2 15:04:05.000000000",
	}

	sDate := strings.Trim(string(b), `"`)
	for _, l := range lyaouts {
		bt.Time, err = time.Parse(l, sDate)
		if err == nil {
			bt.layout = l
			return
		}
	}

	return errors.New(fmt.Sprintf("%s unable parse time brute force", sDate))

}

func (bt *BruteForceTime) MarshalJSON() ([]byte, error) {
	if y := bt.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	if bt.layout == "" {
		bt.layout = time.RFC1123
	}
	return []byte(fmt.Sprintf("%s", bt.Format(bt.layout))), nil
}

func (bt *BruteForceTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	lyaouts := []string{
		`Mon, 2 Jan 2006 15:04:05 -0700`,
		"Mon Jan _2 15:04:05 2006",
		"Mon Jan _2 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"02 Jan 06 15:04 MST",
		"02 Jan 06 15:04 -0700",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"3:04PM",
		"Jan _2 15:04:05",
		"Jan _2 15:04:05.000",
		"Jan _2 15:04:05.000000",
		"Jan _2 15:04:05.000000000",
	}
	var t string
	if err := d.DecodeElement(&t, &start); err != nil {
		return err
	}
	for _, l := range lyaouts {
		bt.Time, err = time.Parse(l, t)
		if err == nil {
			bt.layout = l
			return
		}
	}
	return errors.New(fmt.Sprintf("%s unable parse xml time brute force", t))
}

func (bt *BruteForceTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if y := bt.Time.Year(); y < 0 || y >= 10000 {
		return errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	if bt.layout == "" {
		bt.layout = time.RFC1123
	}
	s := bt.Time.Format(bt.layout)
	return e.EncodeElement([]byte(fmt.Sprintf("%s", s)), start)
}

func (bt *BruteForceTime) ConvertToTime() time.Time {
	return bt.Time
}
