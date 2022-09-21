import { marked } from "marked";
import { Navigate, useParams } from 'react-router-dom';
import notes from './gen/notes.json';

function Note() {
    const params = useParams();
    const i = params.id ? parseInt(params.id) : -1;

    if (i < 0 || i >= notes.length) {
        return <Navigate replace to="/404" />;
    }

    return (
        <p dangerouslySetInnerHTML={{ __html: marked(notes[i].data) }} />
    );
}

export default Note;