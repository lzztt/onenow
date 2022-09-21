import { Link } from 'react-router-dom';
import notes from './gen/notes.json';

const title = (data: string) => {
    return data.split('\n', 1).shift()?.substring(2);
}

function Home() {
    return (
        <ul>
            {
                notes.map((n, i) => (
                    <li key={i}>
                        <Link to={"/note/" + i}>{title(n.data)}</Link>
                    </li>
                ))
            }
        </ul>
    );
}

export default Home;