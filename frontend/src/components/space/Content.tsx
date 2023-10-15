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
    const [content = [], updateContent ] = useState<Content[]>();

    const newContent = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)
        const contentName = form.get("name")

        if (userId > -1 && contentName) {
            Post('space', JSON.stringify({userId, contentName}))
            .then(res => {
                getContent()
            }).catch((err: Error) => {
                console.log('an error occurred', err);
            })
        }
    }

    const removeContent = (event: React.FormEvent<HTMLButtonElement>) => {
        event.preventDefault();
    }

    const getContent = () => {
        Get('content')
        .then(res => {
            const content = res as Content[]
            updateContent(content)
        })
        .catch(err => {
            console.log(err)
        })
    }

    useEffect(() => {
        getContent()
    }, [])

    return (
        <>
            <h1>Spaces</h1>

            <h2>Your spaces</h2>
            {
                content &&
                content.map((d) => (
                    <div key={d.id} style={{lineHeight: '35px', width: '400px', display: 'flex', justifyContent: 'space-between'}}>
                        <span>{d.contentKey}</span><span>{d.contentValue}</span>
                        <button style={{marginLeft: '15px', height: '24px', display: 'inline-block'}} onSubmit={removeContent} type="button">Remove</button>
                    </div>
                ))
            }
            <br />
            <h2>Add Content name and value</h2>
            <form onSubmit={newContent}>
                <label htmlFor="name">Content name</label>
                <input id="name" type="text" name="name" />
                <label htmlFor="name">Content value</label>
                <input id="value" type="text" name="value" />
                <button id="login-button" type="submit">Create</button>
            </form>

           
        </>
    )
}

export default Content