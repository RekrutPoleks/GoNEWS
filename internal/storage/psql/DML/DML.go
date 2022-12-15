package DML

const (
	INSERT_NEWS    = `INSERT INTO gonews(title, content, pubTime, link, idchannel) VALUES ($1, $2, $3, $4, $5);`
	INSERT_CHANNEL = `INSERT INTO rsschannel(channel) VALUES ($1) RETURNING id`

	Query_NEWS_LIMIT     = `SELECT id, title, content, pubTime, link FROM gonews ORDER BY pubtime ASC LIMIT $1;`
	Query_GET_ID_CHANNEL = `SELECT id FROM rsschannel WHERE channel = $1;`
	Query_LastRecord     = `SELECT pubTime FROM gonews WHERE idchannel = $1 ORDER BY pubtime ASC LIMIT 1;`

	TEST_SELECT_ALL = `SELECT * FROM gonews;`
)
