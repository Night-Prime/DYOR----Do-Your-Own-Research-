import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    triggered: false
}

export const alertSlice = createSlice({
    name: 'alert',
    initialState,
    reducers: {
        alertOn : (state) => {
            state.triggered = true;
        },

        alertOff : () => initialState
    }
});

export const {alertOn, alertOff} = alertSlice.actions;
export default alertSlice.reducer;