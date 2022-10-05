package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"one.now/backend/controller"
	authv1 "one.now/backend/gen/proto/auth/v1"
	notev1 "one.now/backend/gen/proto/note/v1"
	"one.now/backend/handler"
	"one.now/backend/repository"
)

const (
	devOrigin = "https://localhost:3000"
	certFile  = "../cert/localhost.pem"
	keyFile   = "../cert/localhost-key.pem"
)

var (
	dir   = flag.String("dir", "", "The directory contains note files")
	email = flag.String("email", "", "Allowed email to login")
	dev   = flag.Bool("dev", false, "Run server in dev mode")
)

func main() {
	flag.Parse()

	gs := grpc.NewServer()

	notev1.RegisterNoteServiceServer(gs, handler.NewNoteService(controller.NewNoteCtrler(*dir)))
	authv1.RegisterAuthServiceServer(gs, handler.NewAuthService(controller.NewAuthCtrler(*email)))

	var originFunc func(string) bool
	if *dev {
		reflection.Register(gs)
		originFunc = func(origin string) bool { return origin == devOrigin }
	} else {
		originFunc = func(origin string) bool { return false }
	}

	wrappedServer := grpcweb.WrapServer(gs, grpcweb.WithOriginFunc(originFunc))

	http.Handle("/", handler.EnableSession(wrappedServer, repository.NewInMemorySession()))

	log.Println("Serving on https://localhost:3443")
	if err := http.ListenAndServeTLS("localhost:3443", certFile, keyFile, nil); err != nil {
		log.Fatal(err)
	}
}
