import React, { useState, useEffect, useContext } from 'react'
import { Link } from "react-router-dom"
import { AppContext } from '../context/AppContext'
import { Get, Post, Delete } from '../service/http'

interface Space {
    id: number
    spaceName: string
}

const Spaces = () => {
    const { app } = useContext(AppContext);
    const userId = app?.id!!
    const [ spaces = [], updateSpaces ] = useState<Space[]>();

    const newSpace = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)
        const spaceName = form.get("name")
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
        const id = (event.currentTarget as HTMLButtonElement).value
        Delete('space', parseInt(id, 10))
            .then(res => {
                getSpaces()
        })
        .catch(err => {
            console.log(err)
        })
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

            <h2>Your spaces</h2>
            {
                spaces &&
                spaces.map((d) => (
                    <div key={d.id} style={{lineHeight: '35px', width: '400px', display: 'flex', justifyContent: 'space-between'}}>
                        <Link to={`/content/${d.id}`}><span>{d.spaceName}</span></Link>
                        <button value={d.id} 
                            style={{marginLeft: '15px', height: '24px', display: 'inline-block', marginTop: '10px'}} 
                            onClick={removeSpace} 
                            type="button">Remove</button>
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

export default Spaces