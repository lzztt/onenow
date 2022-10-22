package main

import (
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

func main() {
	c := getConfig()

	gs := grpc.NewServer()

	notev1.RegisterNoteServiceServer(gs, handler.NewNoteService(controller.NewNoteCtrler(c.Data.NoteDir)))
	authv1.RegisterAuthServiceServer(gs, handler.NewAuthService(controller.NewAuthCtrler(c.Secret.AllowedEmail)))

	wrappedServer := grpcweb.WrapServer(gs)

	s := repository.NewInMemorySession()
	http.Handle("/", handler.EnableSession(wrappedServer, &s))

	log.Println("Serving on https://localhost:" + c.Server.Port)
	if err := http.ListenAndServeTLS("localhost:"+c.Server.Port, c.Server.CertFile, c.Server.KeyFile, nil); err != nil {
		log.Fatal(err)
	}
}
