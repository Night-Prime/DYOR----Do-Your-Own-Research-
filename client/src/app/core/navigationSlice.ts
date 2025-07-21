import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Header } from "../data/models";
import { headerCopy } from "../data/header";

interface NavigationState {
    isNavigated: boolean;
    header: Header | null;
}

const initialState: NavigationState = {
    isNavigated: false, // Changed from true to false initially
    header: null,
}

export const navigationSlice = createSlice({
    name: 'navigation',
    initialState,
    reducers: {
        setCurrentPage: (state, action: PayloadAction<string>) => {
            let pagePath = action.payload.trim();
            console.log(pagePath);
            if (!pagePath || pagePath === '/') pagePath = 'dashboard';
            
            const basePath = pagePath.split('/')[0];
            const headerData = headerCopy.find(item => item.page === basePath);

            if (headerData) {
                state.header = headerData;
                state.isNavigated = true;
            } else {
                // Fallback to dashboard if route not found
                const fallback = headerCopy.find(item => item.page === 'dashboard');
                state.header = fallback || null;
                state.isNavigated = !!fallback;
            }
        },
        resetNavigation: (state) => {
            state.isNavigated = false;
            state.header = null;
        }
    }
});

export const { setCurrentPage, resetNavigation } = navigationSlice.actions;
export default navigationSlice.reducer;