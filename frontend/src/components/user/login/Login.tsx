import React, { useState, useContext } from 'react';
import { useNavigate } from "react-router-dom";
import { Post } from '../../service/http';
import { AppContext } from '../../context/AppContext';

const Login = () => {

    const [isLoading, setLoading] = useState(false);
    const [error, setError] = useState(false);
    const { updateState } = useContext(AppContext);

    const navigate = useNavigate();

    const loginSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)

        const username = form.get('username');
        const password = form.get('password');

        setLoading(true);
        Post('login', JSON.stringify({username, password}))
        .then(res => {
            setLoading(false);
            updateState({app: {loggedIn: true}})
            console.log('login succeeded', res);
            navigate("/");
        }).catch((err: Error) => {
            setLoading(false);
            setError(true)
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
                        <input type="password" name="password" id="password" />
                    </div>
                    <button id="login-button" type="submit">Login</button>
                </form>
                <br />
                {error && 
                    <div id="usernameError">Oops, login failed.</div>
                }
            </>
            }
        </>
    )
}

export default Login;