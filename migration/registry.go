package migration

import (
	"io"
	"os"
	"path"
)

const migrationsDir = "migration/migrations"

type Registry struct {
	currIdx      int
	migrationDir string
	migrations   []string
}

func (r *Registry) Migration() (io.Reader, error) {
	file, err := os.Open(path.Join(r.migrationDir, r.migrations[r.currIdx]))
	if err != nil {
		return nil, err
	}

	r.currIdx++

	return file, nil
}

func (r *Registry) Next() bool {
	return r.currIdx < len(r.migrations)
}

func NewRegistry() (*Registry, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	r := Registry{
		migrationDir: path.Join(curDir, migrationsDir),
		migrations: []string{
			"1.sql",
		},
	}

	return &r, nil
}
