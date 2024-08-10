import { useContext, useEffect } from 'react'
import './App.css'
import { ISocketContext } from './@types/socketContext.type'
import { SocketContext } from './context/socket.context'
import HomePage from './pages/HomePage'
import { useMsal, useMsalAuthentication } from '@azure/msal-react'
import { InteractionType } from '@azure/msal-browser'

function App() {
  console.log("-> ", process.env.BROKER_HOST)
  // const { user, handleUser } = useContext(AuthContext) as IAuthContext;
  useMsalAuthentication(InteractionType.Redirect, {
    scopes:["User.Read"]
  })
  const {instance} = useMsal()
  useEffect(()=>{
    console.log("get active : ", instance.getActiveAccount())
  }, [instance])

  const { myState } = useContext(SocketContext) as ISocketContext;
  
  return (
    <>
      <HomePage/>
    </>
  )
}

export default App
