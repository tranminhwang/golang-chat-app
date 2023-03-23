import { useRouter } from "next/router";
import { useState, createContext, useEffect } from "react";

export interface IUserInfo {
  id: string;
  username: string;
  email: string;
}

export const initialUser: IUserInfo = {
  id: "",
  username: "",
  email: "",
};

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: (authenticated: boolean) => void;
  user: IUserInfo;
  setUser: (user: IUserInfo) => void;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: initialUser,
  setUser: () => {},
});

export const AuthContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUser] = useState<IUserInfo>(initialUser);
  const route = useRouter();

  useEffect(() => {
    const userInfo = localStorage.getItem("user_info");
    if (!userInfo) {
      if (window.location.pathname !== "/register") {
        route.push("/login");
        return;
      }
    } else {
      console.log("userInfo", userInfo);
      const user: IUserInfo = JSON.parse(userInfo);
      setUser(user);
      setAuthenticated(true);
    }
  }, [authenticated]);

  return (
    <AuthContext.Provider
      value={{ authenticated, setAuthenticated, user, setUser }}
    >
      {children}
    </AuthContext.Provider>
  );
};
