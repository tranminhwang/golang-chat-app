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
