package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ClientTime time.Time
	DeviceID   uuid.UUID
	DeviceOS   string
	Session    string
	Sequence   int64
	Event      string
	ParamInt   int16
	ParamStr   string
	IP         string
	ServerTime time.Time
}

func dtoEventToModelEvent(dto EventDTO, ip string) Event {
	return Event{
		ClientTime: dto.ClientTime.Time,
		DeviceID:   dto.DeviceID,
		DeviceOS:   dto.DeviceOS,
		Session:    dto.Session,
		Sequence:   dto.Sequence,
		Event:      dto.Event,
		ParamInt:   dto.ParamInt,
		ParamStr:   dto.ParamStr,
		ServerTime: time.Now(),
		IP:         ip,
	}
}
