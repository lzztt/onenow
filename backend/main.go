package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	authv1 "one.now/backend/gen/proto/auth/v1"
	notev1 "one.now/backend/gen/proto/note/v1"
	"one.now/backend/handler"
)

var (
	dir   = flag.String("dir", "", "The directory contains note files")
	email = flag.String("email", "", "Allowed email to login")
)

func main() {
	flag.Parse()

	gs := grpc.NewServer()

	notev1.RegisterNoteServiceServer(gs, handler.NewNoteService(*dir))
	authv1.RegisterAuthServiceServer(gs, handler.NewAuthService(*email))

	wrappedServer := grpcweb.WrapServer(gs, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	http.Handle("/", handler.EnableSession(wrappedServer))

	log.Println("Serving on http://0.0.0.0:3080")
	if err := http.ListenAndServe(":3080", nil); err != nil {
		log.Fatal(err)
	}
}
