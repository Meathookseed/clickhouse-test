package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CustomTime struct {
	time.Time
}

func (m *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// Get rid of the quotes "" around the value.
	s = s[1 : len(s)-1]

	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return errors.WithStack(err)
	}

	m.Time = t

	return
}

type EventDTO struct {
	ClientTime CustomTime `json:"client_time"`
	DeviceID   uuid.UUID  `json:"device_id"`
	DeviceOS   string     `json:"device_os"`
	Session    string     `json:"session"`
	Sequence   int64      `json:"sequence"`
	Event      string     `json:"event"`
	ParamInt   int16      `json:"param_int"`
	ParamStr   string     `json:"param_str"`
}
