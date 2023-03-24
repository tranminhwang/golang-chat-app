import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { AuthContextProvider } from "@/contexts/AuthContextProvider";
import WebSocketContentProvider from "@/contexts/WebSocketContextProvider";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AuthContextProvider>
      <WebSocketContentProvider>
        <Component {...pageProps} />
      </WebSocketContentProvider>
    </AuthContextProvider>
  );
}
