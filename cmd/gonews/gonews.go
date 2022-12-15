package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RekrutPoleks/GoNEWS/api"
	handlernews "github.com/RekrutPoleks/GoNEWS/internal/handlerNews"
	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSSlayout"
	"github.com/RekrutPoleks/GoNEWS/internal/selector"
	"github.com/RekrutPoleks/GoNEWS/internal/storage/psql"
)

type Config struct {
	Period int `json:"request_period"`
	Rss    []string
}

func ReadConfig() *Config {
	file, err := ioutil.ReadFile("/home/future/go/src/github.com/RekrutPoleks/GoNEWS/cmd/gonews/config.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := Config{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func main() {

	ctx := context.Background()
	db, err := psql.NewDB(ctx, fmt.Sprintf("postgres://%s:%s@%s:5432/gonews", os.Getenv("USERNAME"), os.Getenv("PWDDB"), os.Getenv("IPDB")))
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(db)

	errChan := make(chan error, 4)
	quePutNews := make(chan []RSSlayout.StructNews)
	cfg := ReadConfig()
	go func() {
		for e := range errChan {
			log.Printf("--%s--\tError: %s\n", time.Now().Format(time.RFC822), e)

		}
	}()
	for _, url := range cfg.Rss {
		ext := handlernews.InitExtractNews(ctx, url)
		var starttime int64
		id, err := db.GetIdChannel(url)
		if err != nil {
			errChan <- err
		}
		if id == 0 {
			id, err = db.PutChannel(url)
			if err != nil {
				errChan <- err
				return
			}
			starttime = 0
		} else {
			starttime, _ = db.LastRecordbyID(id)

		}
		s := selector.InitSelectorAndPrepare(starttime, id)
		go func() {
			ticker := time.NewTicker(time.Duration(cfg.Period) * time.Minute)
			for _ = range ticker.C {
				rss, err := ext()
				if err != nil {
					errChan <- err
					continue
				}
				n := s(rss.GetItems())
				if len(n) > 0 {
					quePutNews <- n
				}
			}
		}()
	}

	go func() {
		for n := range quePutNews {
			err := db.PutNews(n)
			if err != nil {
				errChan <- err
				continue
			}
		}
	}()

	http.ListenAndServe(":5056", api.Router())

}
