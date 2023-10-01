import React from 'react';
import {
    Link
  } from "react-router-dom";
import { Props } from './LayoutTypeProps';

const Layout =({children} : Props) =>{
    return(
        <>
            <div>
                <ul>
                    <li className='links'><Link to="/">Home</Link></li>
                    <li className='links'><Link to="/register">Register</Link></li>
                    <li className='links'><Link to="/login">Login</Link></li>
                </ul>
            </div>
            <main>{children}</main>
        </>
    )
}

export default Layout;