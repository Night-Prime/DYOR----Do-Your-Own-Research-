import { createSlice, PayloadAction } from "@reduxjs/toolkit";
interface AlertState {
    open: boolean;
    type: "success" | "error" | "warning" | "info";
    message: string;
    autoClose: boolean;
    autoCloseDuration: number;
}

const initialState: AlertState = {
    open: false,
    type: "info",
    message: "",
    autoClose: true,
    autoCloseDuration: 5000,
};

export const alertSlice = createSlice({
  name: "alert",
  initialState,
  reducers: {
    showAlert: (
      state,
      action: PayloadAction<{
        type: AlertState["type"];
        message: string;
        autoClose?: boolean;
        duration?: number;
      }>
    ) => {
      state.open = true;
      state.type = action.payload.type;
      state.message = action.payload.message;
      state.autoClose = action.payload.autoClose ?? true;
      state.autoCloseDuration = action.payload.duration ?? 5000;
    },
    hideAlert: (state) => {
      state.open = false;
    },
    // Optional: Add a reset action
    resetAlert: () => initialState,
  },
});

export const { showAlert, hideAlert, resetAlert } = alertSlice.actions;
export default alertSlice.reducer;
