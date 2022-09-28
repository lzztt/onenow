package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/note"
)

type Metadata struct {
	createTime     int
	lastUpdateTime int
	id             uuid.UUID
	file           string
}

func loadMetadata(files []string) []*Metadata {
	m := make(map[uuid.UUID]*Metadata)

	for _, file := range files {
		t := strings.Split(file, "/")
		p := strings.SplitN(t[len(t)-1], "_", 2)
		id := uuid.MustParse(p[0])
		time, err := strconv.Atoi(p[1][:len(p[1])-3])
		if err != nil {
			log.Fatal(err)
		}

		if data, ok := m[id]; ok {
			if data.createTime > time {
				data.createTime = time
			}

			if data.lastUpdateTime < time {
				data.lastUpdateTime = time
				data.file = file
			}
		} else {
			m[id] = &Metadata{
				createTime:     time,
				lastUpdateTime: time,
				id:             id,
				file:           file,
			}
		}
	}

	v := make([]*Metadata, 0, len(m))

	for _, val := range m {
		v = append(v, val)
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i].createTime < v[j].createTime
	})

	return v
}

func getNotes() []*pb.Note {
	flag.Parse()
	files := flag.Args()

	metadata := loadMetadata(files)

	notes := make([]*pb.Note, len(metadata))
	for i, v := range metadata {
		id := v.id
		data, err := os.ReadFile(v.file)
		if err != nil {
			log.Fatal(err)
		}

		notes[i] = &pb.Note{
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
