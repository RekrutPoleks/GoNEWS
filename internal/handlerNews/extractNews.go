package handlernews

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSSlayout"
	"io"
	"net/http"
	"time"
)

func InitExtractNews(ctx context.Context, url string) func() (RSSlayout.RSS, error) {
	count := 0
	listlayout := []RSSlayout.RSS{
		&RSSlayout.RSSRFC1123{},
		&RSSlayout.RSSRFC1123Z{},
		&RSSlayout.RSSBFTime{},
	}

	changelayout := func() (RSSlayout.RSS, error) {
		if count > len(listlayout)-1 {
			return nil, fmt.Errorf("Unable to parse RSS")
		}
		i := listlayout[count]
		count++
		return i, nil
	}

	rss, _ := changelayout()

	return func() (RSSlayout.RSS, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		client := http.Client{Timeout: time.Duration(10 * time.Minute)}
		resp, err := client.Do(req)

		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Request status code %d", resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = xml.Unmarshal(body, &rss)
		if err != nil {
			var t *time.ParseError
			if errors.As(err, &t) {
				for {
					rss, err = changelayout()
					if err != nil {
						return nil, err
					}
					err = xml.Unmarshal(body, &rss)
					if err == nil {
						break
					}
				}
			}

		}
		return rss, nil
	}

}
