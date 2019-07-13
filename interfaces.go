// Copyright 2019 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package dbq

import (
	"context"
	stdSql "database/sql"
)

// ExecContexter is for modifying the database state.
type ExecContexter interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (stdSql.Result, error)
}

// QueryContexter is for querying the database.
type QueryContexter interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*stdSql.Rows, error)
}

// SQLBasic allows for querying and executing statements.
type SQLBasic interface {
	ExecContexter
	QueryContexter
}
