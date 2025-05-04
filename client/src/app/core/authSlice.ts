/* eslint-disable @typescript-eslint/no-explicit-any */
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";

// initialize the state
interface AuthState {
    isAuthenticated : boolean;
    user: null | {id: string; email:string};
}

const initialState: AuthState = {
    isAuthenticated : false,
    user: null
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers : {
        loginSuccess: (state, action: PayloadAction<{user: any}>) => {
            state.isAuthenticated = true;
            state.user = action.payload.user;
        },

        logoutSuccess: () => initialState
    }
});

export const checkAuthStatus = () => async(dispatch : any) => {
    try {

        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/user/verify`, {
            withCredentials: true
        })

        if(response.statusText == "OK"){
            const user = response.data
            dispatch(loginSuccess({user}))
        }

    } catch(error) {
        console.log("Error: ", error)
    }
}

export const {loginSuccess, logoutSuccess} = authSlice.actions;
export default authSlice.reducer;