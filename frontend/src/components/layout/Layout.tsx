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
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/register">Register</Link></li>
                    <li><Link to="/login">Login</Link></li>
                </ul>
            </div>
            <main>{children}</main>
        </>
    )
}

export default Layout;