import { configureStore } from "@reduxjs/toolkit";
import authReducer from "./authSlice";
import alertReducer from "./alertSlice"

export const store = configureStore({
    reducer : {
        auth: authReducer,
        alert: alertReducer
    }
})


// infer the types
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;