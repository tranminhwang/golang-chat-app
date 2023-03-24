import React, { createContext, ReactNode, useState } from "react";

type Conn = WebSocket | null;
interface IWebSocketContext {
  conn: Conn;
  setConn: React.Dispatch<React.SetStateAction<Conn>>;
}

export const WebSocketContext = createContext<IWebSocketContext>({
  conn: null,
  setConn: () => {},
});

const WebSocketContentProvider = ({ children }: { children: ReactNode }) => {
  const [conn, setConn] = useState<Conn>(null);

  return (
    <WebSocketContext.Provider
      value={{
        conn,
        setConn,
      }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};

export default WebSocketContentProvider;
