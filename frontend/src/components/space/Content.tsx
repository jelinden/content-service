import React, { useState, useEffect, useContext } from 'react';
import { useParams, useNavigate } from 'react-router';
import { AppContext } from '../context/AppContext';
import { Get, Post, Delete } from '../service/http'

interface Content {
    id: number
    name: string
    value: string
}

const Content = () => {
    const params = useParams()
    const { app } = useContext(AppContext);
    const navigate = useNavigate();
    const userId = app?.id!!
    const [ content = [], updateContent ] = useState<Content[]>();
    const spaceID = Number(params.spaceID)
    

    const newContent = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = new FormData(event.currentTarget)
        const contentName = form.get("name")
        const contentValue = form.get("value")

        if (userId > -1 && contentName) {
            Post('content', JSON.stringify({spaceID, name: contentName, value: contentValue}))
            .then(res => {
                getContent()
            }).catch((err: Error) => {
                console.log('an error occurred', err);
            })
        }
    }

    const removeContent = (event: React.FormEvent<HTMLButtonElement>) => {
        event.preventDefault();
        const id = (event.currentTarget as HTMLButtonElement).value
        Delete('content', parseInt(id, 10))
            .then(res => {
                getContent()
        })
        .catch(err => {
            console.log(err)
        })
    }

    const getContent = () => {
        Get(`content/${spaceID}`)
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
            <h1>Space {spaceID}</h1>

            <h2>Your Content</h2>
            <table>
            {
                content && content.map((d) => (
                    
                    <tr key={d.id} style={{lineHeight: '35px', minWidth: '450px'}} className="border-bottom">
                        <td>{d.name}</td>
                        <td>{d.value}</td>
                        <td><button value={d.id} 
                            style={{marginLeft: '15px', height: '24px', display: 'inline-block', marginTop: '8px'}} 
                            onClick={removeContent} type="button">Remove</button></td>
                    </tr>
                    
                ))
            }
            </table>
            <br />
            <h2>Add Content name and value</h2>
            <form onSubmit={newContent}>
                <div style={{minHeight: '80px'}}>
                    <div>
                        <label htmlFor="name">Content name</label>
                        <input id="name" type="text" name="name" />
                    </div>
                    <div>
                        <label htmlFor="name">Content value</label>
                        <textarea id="value" name="value" />
                    </div>
                </div>
                <button className="button" id="login-button" type="submit">Create</button>
                <button className="back" onClick={(event) => {
                        event.preventDefault()
                        navigate('/space')
                    }} type="button">Back</button>
            </form>
            
        </>
    )
}

export default Content