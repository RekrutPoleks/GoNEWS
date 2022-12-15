package psql

import (
	"context"

	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSSlayout"
	"github.com/RekrutPoleks/GoNEWS/internal/storage"
	"github.com/RekrutPoleks/GoNEWS/internal/storage/psql/DML"
	"github.com/jackc/pgx/v5/pgxpool"
)

// urlExample := "postgres://username:password@localhost:5432/database_name"

type DB struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewDB(ctx context.Context, url string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DB{db, ctx}, nil
}

func (d *DB) PutNews(news []RSSlayout.StructNews) error {
	tx, err := d.db.Begin(d.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(d.ctx)
	pInsert, err := tx.Prepare(d.ctx, "putNews", DML.INSERT_NEWS)
	if err != nil {
		return nil
	}
	for _, n := range news {
		_, err = tx.Exec(d.ctx, pInsert.Name, n.Title, n.Content, n.PubDate.Unix(), n.Link, n.Channel)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(d.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetNews(limit int) ([]storage.Post, error) {
	rows, err := d.db.Query(d.ctx, DML.Query_NEWS_LIMIT, limit)
	if err != nil {
		return nil, err
	}

	r := storage.Post{}
	news := make([]storage.Post, 0)
	for rows.Next() {
		if err = rows.Scan(&r.ID, &r.Title, &r.Content, &r.PubTime, &r.Link); err != nil {
			return nil, err
		}
		news = append(news, r)
	}
	return news, nil
}

func (d *DB) GetIdChannel(channel string) (id int, err error) {
	qr := d.db.QueryRow(d.ctx, DML.Query_GET_ID_CHANNEL, &channel)
	if err = qr.Scan(&id); err != nil {
		return
	}
	return
}

func (d *DB) PutChannel(channel string) (id int, err error) {
	qr := d.db.QueryRow(d.ctx, DML.INSERT_CHANNEL, &channel)
	if err = qr.Scan(&id); err != nil {
		return
	}
	return
}

func (d *DB) LastRecordbyID(id int) (time int64, err error) {
	qr := d.db.QueryRow(d.ctx, DML.Query_LastRecord, &id)
	if err = qr.Scan(&time); err != nil {
		return
	}
	return
}
