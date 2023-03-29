import { Link } from 'react-router-dom';
import * as pb from "./gen/proto/note/v1/note";

const title = (data: string) => {
    return data.split('\n', 1).shift()?.substring(2);
}

type Props = {
    notes: pb.Note[]
}

function Home(props: Props) {
    return (
        <ul>
            {
                props.notes.map(n => {
                    return (
                        <li key={n.id.toString()}>
                            <Link to={"/note/" + n.id}>{title(n.body)}</Link>
                        </li>
                    );
                })
            }
        </ul>
    );
}

export default Home;
