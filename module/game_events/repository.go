package events

import (
	"context"
	"project-clickhouse/db"

	"github.com/pkg/errors"
)

type RepositoryInterface interface {
	CreateMany(ctx context.Context, events []Event) error
}

type Repository struct {
	db *db.Connection
}

func (r *Repository) CreateMany(ctx context.Context, events []Event) error {
	sqlStr := `INSERT INTO project_clickhouse.events (
		   client_time,
		   device_id, 
		   device_os,
		   session,
		   sequence,
		   param_int,
		   param_str,
		   event,
		   server_time,
		   ip
	   ) VALUES
	`
	vals := make([]interface{}, 0, len(events)*10)

	for _, event := range events {
		sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"

		vals = append(
			vals,
			event.ClientTime,
			event.DeviceID,
			event.DeviceOS,
			event.Session,
			event.Sequence,
			event.ParamInt,
			event.ParamStr,
			event.Event,
			event.ServerTime,
			event.IP,
		)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	stmt, err := r.db.PrepareContext(ctx, sqlStr)
	if err != nil {
		return errors.WithStack(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, vals...)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewGameEventRepository(db *db.Connection) RepositoryInterface {
	return &Repository{db: db}
}
