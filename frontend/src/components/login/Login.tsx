import React, {useState} from 'react';
import Post from '../service/http';

const Login = () => {

    const [isLoading, setLoading] = useState(false);

    const loginSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)

        const username = form.get('username');
        const password = form.get('password');
        // TODO: validate
        setLoading(true);
        Post('login', JSON.stringify({username, password}))
        .then(res => {
            setLoading(false);
            console.log('login succeeded', res);

        }).catch((err: Error) => {
            setLoading(false);
            console.log('an error occurred', err);
        })
    }
    
    return (
        <>
            {isLoading &&
                <div className="spinner-container">
                    <div className="loading-spinner" />
                </div>
            }

            {!isLoading &&
            <>
                <h1>Login</h1>
                <form onSubmit={loginSubmit}>
                    <div>
                        <label htmlFor="username">Username/Email</label>
                        <input type="text" name="username" id="username"/>
                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input type="text" name="password" id="password"/>
                    </div>
                    <button id="login-button" type="submit">Login</button>
                </form>
            </>
            }
        </>
    )
}

export default Login;