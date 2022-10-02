import { SyntheticEvent, useState } from 'react';
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { authService } from './Gateway';
import { Navigate } from 'react-router-dom';

type Props = {
    setLogin: (login: boolean) => void
}

function Logout(props: Props) {
    const [error, setError] = useState<RpcError | undefined>();

    const handleSubmit = async (event: SyntheticEvent) => {
        event.preventDefault();
        try {
            await authService.logout({});
            props.setLogin(false);
            console.log('Logged out');
        } catch (error) {
            setError(error as RpcError);
        }
    };

    if (error !== undefined) {
        sessionStorage.setItem("error", (error as RpcError).toString());
        return <Navigate replace to="/error" />;
    }

    return (
        <form onSubmit={handleSubmit}>
            <input type="submit" value="Submit" />
        </form>
    );
}

export default Logout;
