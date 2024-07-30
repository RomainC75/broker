import { useState, createContext, PropsWithChildren } from "react";
import { ISocketContext } from "../@types/socketContext.type";

const SocketContext = createContext<ISocketContext | null>(null);

const BROKER_HOST = process.env.BROKER_HOST || "";
const BROKER_PORT = process.env.BROKER_PORT || "";

function SocketProviderWrapper(props: PropsWithChildren) {
  const [myState, setMyState] = useState<number>(3);

  const socket = new WebSocket(`ws://${BROKER_HOST}:${BROKER_PORT}/reader`);

  socket.addEventListener("open", (event) => {
    socket.send("Connection established");
  });

  socket.addEventListener("message", (event) => {
    console.log("Message from server ", event.data);
  });

  return (
    <SocketContext.Provider value={{ myState }}>
      {props.children}
    </SocketContext.Provider>
  );
}

export { SocketContext, SocketProviderWrapper };
