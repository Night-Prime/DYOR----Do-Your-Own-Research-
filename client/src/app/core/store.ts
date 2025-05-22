import { configureStore } from "@reduxjs/toolkit";
import { persistStore, persistReducer } from 'redux-persist';
import createWebStorage from "redux-persist/lib/storage/createWebStorage";
import authReducer from "./authSlice";
import alertReducer from "./alertSlice";

// configuration for Persists:
const persistConfig = {
    key: 'auth',
    storage: createWebStorage('session'),
    whitelist : ['auth']
}

const persistAuthReducer = persistReducer(persistConfig, authReducer)

export const store = configureStore({
    reducer : {
        auth: persistAuthReducer,
        alert: alertReducer
    },
    middleware: (getDefaultMiddleware) => 
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: ['persist/PERSIST']
            }
        })
})

export const persistor = persistStore(store);

// infer the types
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;