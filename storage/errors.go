package storage

import (
	"fmt"
)

type sentinelError struct {
	message string
}

var (
	ErrDatabaseQuery   = sentinelError{"Error querying database"}
	ErrScanningRow     = sentinelError{"Error Scanning Row"}
	ErrIterating       = sentinelError{"Error when iterating sql rows"}
	ErrQueryScanOneRow = sentinelError{"Error when querying and scanning one row"}
)

func (e sentinelError) Error() string {
	return e.message
}

func WrapSQLError(sentinelErr sentinelError, err error) error {
	return fmt.Errorf("%v: %w", sentinelErr, err)
}
