import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { getUsers } from "@/services";

function RoomPage() {
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  const router = useRouter();
  const roomId = router.query.roomId;
  console.log("roomId", roomId);

  useEffect(() => {
    const fetchUsers = async () => {
      if (roomId) {
        const users = await getUsers(roomId.toString());
        console.log("users", users);
      }
    };
    fetchUsers();
  }, [roomId]);

  return <div>RoomPage</div>;
}

export default RoomPage;
