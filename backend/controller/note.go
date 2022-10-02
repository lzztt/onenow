package controller

import (
	"context"

	"one.now/backend/entity"
	"one.now/backend/repository"
)

type NoteCtrler struct {
	notes []*entity.Note
}

func (c NoteCtrler) GetNoteList(ctx context.Context) []*entity.Note {
	return c.notes
}

func NewNoteCtrler(dir string) NoteCtrler {
	return NoteCtrler{
		notes: repository.GetNotes(dir),
	}
}
