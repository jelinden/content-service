import React from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom"
import Layout from './components/layout/Layout'
import Profile from './components/user/Profile'
import Space from './components/space/Space'
import Content from './components/space/Content'
import Login from './components/user/login/Login'
import Register from './components/user/register/Register'
import Home from './components/Home'
import { AppContextProvider } from './components/context/AppContext'
import './App.css'

function App() {

  return (
    <Router>
      <AppContextProvider>
        <Layout>
          <Routes>
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="/space" element={<Space />} />
            <Route path="/content/:spaceID" element={<Content />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/" element={<Home />} />
          </Routes>
        </Layout>
      </AppContextProvider>
    </Router>
  );
}

export default App;
