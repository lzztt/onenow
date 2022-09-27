package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/note"
)

func getNotes() []*pb.Note {
	flag.Parse()
	files := flag.Args()
	notes := make([]*pb.Note, len(files))

	for _, file := range files {
		t := strings.Split(file, "/")
		p := strings.SplitN(t[len(t)-1], "_", 2)
		i, err := strconv.Atoi(p[0])
		if err != nil {
			log.Fatal(err)
		}
		id := uuid.MustParse(p[1][:len(p[1])-3])

		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		notes[i-1] = &pb.Note{
			Uuid: id[:],
			Body: string(data),
		}
	}

	return notes
}

func main() {
	gs := grpc.NewServer()
	pb.RegisterNoteServiceServer(gs, note.NoteService{Notes: getNotes()})
	wrappedServer := grpcweb.WrapServer(gs, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	http.Handle("/", wrappedServer)

	log.Println("Serving on http://0.0.0.0:3080")
	if err := http.ListenAndServe(":3080", nil); err != nil {
		log.Fatal(err)
	}
}
