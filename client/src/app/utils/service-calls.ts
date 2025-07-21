/* eslint-disable @typescript-eslint/no-explicit-any */

// specifically made this to fetch data coming from external service through my
// server acting as a proxy server, so that I can avoid CORS issues.

// this would move from using Http calls to Websockets in the near future
// to get live updates of the assets (just gotta pay for premium first)

import axios from "axios";

export interface T {
    success?: boolean;
    data?: any;
    errors?: Record<string, string[]>;
}


export const getAssetData = async (stock_symbol?: string[], crypto_symbol?: string[]): Promise<any> => {
    try {
        console.log("Available: ", stock_symbol, crypto_symbol);
        const stockParams = stock_symbol ? `stock_symbols=${stock_symbol.join(",")}` : "";
        const cryptoParams = crypto_symbol ? `crypto_symbols=${crypto_symbol.join(",")}` : "";
        const queryParams = [stockParams, cryptoParams].filter(Boolean).join("&");
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/asset/get-live-update?${queryParams}`, {
            withCredentials: true,
        });
        //TODO we set an env that triggers storing data to the localStorage 
        localStorage.setItem("recommended", JSON.stringify(response))
        return {
            success: true,
            data: response,
        };
    } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
            return {
                errors: error.response.data.errors || { general: ["An error occurred during signup."] },
            };
        }
        return {
            errors: { general: ["Failed to connect to the signup service."] },
        };
    }
};

export const getAssetsInsights = async(asset: string, promptType : string):Promise<any> => {
    try {
        const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/asset/get-ai-summary`, {
            asset_info : asset,
            prompt_type: promptType
        });

        return {
            success: true,
            data: response.data,
        };

    } catch(error) {
        if (axios.isAxiosError(error) && error.response) {
            return {
                errors: error.response.data.errors || { general: ["An error occurred during signup."] },
            };
        }
        return {
            errors: { general: ["Failed to connect to the signup service."] },
        };
    }
}