import React, {useState} from 'react';
import { Link, useNavigate } from "react-router-dom";
import Post from '../service/http';
import validateFields from './ValidateFields';

const Register = () => {

    const [isLoading, setLoading] = useState(false);
    const [error, setError] = useState<Error>();
    const navigate = useNavigate();

    const registerSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)

        const username = form.get('username')?.toString();
        const password = form.get('password')?.toString();
        var usernameError = document.getElementById('usernameError') as HTMLElement
        var passwordError = document.getElementById('passwordError') as HTMLElement
        const validationErrors = validateFields(username ? username : '', password ? password : '')
        
        if (validationErrors.username !== '' || validationErrors.password !== '') {
            usernameError.innerHTML = validationErrors.username
            passwordError.innerHTML = validationErrors.password
            return
        }

        setLoading(true);
        Post('register', JSON.stringify({username, password}))
        .then(res => {
            setLoading(false);
            navigate("/login");

        }).catch((err: Error) => {
            setLoading(false);
            setError(err)
            console.log('an error occurred', err);
        })
    }
    
    return (
        <>
            {isLoading && !error &&
                <div className="spinner-container">
                    <div className="loading-spinner" />
                </div>
            }

            {!isLoading && !error &&
            <>
                <h1>Register</h1>
                <form onSubmit={registerSubmit}>
                    <div>
                        <label htmlFor="username" id="usernameError"></label>
                        <label htmlFor="username">Username/Email</label>
                        <input type="text" name="username" id="username" required/>
                    </div>
                    <div>
                        <label htmlFor="password" id="passwordError"></label>
                        <label htmlFor="password">Password</label>
                        <input 
                            type="text" 
                            name="password" 
                            id="password"
                            required
                        />
                    </div>
                    <button id="register-button" type="submit">Register</button>
                </form>
            </>
            }

            {!isLoading && error &&
                <div className="errors">
                    <div>Oops, an error occurred. Please try again.</div>
                    <br />
                    <Link to="/register" onClick={() => window.location.reload()}>Back to register page</Link>
                </div>
            }
        </>
    )
}

export default Register;
