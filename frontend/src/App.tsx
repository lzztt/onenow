import './App.css';
import Home from './Home';
import Note from './Note';
import Navbar from './Navbar';
import { useEffect, useState } from "react";
import {
  BrowserRouter,
  Link,
  Navigate,
  Route,
  Routes,
  useLocation,
} from "react-router-dom";
import * as pb from "./gen/proto/note/v1/note";
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { authService, noteService } from './Gateway';
import Logout from './Logout';
import Login from './Login';
import Error from './Error';

declare global {
  interface Window {
    app: {
      gtagPageView: Function
    };
  }
}

const NotFound = <div className="err">404 Not Found :(</div>

type Props = {
  notes: pb.Note[]
  setLogin: (login: boolean) => void
}

function PageRouter(props: Props) {
  const location = useLocation();
  useEffect(() => {
    window.app?.gtagPageView();
  }, [location]);

  const sreNotes = props.notes.slice(-35, -29);
  const homeNotes = props.notes.slice(0, -35).concat(props.notes.slice(-29));

  return (
    <Routes>
      <Route path="/" element={<Home notes={homeNotes} />} />,
      <Route path="/note/:id" element={<Note notes={props.notes} />} />
      <Route path="/login" element={<Login setLogin={props.setLogin} />} />
      <Route path="/logout" element={<Logout setLogin={props.setLogin} />} />
      <Route path="/error" element={<Error />} />
      <Route path="/sre" element={<Home notes={sreNotes} />} />
      <Route path="*" element={NotFound} />
    </Routes>
  );
}

function App() {
  const [notes, setNotes] = useState<pb.Note[]>([]);
  const [error, setError] = useState<RpcError | undefined>();
  const [login, setLogin] = useState<boolean>(false);

  useEffect(() => {
    const getNotes = async () => {
      let data: pb.Note[] = [];

      if (process.env.NODE_ENV !== 'test') {
        try {
          const loginResp = await authService.login({ email: "" });
          if (loginResp.response.ok) {
            setLogin(true);
          }

          const noteResp = await noteService.getNoteList({});
          data = noteResp.response.notes;
        } catch (error) {
          setError(error as RpcError);
        }
      }

      setNotes(data);
    }

    getNotes();
  }, []);

  if (error !== undefined) {
    sessionStorage.setItem("error", (error as RpcError).toString());
    return <Navigate replace to="/error" />;
  }

  return (
    <BrowserRouter>
      <div>
        <h1>
          <Link to="/" title="onenow.life">
            <svg viewBox="0 0 400 80" xmlns="http://www.w3.org/2000/svg">
              <rect x="0" y="0" width="400" height="80" fill="#ffffff" />
              <text x="50%" y="60%" dominantBaseline="middle" textAnchor="middle" fontSize="60" fontFamily="Montserrat, sans-serif" fontWeight="500">
                <tspan fill="#5cb85c">one</tspan><tspan fill="#6495ed">now</tspan><tspan>.</tspan><tspan fill="#ff7f50">life</tspan>
              </text>
            </svg>
          </Link>
        </h1>
        <Navbar>
          <Link to="/sre">SRE</Link>
          <Link to={login ? "/logout" : "login"}>{login ? "Logout" : "Login"}</Link>
        </Navbar>
        <PageRouter notes={notes} setLogin={setLogin} />
      </div>
    </BrowserRouter>
  );
}

export default App;
