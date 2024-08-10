import { useState, createContext, PropsWithChildren, useEffect } from "react";
import { ISocketContext } from "../@types/socketContext.type";
import useWebSocket from 'react-use-websocket';
import { dataJson } from "./data";
import { TopicMapDto } from "../@types/back_types";

const SocketContext = createContext<ISocketContext | null>(null);

const BROKER_HOST = process.env.BROKER_HOST_FROM_FRONT || "";
// const BROKER_HOST = "host.docker.internal"
const BROKER_PORT = process.env.BROKER_PORT || "";

const WS_URL = `ws://${BROKER_HOST}:${BROKER_PORT}/reader`;
const socket = new WebSocket(WS_URL)

function SocketProviderWrapper(props: PropsWithChildren) {
  const [myState, setMyState] = useState<number>(3);
  // const [data, setData] = useState<TopicMapDto | null>()
  const [data, setData] = useState<TopicMapDto | null>(JSON.parse(dataJson))
  
  

  useEffect(() => {
    socket.onopen = () => {
      // setMessage('Connected')
    };

    socket.onmessage = (e) => {
      // setMessage("Get message from server: " + e.data)
      console.log("->", e.data)
      const content = JSON.parse(e.data)
      console.log("==>", content)
      setData(content)

    };

    socket.onclose = () => {
      console.log("closed ")
    };

    console.log("- ", data)


    return () => {
      socket.close()
    }
  }, [])

  return (
    <SocketContext.Provider value={{ myState, data }}>
      {props.children}
    </SocketContext.Provider>
  );
}

export { SocketContext, SocketProviderWrapper };
