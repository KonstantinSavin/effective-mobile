package sqldb

import (
	"database/sql"
	"effective-mobile/music-lib/internal/storage"
	"effective-mobile/music-lib/pkg/logging"

	_ "github.com/lib/pq"
)

type Storage struct {
	db      *sql.DB
	songRep *SongRep
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (st *Storage) Song() storage.SongRep {
	if st.songRep != nil {
		return st.songRep
	}

	st.songRep = &SongRep{
		storage: st,
		logger:  logging.GetLogger(),
	}

	return st.songRep
}
