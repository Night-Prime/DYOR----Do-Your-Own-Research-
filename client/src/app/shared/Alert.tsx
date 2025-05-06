"use client"

import React from "react";
import { Alert } from "@material-tailwind/react";

// A custom reusable alert component (need to make this work with redux)
type AlertType = "success" | "error" | "warning" | "info";

interface DyorAlertProps {
  title?: string;
  message: string;
  type: AlertType;
  open: boolean;
  onClose?: () => void;
  autoClose?: boolean;
  autoCloseDuration?: number;
}

const iconMap = {
  success: (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="h-6 w-6">
      <path fillRule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zm13.36-1.814a.75.75 0 10-1.22-.872l-3.236 4.53L9.53 12.22a.75.75 0 00-1.06 1.06l2.25 2.25a.75.75 0 001.14-.094l3.75-5.25z" clipRule="evenodd" />
    </svg>
  ),
  error: (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="h-6 w-6">
      <path fillRule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25zm-1.72 6.97a.75.75 0 10-1.06 1.06L10.94 12l-1.72 1.72a.75.75 0 101.06 1.06L12 13.06l1.72 1.72a.75.75 0 101.06-1.06L13.06 12l1.72-1.72a.75.75 0 10-1.06-1.06L12 10.94l-1.72-1.72z" clipRule="evenodd" />
    </svg>
  ),
  warning: (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="h-6 w-6">
      <path fillRule="evenodd" d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z" clipRule="evenodd" />
    </svg>
  ),
  info: (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="h-6 w-6">
      <path fillRule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zm8.706-1.442c1.146-.573 2.437.463 2.126 1.706l-.709 2.836.042-.02a.75.75 0 01.67 1.34l-.04.022c-1.147.573-2.438-.463-2.127-1.706l.71-2.836-.042.02a.75.75 0 11-.671-1.34l.041-.022zM12 9a.75.75 0 100-1.5.75.75 0 000 1.5z" clipRule="evenodd" />
    </svg>
  )
}


export const DyorAlert: React.FC<DyorAlertProps> = ({
  message,
  type,
  open,
  onClose,
  autoClose = true,
  autoCloseDuration = 5000,
}) => {
  React.useEffect(() => {
    if (autoClose && open) {
      const timer = setTimeout(() => {
        onClose?.();
      }, autoCloseDuration);
      return () => clearTimeout(timer);
    }
  }, [autoClose, autoCloseDuration, onClose, open]);

  const alertClasses = {
    success: "bg-lime-500",
    error: "bg-red-600", 
    warning: "bg-amber-600",
    info: "bg-blue-600",
  };

  if (!open) return null;

  return (
    <div className="fixed inset-0 pointer-events-none z-50">
      <Alert
        className={`absolute top-4 right-4 w-1/4 rounded-3xl shadow-xl ${alertClasses[type]} p-3`}
        icon={iconMap[type]}
        variant="filled"
        animate={{
          mount: { y: 0, opacity: 1 },
          unmount: { y: -100, opacity: 0 },
        }}
      >
          <span className="mx-2"><b>{message}</b></span>
      </Alert>
    </div>
  );
};
