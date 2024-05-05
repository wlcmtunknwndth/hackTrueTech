package postgres

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/hackBPA/internal/storage"
	"time"
)

func (s *Storage) GetEvent(ctx context.Context, id uint) (*storage.Event, error) {
	const op = "storage.postgres.events.GetEvent"

	var event storage.Event

	err := s.driver.QueryRowContext(ctx, getEvent, id).Scan(
		event.Id, event.Price, event.Restrictions, event.Date,
		event.Location, event.Name, event.ImgPath, event.Description,

		event.Features.Disability, event.Features.Deaf, event.Features.Blind,
		event.Features.Neural,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &event, nil
}

func (s *Storage) getId(name string, date time.Time) (uint, error) {
	const op = "storage.postgres.events.getId"

	var id uint
	if err := s.driver.QueryRow(getId, name, date).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) CreateEvent(ctx context.Context, event *storage.Event) (uint, error) {
	const op = "storage.postgres.events.CreateEvent"

	_, err := s.driver.ExecContext(ctx, createEvent, &event.Price,
		&event.Restrictions, &event.Date, &event.Location,
		&event.Name, &event.ImgPath, &event.Description,

		&event.Features.Disability, &event.Features.Deaf,
		&event.Features.Blind, &event.Features.Neural,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id uint
	if id, err = s.getId(event.Name, event.Date); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) DeleteEvent(ctx context.Context, id uint) error {
	const op = "storage.postgres.events.DeleteEvent"

	if _, err := s.driver.ExecContext(ctx, deleteEvent, id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
