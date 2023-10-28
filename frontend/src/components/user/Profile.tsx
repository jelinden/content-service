import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext';


const Profile = () => {
    const { app, updateState } = useContext(AppContext);

    return (
        <>
            <h1>Profile</h1>

            <div>User id: <span>{app?.id}</span></div>
            <div>API token: <span>{app?.apiToken}</span></div>
            <div style={{marginTop: '20px'}}>
                <div>You can use the api token to get your content in a space, for example with curl.</div>
                
                <div style={{fontSize: '11px', padding: '10px', backgroundColor: '#F8F8F8', maxWidth: '700px', border: '1px solid blue', margin: '20px'}}>
                    curl "https://content-service.jelinden.fi/api/space/YOUR_SPACE_ID/entries?token=YOUR_TOKEN_REPLACED_HERE"
                </div>
                Replace YOUR_SPACE_ID and YOUR_TOKEN_REPLACED_HERE with your space id and your token.
                <div>(you can view the space id in your <a href="/space">Spaces</a> page, click in a Space)</div>
            </div>
        </>
    )
}

export default Profile