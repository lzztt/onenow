package mapper

import (
	"one.now/backend/entity"
	pb "one.now/backend/gen/proto/note/v1"
)

func ToProto(n *entity.Note) *pb.Note {
	return &pb.Note{
		Uuid: n.Id[:],
		Body: n.Body,
	}
}
