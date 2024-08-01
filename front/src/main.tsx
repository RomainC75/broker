import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { SocketProviderWrapper } from './context/socket.context.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  // <React.StrictMode>
  <>
    <SocketProviderWrapper>
    <App />

    </SocketProviderWrapper>
  </>
  // </React.StrictMode>,
  
)
