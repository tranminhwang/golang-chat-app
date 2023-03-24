import { API_URL } from "../constants";

export interface ICreateRoomPayload {
  id: string;
  name: string;
}
export const createRoom = async (payload: ICreateRoomPayload): Promise<any> => {
  const response = await fetch(`${API_URL}/ws/create_room`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("access_token")}`,
    },
    credentials: "include",
    body: JSON.stringify(payload),
  });
  if (response.ok) {
    const data = await response.json();
    return data;
  } else {
    throw new Error("failed to login");
  }
};

export const getRooms = async (): Promise<any> => {
  const res = await fetch(`${API_URL}/ws/get_rooms`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("access_token")}`,
    },
    credentials: "include",
  });
  if (res.ok) {
    return await res.json();
  }
  throw new Error("failed to get rooms");
};

export const getUsers = async (roomId: string): Promise<any> => {
  const res = await fetch(`${API_URL}/ws/get_clients/${roomId}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("access_token")}`,
    },
    credentials: "include",
  });
  if (res.ok) {
    return await res.json();
  }
  throw new Error("failed to get users");
};
