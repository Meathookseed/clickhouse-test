package events

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scizorman/go-ndjson"
)

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository: repository}
}

func (s *Service) Deserialize(data []byte) ([]EventDTO, error) {
	var events []EventDTO

	err := ndjson.Unmarshal(data, &events)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return events, nil
}

func (s *Service) CreateEvents(ctx context.Context, events []EventDTO, ip string) error {
	dbEvents := make([]Event, len(events))
	for i, e := range events {
		dbEvents[i] = dtoEventToModelEvent(e, ip)
	}

	return s.repository.CreateMany(ctx, dbEvents)
}
