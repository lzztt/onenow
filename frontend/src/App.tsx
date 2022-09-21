import './App.css';
import Home from './Home';
import Note from './Note';
import {
  BrowserRouter,
  Link,
  Route,
  Routes,
} from "react-router-dom";

const NotFound = <div className="err">404 Not Found :(</div>

function PageRouter() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />,
      <Route path="/note/:id" element={<Note />} />
      <Route path="*" element={NotFound} />
    </Routes>
  );
}

function App() {
  return (
    <BrowserRouter>
      <div>
        <h1><Link to="/">One Now</Link></h1>
        <PageRouter />
      </div>
    </BrowserRouter>
  );
}

export default App;
