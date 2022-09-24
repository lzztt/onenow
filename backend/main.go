package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/note"
)

//go:embed gen/note
var gen embed.FS

func getNotes() []*pb.Note {
	noteFS, err := fs.Sub(gen, "gen/note")
	if err != nil {
		log.Fatal(err)
	}

	matches, err := fs.Glob(noteFS, "*_*.md")
	if err != nil {
		log.Fatal(err)
	}

	notes := make([]*pb.Note, len(matches))

	for _, file := range matches {
		i, err := strconv.Atoi(strings.SplitN(file, "_", 2)[0])
		if err != nil {
			log.Fatal(err)
		}

		data, err := fs.ReadFile(noteFS, file)
		if err != nil {
			log.Fatal(err)
		}

		notes[i-1] = &pb.Note{
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
