import React, { useEffect, useContext } from 'react';
import {
    Link,
    useNavigate
  } from "react-router-dom";
import { Props } from './LayoutTypeProps';
import { Get, Post } from '../service/http'
import { AppContext } from '../context/AppContext';

export interface Profile {
    id: number
    username: string
    apiToken: string
}

const Layout = ({children} : Props) => {

    const { app, updateState } = useContext(AppContext);
    const navigate = useNavigate();

    const logout = (event: React.MouseEvent<HTMLElement>) => {
        event.preventDefault();
        Post('logout', "")
        .then(res => {
            if (app?.loggedIn !== false || app.username !== '') {
                updateState({app: {id: -1, loggedIn: false, username: '', apiToken: ''}})
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
                        updateState({app: {id: p.id, loggedIn: true, username: p.username, apiToken: p.apiToken}})
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
            <div style={{borderBottom: '1px solid grey', backgroundColor: '#D7E4FF', paddingTop:'1px'}}>
                <ul>
                    <li className='links'><Link to="/">Home</Link></li>
                    { app && app.username &&
                        <>
                        <li className='links'><Link to="/profile">Profile</Link></li>
                        <li className='links'><Link to="/space">Spaces</Link></li>
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