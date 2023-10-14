import React, { useState, useEffect, useContext } from 'react';
import { AppContext } from '../context/AppContext';
import { Get, Post } from '../service/http'

interface Space {
    id: number
    spaceName: string
}

const Spaces = () => {
    const { app } = useContext(AppContext);
    const userId = app?.id!!
    const [spaces = [], updateSpaces ] = useState<Space[]>();

    const newSpace = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)
        const spaceName = form.get("name")
        console.log('name', spaceName, userId)
        if (userId > -1 && spaceName) {
            Post('space', JSON.stringify({userId, spaceName}))
            .then(res => {
                getSpaces()
            }).catch((err: Error) => {
                console.log('an error occurred', err);
            })
        }
    }

    const getSpaces = () => {
        Get('spaces')
        .then(res => {
            const spaces = res as Space[]
            updateSpaces(spaces)
        })
        .catch(err => {
            console.log(err)
        })
    }

    useEffect(() => {
        getSpaces()
    }, [])

    return (
        <>
            <h1>Spaces</h1>

            <div>Your spaces:</div>
            {
                spaces &&
                spaces.map((d) => (
                    <div key={d.id}><span>{d.spaceName}</span></div>
                ))
            }
            <div>Add a new Space</div>
            <form onSubmit={newSpace}>
                <label htmlFor="name">Space name</label>
                <input id="name" type="text" name="name" />
                <button id="login-button" type="submit">Create</button>
            </form>

           
        </>
    )
}

export default Spaces