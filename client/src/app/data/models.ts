export interface Portfolio {
    id:string;
    user_id: string;
    name: string;
    total_value: number;
    created_at: string;
    updated_at: string;
    deleted_at: string | null;
}

export interface User {
    id: string;
    first_name: string;
    last_name: string;
    avatar: string;
    email: string;
    role: string;
    password:string;
    created_at:string;
    updated_at: string;
    deleted_at: null;
    portfolios: Portfolio[]
}

export interface ApiResponse<T> {
    success: boolean;
    data?: T;
    errors?: Record<string, string[]>;
}

export interface Asset {
    name: string;
    symbol: string
}

export interface AssetPayload {
    type : string;
    symbols: Asset[];
    portfolioID: string;
}
