package RSSlayout

import (
	"encoding/xml"
	"time"

	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSStime"
)

type RSS interface {
	GetItems() []StructNews
}

type StructNews struct {
	Title   string
	Content string
	PubDate time.Time
	Link    string
	Channel int
}

func (sn *StructNews) SetChannel(id int) {
	sn.Channel = id
}

type InterfaceNews interface {
	ReturnStructNews() StructNews
}

//Психанул
type RSSChannelsBFTime struct {
	Title   string                 `xml:"title"`
	Content string                 `xml:"description"`
	PubDate RSStime.BruteForceTime `xml:"pubDate"`
	Link    string                 `xml:"link"`
}

func (r *RSSChannelsBFTime) ReturnStructNews() StructNews {
	return StructNews{Title: r.Title, Content: r.Content, PubDate: r.PubDate.ConvertToTime(), Link: r.Link}
}

type RSSBFTime struct {
	XMLName xml.Name            `xml:"rss"`
	Items   []RSSChannelsBFTime `xml:"channel>item"`
}

func (rss *RSSBFTime) GetItems() []StructNews {
	r := make([]StructNews, len(rss.Items))
	for count, i := range rss.Items {
		r[count] = i.ReturnStructNews()

	}
	return r
}

type RSSChannelsRFC1123 struct {
	Title   string              `xml:"title"`
	Content string              `xml:"description"`
	PubDate RSStime.TimeRFC1123 `xml:"pubDate"`
	Link    string              `xml:"link"`
}

func (r *RSSChannelsRFC1123) ReturnStructNews() (sn StructNews) {
	return StructNews{Title: r.Title, Content: r.Content, PubDate: time.Time(r.PubDate), Link: r.Link}
}

type RSSRFC1123 struct {
	XMLName xml.Name             `xml:"rss"`
	Items   []RSSChannelsRFC1123 `xml:"channel>item"`
}

func (rss *RSSRFC1123) GetItems() []StructNews {
	r := make([]StructNews, len(rss.Items))
	for count, i := range rss.Items {
		r[count] = i.ReturnStructNews()

	}
	return r
}

type RSSChannelsRFC1123Z struct {
	Title   string               `xml:"title"`
	Content string               `xml:"description"`
	PubDate RSStime.TimeRFC1123Z `xml:"pubDate"`
	Link    string               `xml:"link"`
}

func (rz *RSSChannelsRFC1123Z) ReturnStructNews() (sn StructNews) {
	return StructNews{Title: rz.Title, Content: rz.Content, PubDate: time.Time(rz.PubDate), Link: rz.Link}
}

type RSSRFC1123Z struct {
	XMLName xml.Name              `xml:"rss"`
	Items   []RSSChannelsRFC1123Z `xml:"channel>item"`
}

func (rss *RSSRFC1123Z) GetItems() []StructNews {
	r := make([]StructNews, len(rss.Items))
	for count, i := range rss.Items {
		r[count] = i.ReturnStructNews()

	}
	return r
}
