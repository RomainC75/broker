import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { SocketProviderWrapper } from "./context/socket.context.tsx";
import AuthProvider from "./auth/Auth.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  // <React.StrictMode>
    <AuthProvider>
      <SocketProviderWrapper>
        <App />
        {/* <div>HHHHHHEEEE</div> */}
      </SocketProviderWrapper>
    </AuthProvider>
  // </React.StrictMode>,
);
