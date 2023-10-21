import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext';


const Profile = () => {
    const { app, updateState } = useContext(AppContext);

    return (
        <>
            <h1>Profile</h1>

            <div>User id: <span>{app?.id}</span></div>
            <div>API token: <span>{app?.apiToken}</span></div>
        </>
    )
}

export default Profile