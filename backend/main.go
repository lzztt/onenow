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

//go:embed gen/note gen/build
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

	distFS, err := fs.Sub(gen, "gen/build")
	if err != nil {
		log.Fatal(err)
	}

	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedServer.IsAcceptableGrpcCorsRequest(req) || wrappedServer.IsGrpcWebRequest(req) {
			wrappedServer.ServeHTTP(resp, req)
			return
		}

		http.FileServer(http.FS(distFS)).ServeHTTP(resp, req)
	})

	http.Handle("/", handler)

	log.Println("Serving on http://0.0.0.0:3080")
	if err := http.ListenAndServe(":3080", nil); err != nil {
		log.Fatal(err)
	}
}
