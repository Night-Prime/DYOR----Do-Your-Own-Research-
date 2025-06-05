// this would hold most portfolio Data

import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Portfolio } from "../data/models";

const initialState: Portfolio = {
    id: "",
    user_id: "",
    name: "",
    total_value: 0,
    investment_goals: [],
    asset_preference: [],
    created_at: "",
    updated_at: "",
    deleted_at: null,
    assets: [],
}

export const portfolioSlice = createSlice({
    name: 'portfolio',
    initialState,
    reducers: {
        setPortfolio: (state, action: PayloadAction<Portfolio | undefined>) => {
            return { ...state, ...action.payload };
        },
        clearPortfolio: () => initialState,
    }
});

export const { setPortfolio, clearPortfolio } = portfolioSlice.actions;
export default portfolioSlice.reducer;