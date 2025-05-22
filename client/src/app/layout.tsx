"use client";

import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { Provider } from "react-redux";
import { ThemeProvider } from "@material-tailwind/react";
import { persistor, store } from "./core/store";
import { PersistGate } from "redux-persist/integration/react";
import { GlobalAlert } from "./shared/GlobalAlert";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// export const metadata: Metadata = {
//   title: "Do Your Own Research",
//   description: "D.Y.O.R platform is used to track the markets and provide real-time insights and information into the current state of stocks, bonds and crypto investments made by the users.",
// };

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <ThemeProvider>
          <Provider store={store}>
          <PersistGate loading={null} persistor={persistor}>
              {children}
              <GlobalAlert />
              </PersistGate>
            </Provider>
        </ThemeProvider>
      </body>
    </html>
  );
}
