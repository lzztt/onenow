package handler

import (
	"context"

	"one.now/backend/controller"
	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/mapper"
)

type NoteService struct {
	pb.UnimplementedNoteServiceServer

	ctrler controller.NoteCtrler
}

func (s NoteService) GetNoteList(ctx context.Context, req *pb.GetNoteListRequest) (*pb.GetNoteListResponse, error) {
	return &pb.GetNoteListResponse{
		Notes: mapper.ToProtoArray(s.ctrler.GetNoteList(ctx)),
	}, nil
}

func NewNoteService(c controller.NoteCtrler) NoteService {
	return NoteService{
		ctrler: c,
	}
}
