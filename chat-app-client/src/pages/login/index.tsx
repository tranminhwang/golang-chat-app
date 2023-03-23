import React, { useState, useContext, useEffect } from "react";
import { useRouter } from "next/router";
import { PrimaryButton } from "@/components/shared/Button";
import { AuthContext, IUserInfo } from "../../contexts/AuthContextProvider";
import { serviceLogin } from "../../services/auth";

function Login() {
  const [userInfo, setUserInfo] = useState<IUserInfo>();
  const [loading, setLoading] = useState(false);
  const route = useRouter();
  const { authenticated } = useContext(AuthContext);

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const emailInput = e.currentTarget.email.value;
    const passwordInput = e.currentTarget.password.value;

    try {
      // await new Promise((resolve) => setTimeout(resolve, 3000));
      setLoading(true);
      const { accessToken, ...userInfo } = await serviceLogin({
        email: emailInput,
        password: passwordInput,
      });
      localStorage.setItem("user_info", JSON.stringify(userInfo));
      return route.push("/");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (authenticated) {
      route.push("/");
      return;
    }
  }, [authenticated, route]);

  return (
    <div className="w-full h-screen flex justify-center items-center bg-gray-100 dark:bg-gray-900">
      <form
        className="w-1/3  px-8 py-10 rounded-lg shadow-lg bg-white dark:bg-gray-800"
        onSubmit={onSubmit}
      >
        <div className="mb-6">
          <label
            htmlFor="email"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Your email
          </label>
          <input
            name="email"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="name@flowbite.com"
            required
          />
        </div>
        <div className="mb-6">
          <label
            htmlFor="password"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Your password
          </label>
          <input
            type="password"
            name="password"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            required
          />
        </div>
        <PrimaryButton disabled={loading}>Login</PrimaryButton>
      </form>
    </div>
  );
}

export default Login;
