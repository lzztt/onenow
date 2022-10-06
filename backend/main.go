package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"

	"one.now/backend/controller"
	authv1 "one.now/backend/gen/proto/auth/v1"
	notev1 "one.now/backend/gen/proto/note/v1"
	"one.now/backend/handler"
	"one.now/backend/repository"
)

const (
	certFile = "../cert/localhost.pem"
	keyFile  = "../cert/localhost-key.pem"
)

var (
	dir   = flag.String("dir", "", "The directory contains note files")
	email = flag.String("email", "", "Allowed email to login")
)

func main() {
	flag.Parse()

	gs := grpc.NewServer()

	notev1.RegisterNoteServiceServer(gs, handler.NewNoteService(controller.NewNoteCtrler(*dir)))
	authv1.RegisterAuthServiceServer(gs, handler.NewAuthService(controller.NewAuthCtrler(*email)))

	wrappedServer := grpcweb.WrapServer(gs)

	s := repository.NewInMemorySession()
	http.Handle("/", handler.EnableSession(wrappedServer, &s))

	log.Println("Serving on https://localhost:3443")
	if err := http.ListenAndServeTLS("localhost:3443", certFile, keyFile, nil); err != nil {
		log.Fatal(err)
	}
}
