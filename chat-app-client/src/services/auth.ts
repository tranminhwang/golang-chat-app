import { API_URL } from "../constants";

export interface ILoginPayload {
  email: string;
  password: string;
}

export interface ILoginResponse {
  accessToken: string;
  id: string;
  email: string;
  username: string;
}

export const serviceLogin = async (
  payload: ILoginPayload
): Promise<ILoginResponse> => {
  const response = await fetch(`${API_URL}/api/v1/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });
  if (response.ok) {
    const data = await response.json();
    return data;
  } else {
    throw new Error("failed to login");
  }
};
