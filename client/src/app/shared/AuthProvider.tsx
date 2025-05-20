"use client"

import { useAppDispatch, useAppSelector } from "../hooks/hook";
import { useEffect } from "react";
import { checkAuthStatus } from "../core/authSlice";
import { useRouter } from "next/navigation";

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const router = useRouter();
    const dispatch = useAppDispatch();
    const { isAuthenticated } = useAppSelector(state => state.auth);

    useEffect(() => {
        const verifyAuth = async () => {
            const isAuthenticated = await dispatch(checkAuthStatus());
            if (!isAuthenticated) {
                router.replace('/');
            }
        };
        
        verifyAuth();
    }, [router, dispatch]);

    return isAuthenticated ? children : null;
};