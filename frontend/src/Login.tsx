import { SyntheticEvent, useState } from 'react';
import { RpcError } from "@protobuf-ts/runtime-rpc";
import { authService } from './Gateway';
import { Navigate } from 'react-router-dom';

type Props = {
    setLogin: (login: boolean) => void
}

function Login(props: Props) {
    const [email, setEmail] = useState<string>("");
    const [error, setError] = useState<RpcError | undefined>();

    const handleSubmit = async (event: SyntheticEvent) => {
        event.preventDefault();
        console.log(email)
        if (email.length > 0) {
            try {
                const resp = await authService.login({ email });
                if (resp.response.ok) {
                    console.log('Logged in');
                    props.setLogin(true);
                } else {
                    alert("Login Failed :(");
                }
            } catch (error) {
                setError(error as RpcError);
            }
        }
    };

    if (error !== undefined) {
        sessionStorage.setItem("error", (error as RpcError).toString());
        return <Navigate replace to="/error" />;
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Email:
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
            </label>
            <input type="submit" value="Submit" />
        </form>
    );
}

export default Login;
