package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	pb "one.now/backend/gen/proto/note/v1"
	"one.now/backend/handler"
)

var (
	dir = flag.String("dir", "", "The directory contains note files")
)

func main() {
	flag.Parse()

	gs := grpc.NewServer()
	pb.RegisterNoteServiceServer(gs, handler.New(*dir))
	wrappedServer := grpcweb.WrapServer(gs, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	http.Handle("/", wrappedServer)

	log.Println("Serving on http://0.0.0.0:3080")
	if err := http.ListenAndServe(":3080", nil); err != nil {
		log.Fatal(err)
	}
}
