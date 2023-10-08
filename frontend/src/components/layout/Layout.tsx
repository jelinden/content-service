import React, { useEffect, useContext } from 'react';
import {
    Link,
    useNavigate
  } from "react-router-dom";
import { Props } from './LayoutTypeProps';
import { Get, Post } from '../service/http'
import { AppContext } from '../context/AppContext';

interface Profile {
    username: string
    apiToken: string
}

const Layout =({children} : Props) => {

    const { app, updateState } = useContext(AppContext);
    const navigate = useNavigate();

    const logout = (event: React.MouseEvent<HTMLElement>) => {
        event.preventDefault();
        Post('logout', "")
        .then(res => {
            if (app?.loggedIn !== false || app.username !== '') {
                updateState({app: {loggedIn: false, username: '', apiToken: ''}})
            }
            navigate("/");
        }).catch((err: Error) => {
            console.log('an error occurred', err);
        })
    }

    useEffect(() => {
        const fetchProfile = () => {
            Get('profile')
                .then(res => {
                    const p = res as Profile
                    if (app?.loggedIn !== true || app.username !== p.username) {
                        updateState({app: {loggedIn: true, username: p.username, apiToken: p.apiToken}})
                    }
                })
                .catch(err => {
                    console.log(err)
                })
        }

        fetchProfile()
      }, [app])

    return(
        <>
            <div>
                <ul>
                    <li className='links'><Link to="/">Home</Link></li>
                    { app && app.username &&
                        <>
                        <li className='links'><Link to="/profile">Profile</Link></li>
                        <li className='links'>Logged in as {app.username}</li>
                        <li className='links'><Link to="#" onClick={logout}>Logout</Link></li>
                        </>
                    }
                    { (!app || !app.username) &&
                        <>
                        <li className='links'><Link to="/register">Register</Link></li>
                        <li className='links'><Link to="/login">Login</Link></li>
                        </>
                    }
                </ul>
            </div>
            <main>{children}</main>
        </>
    )
}

export default Layout;