import React, { useContext } from "react";
import { useRouter } from "next/router";
import { AuthContext } from "@/contexts/AuthContextProvider";
import { WebSocketContext } from "@/contexts/WebSocketContextProvider";
import { WS_URL } from "@/constants";

export interface IRoomItem {
  id: string;
  name: string;
}

const RoomItem: React.FC<IRoomItem> = ({ id, name }) => {
  const route = useRouter();
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebSocketContext);

  const handleJoinRoom = () => {
    const ws = new WebSocket(
      `${WS_URL}/ws/join_room/${id}?userId=${user.id}&username=${user.username}`
    );
    if (ws.OPEN) {
      setConn(ws);
      route.push(`/room/${id}`);
      return;
    }
  };

  return (
    <div className="flex justify-between text-[#0063ec] bg-white rounded-lg px-8 py-6">
      <h3 className="text-2xl font-medium">{name}</h3>
      <button
        onClick={handleJoinRoom}
        type="button"
        className="px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-full hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        Join
      </button>
    </div>
  );
};

export default RoomItem;
