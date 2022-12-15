package handlernews

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSSlayout"
	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSStime"
)

func TestEnvelopeStructRFC1123Z(t *testing.T) {
	rss := RSSlayout.RSSRFC1123Z{
		Items: []RSSlayout.RSSChannelsRFC1123Z{
			RSSlayout.RSSChannelsRFC1123Z{
				Title:   "test1",
				Content: "testtesttesttetst",
				PubDate: RSStime.TimeRFC1123Z(time.Now()),
				Link:    "http://sfdasdf/asfdafsd",
			},
			RSSlayout.RSSChannelsRFC1123Z{
				Title:   "test2",
				Content: "testtesttesttetst",
				PubDate: RSStime.TimeRFC1123Z(time.Now()),
				Link:    "http://sfdasdf/asfdafsd",
			},
		},
	}

	t.Run("TestMarshalXML", func(t *testing.T) {
		bytes, _ := xml.Marshal(rss)
		t.Logf("%s", bytes)

	})

}

func TestEnvelopeStructRFC1123(t *testing.T) {
	rss := RSSlayout.RSSRFC1123{
		Items: []RSSlayout.RSSChannelsRFC1123{
			RSSlayout.RSSChannelsRFC1123{
				Title:   "test1",
				Content: "testtesttesttetst",
				PubDate: RSStime.TimeRFC1123(time.Now()),
				Link:    "http://sfdasdf/asfdafsd",
			},
			RSSlayout.RSSChannelsRFC1123{
				Title:   "test2",
				Content: "testtesttesttetst",
				PubDate: RSStime.TimeRFC1123(time.Now()),
				Link:    "http://sfdasdf/asfdafsd",
			},
		},
	}

	t.Run("TestMarshalXML", func(t *testing.T) {
		bytes, _ := xml.Marshal(rss)
		t.Logf("%s", bytes)

	})

}

func TestInitExtractNews(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name string
		args args
		want func() (RSSlayout.RSS, error)
	}{
		//TODO: Add test cases.
		{
			name: "Test_extractNews_1",
			args: args{
				ctx: context.Background(),
				url: "https://habr.com/ru/rss/hub/go/all/?fl=ru",
			},
		},
		{
			name: "Test_extractNews_2",
			args: args{
				ctx: context.Background(),
				url: "https://cprss.s3.amazonaws.com/golangweekly.com.xml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testfunc := InitExtractNews(tt.args.ctx, tt.args.url)

			got, err := testfunc()
			if len(got.GetItems()) == 0 {
				t.Errorf("InitExtractNews() = %v, err = %v", got, err)
			}
		})
	}
}
