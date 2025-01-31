import { useContext } from 'react'
import './App.css'
import { ISocketContext } from './@types/socketContext.type'
import { SocketContext } from './context/socket.context'
import HomePage from './pages/HomePage'

function App() {
  console.log("-> ", process.env.BROKER_HOST)
  const { myState } = useContext(SocketContext) as ISocketContext;
  return (
    <>
      <HomePage/>
    </>
  )
}

export default App
