import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { NoteServiceClient } from "./gen/proto/note/v1/note.client";
import { AuthServiceClient } from "./gen/proto/auth/v1/auth.client";

const transport = new GrpcWebFetchTransport({
  baseUrl: process.env.REACT_APP_BACKEND as string,
});

export const noteService = new NoteServiceClient(transport);
export const authService = new AuthServiceClient(transport);
