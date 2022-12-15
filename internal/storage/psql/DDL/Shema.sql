-- type Post struct {
--     ID      int    // номер записи
--     Title   string // заголовок публикации
--     Content string // содержание публикации
--     PubTime int64  // время публикации
--     Link    string // ссылка на источник
-- }


CREATE TABLE IF NOT EXISTS gonews(
    id SERIAL PRIMARY KEY,
    title text,
    content text NOT NULL,
    pubTime BIGINT,
    link text,
    idchannel int REFERENCES rssChanel(id) NOT NULL ON UPDATE CASCADE,
);

CREATE TABLE IF NOT EXISTS rssChanel(
    id SERIAL PRIMARY KEY,
    channel text UNIQUE
);

