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

export const getAssetData = async (type: string, symbol: string) => {
    if(type === 'crypto') {
        const response = await getCryptoData(symbol);
        return response;
    }

    if(type === 'stock') {
        const response = await getStockData(symbol);
        return response;
    }
    
}

const getCryptoData = async (symbol: string): Promise<T> => {
    try {
        const type = 'crypto';
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/asset/get-live-update?type=${type}&symbols=${symbol}`, {
            withCredentials: true,
        });

        return {
            success: true,
            data: response.data.crypto_data.data
        }
    } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
            return {
                errors: error.response.data.errors || { general: ['An error occurred during signup.'] },
            };
        }
        return {
            errors: { general: ['Failed to connect to the signup service.'] },
        };
    }
}

const getStockData = async (symbol: string): Promise<T> => {
    try {
        const type = 'stock';
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/asset/get-live-update?type=${type}&symbols=${symbol}`, {
            withCredentials: true,
        });

        return {
            success: true,
            data: response.data?.stock_data.data.quoteResponse.result[0]
        }
    } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
            return {
                errors: error.response.data.errors || { general: ['An error occurred during signup.'] },
            };
        }
        return {
            errors: { general: ['Failed to connect to the signup service.'] },
        };
    }
}