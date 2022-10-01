package handler

import (
	"context"

	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/mapper"
	"one.now/backend/repository"
)

type NoteService struct {
	pb.UnimplementedNoteServiceServer

	notes []*pb.Note
}

func (s NoteService) GetNoteList(ctx context.Context, req *pb.GetNoteListRequest) (*pb.GetNoteListResponse, error) {
	return &pb.GetNoteListResponse{
		Notes: s.notes,
	}, nil
}

func NewNoteService(dir string) NoteService {
	notes := repository.GetNotes(dir)

	a := make([]*pb.Note, len(notes))
	for i, n := range notes {
		a[i] = mapper.ToProto(n)
	}

	return NoteService{
		notes: a,
	}
}
