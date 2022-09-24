import './App.css';
import Home from './Home';
import Note from './Note';
import { useEffect, useState } from "react";
import {
  BrowserRouter,
  Link,
  Route,
  Routes,
} from "react-router-dom";
import * as pb from "./gen/proto/note/v1/note";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { NoteServiceClient } from "./gen/proto/note/v1/note.client";

const NotFound = <div className="err">404 Not Found :(</div>

type Props = {
  notes: pb.Note[]
}

function PageRouter(props: Props) {
  return (
    <Routes>
      <Route path="/" element={<Home notes={props.notes} />} />,
      <Route path="/note/:id" element={<Note notes={props.notes} />} />
      <Route path="*" element={NotFound} />
    </Routes>
  );
}

const transport = new GrpcWebFetchTransport({
  baseUrl: process.env.REACT_APP_BACKEND as string,
});
const noteService = new NoteServiceClient(transport);

function App() {
  const [notes, setNotes] = useState<pb.Note[]>([]);
  const [error, setError] = useState<RpcError | undefined>();

  useEffect(() => {
    const getNotes = async () => {
      let data: pb.Note[] = [];

      if (process.env.NODE_ENV !== 'test') {
        try {
          const resp = await noteService.getNoteList({});
          data = resp.response.notes;
        } catch (error) {
          setError(error as RpcError);
        }
      }

      setNotes(data);
    }

    getNotes();
  }, []);

  if (error !== undefined) {
    return (
      <div>
        <h1>backend error: {error.message}</h1>
        <button onClick={() => setError(undefined)}>Clear error</button>
      </div>
    );
  }

  return (
    <BrowserRouter>
      <div>
        <h1><Link to="/">One Now</Link></h1>
        <PageRouter notes={notes} />
      </div>
    </BrowserRouter>
  );
}

export default App;
