/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";
import { AppDispatch } from "./store";
import { User } from "../data/models";

// initialize the state
interface AuthState {
    isAuthenticated : boolean;
    user: User | null;
}

const initialState: AuthState = {
    isAuthenticated : false,
    user: null,
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers : {
        loginSuccess: (state, action: PayloadAction<User>) => {
            state.isAuthenticated = true;
            state.user = action.payload;
            console.log("State: ", action.payload);
        },

        logoutSuccess: () => initialState
    }
});

export const checkAuthStatus = () => async (dispatch: AppDispatch) => {
    try {
        const { data } = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/user/verify`, { 
            withCredentials: true 
        });
        dispatch(loginSuccess(data));
        return true;
    } catch (error) {
        if (axios.isAxiosError(error) && error.response?.status === 401) {
            dispatch(logoutSuccess());
        }
        return false;
    }
};
export const {loginSuccess, logoutSuccess} = authSlice.actions;
export default authSlice.reducer;