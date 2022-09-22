package note

import (
	"context"

	pb "one.now/backend/gen/proto/note/v1"
)

type NoteService struct {
	pb.UnimplementedNoteServiceServer

	Notes []*pb.Note
}

func (s NoteService) GetNoteList(ctx context.Context, req *pb.GetNoteListRequest) (*pb.GetNoteListResponse, error) {
	return &pb.GetNoteListResponse{
		Notes: s.Notes,
	}, nil
}
