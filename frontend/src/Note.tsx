import { marked } from "marked";
import { Navigate, useParams } from 'react-router-dom';
import * as pb from "./gen/proto/note/v1/note";

type Props = {
    notes: pb.Note[]
}

function Note(props: Props) {
    const params = useParams();

    if (props.notes.length === 0) {
        return <></>;
    }

    let body: string | null = null;
    props.notes.some(n => {
        if (n.id.toString() === params.id) {
            body = n.body
            return true;
        }

        return false;
    })

    if (body === null) {
        return <Navigate replace to="/404" />;
    }

    return (
        <p dangerouslySetInnerHTML={{ __html: marked(body) as string }} />
    );
}

export default Note;
