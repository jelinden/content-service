import React, { useState, useEffect, useContext } from 'react';
import { AppContext } from '../context/AppContext';
import { Get, Post } from '../service/http'

interface Content {
    id: number
    contentKey: string
    contentValue: string
}

const Content = () => {
    const { app } = useContext(AppContext);
    const userId = app?.id!!
    const [spaces = [], updateSpaces ] = useState<Content[]>();

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

    const removeSpace = (event: React.FormEvent<HTMLButtonElement>) => {
        event.preventDefault();
    }

    const getSpaces = () => {
        Get('spaces')
        .then(res => {
            const spaces = res as Content[]
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

            <h2>Your spaces</h2>
            {
                spaces &&
                spaces.map((d) => (
                    <div key={d.id} style={{lineHeight: '35px', width: '400px', display: 'flex', justifyContent: 'space-between'}}>
                        <span>{d.contentKey}</span>
                        <button style={{marginLeft: '15px', height: '24px', display: 'inline-block'}} onSubmit={removeSpace} type="button">Remove</button>
                    </div>
                ))
            }
            <br />
            <h2>Add a new Space</h2>
            <form onSubmit={newSpace}>
                <label htmlFor="name">Space name</label>
                <input id="name" type="text" name="name" />
                <button id="login-button" type="submit">Create</button>
            </form>

           
        </>
    )
}

export default Content