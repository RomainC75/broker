import { useState, createContext, PropsWithChildren, useEffect } from "react";
import { ISocketContext } from "../@types/socketContext.type";
import useWebSocket from 'react-use-websocket';

const SocketContext = createContext<ISocketContext | null>(null);

const BROKER_HOST = process.env.BROKER_HOST || "";
const BROKER_PORT = process.env.BROKER_PORT || "";

const WS_URL = `ws://${BROKER_HOST}:${BROKER_PORT}/reader`;
const socket = new WebSocket(WS_URL)

function SocketProviderWrapper(props: PropsWithChildren) {
  const [myState, setMyState] = useState<number>(3);

  

  useEffect(() => {
    socket.onopen = () => {
      // setMessage('Connected')
      console.log("- CONNECTED")
    };

    socket.onmessage = (e) => {
      // setMessage("Get message from server: " + e.data)
      console.log("==>", JSON.parse(e.data))

    };

    socket.onclose = () => {
      console.log("closed ")
    };

    return () => {
      socket.close()
    }
  }, [])

  return (
    <SocketContext.Provider value={{ myState }}>
      {props.children}
    </SocketContext.Provider>
  );
}

export { SocketContext, SocketProviderWrapper };
