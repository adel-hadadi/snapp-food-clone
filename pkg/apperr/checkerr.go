package apperr

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const ErrSQLDuplicateEntryCode = "23505"

func IsSQLDuplicateEntry(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == ErrSQLDuplicateEntryCode
	}

	return false
}

func IsSQLNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
