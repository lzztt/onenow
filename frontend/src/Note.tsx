import { marked } from "marked";
import { Navigate, useParams } from 'react-router-dom';
import * as pb from "./gen/proto/note/v1/note";

type Props = {
    notes: pb.Note[]
}

function Note(props: Props) {
    const params = useParams();
    const i = params.id ? parseInt(params.id) : -1;

    if (i < 0 || i >= props.notes.length) {
        return <Navigate replace to="/404" />;
    }

    return (
        <p dangerouslySetInnerHTML={{ __html: marked(props.notes[i].body) }} />
    );
}

export default Note;