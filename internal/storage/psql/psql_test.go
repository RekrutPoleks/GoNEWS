package psql

import (
	"context"
	"fmt"
	"os"
	"testing"

	handlernews "github.com/RekrutPoleks/GoNEWS/internal/handlerNews"
	"github.com/RekrutPoleks/GoNEWS/internal/selector"
	"github.com/RekrutPoleks/GoNEWS/internal/storage/psql/DDL"
	"github.com/RekrutPoleks/GoNEWS/internal/storage/psql/DML"
)

func TestDB_DBNews(t *testing.T) {
	ctx := context.Background()

	db, err := NewDB(ctx, fmt.Sprintf("postgres://%s:%s@%s:5432/gonews", os.Getenv("USERNAME"), os.Getenv("PWDDB"), os.Getenv("IPDB")))
	if err != nil {
		t.Errorf("%v", err)
	}
	handl := handlernews.InitExtractNews(ctx, "https://habr.com/ru/rss/hub/go/all/?fl=ru")
	_, err = db.db.Exec(ctx, DDL.TESTTruncateTableRSSChannel)
	id, err := db.PutChannel("https://habr.com/ru/rss/hub/go/all/?fl=ru")
	if err != nil {
		t.Errorf("Error PutChannel Test Fault")
	}
	sel := selector.InitSelectorAndPrepare(0, id)
	if err != nil {
		t.Errorf("%s :line 23", err)
	}
	rss, err := handl()
	news := rss.GetItems()
	news = sel(news)

	if err != nil {
		t.Errorf("%s", err)
	}

	db.PutNews(news)

	countRecord := func() int {
		lenQuery := 0
		rows, err := db.db.Query(ctx, DML.TEST_SELECT_ALL)
		if err != nil {
			t.Error(err)

		}

		for rows.Next() {
			lenQuery++

		}
		return lenQuery
	}()

	if len(news) != countRecord {
		t.Errorf(" Count record news  got= %d, want %d", len(news), countRecord)
	}

}
