import React, {useState} from 'react';
import Post from '../service/http';

const Register = () => {

    const [isLoading, setLoading] = useState(false);

    const registerSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)

        const username = form.get('username');
        const password = form.get('password');
        // TODO: validate
        setLoading(true);
        Post('register', JSON.stringify({username, password}))
        .then(res => {
            setLoading(false);
            console.log('registering succeeded', res);

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
                <h1>Register</h1>
                <form onSubmit={registerSubmit}>
                    <div>
                        <label htmlFor="username">Username/Email</label>
                        <input type="text" name="username" id="username"/>
                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input type="text" name="password" id="password"/>
                    </div>
                    <button id="register-button" type="submit">Register</button>
                </form>
            </>
            }
        </>
    )
}

export default Register;