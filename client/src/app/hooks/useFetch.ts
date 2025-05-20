import axios, { AxiosError, AxiosRequestConfig } from 'axios';
import { useCallback, useEffect, useRef, useState } from 'react';

const CACHE_TTL = 1000 * 60 * 5;

export type DyorResponse<T = unknown> = {
    status: number | null;
    statusText: string;
    data: T | null;
    error: Error | null;
    loading: boolean;
    refresh: () => void;
}

export const useFetch = <T = unknown>(
    url: string, 
    params: Record<string, unknown> = {}, 
    retries: number = 2,
    config?: AxiosRequestConfig
): DyorResponse<T> => {
    const [status, setStatus] = useState<number | null>(null);
    const [statusText, setStatusText] = useState<string>('');
    const [data, setData] = useState<T | null>(null);
    const [error, setError] = useState<Error | null>(null);
    const [loading, setLoading] = useState<boolean>(false);

    const cache = useRef(new Map<string, { data: T; timestamp: number }>());
    const serviceURL = `${process.env.NEXT_PUBLIC_API_URL}/${url}`

    const getData = useCallback(async (currentRetry = 0): Promise<void> => {
        const abortController = new AbortController();
        const cacheKey = `${serviceURL}-${JSON.stringify(params)}`;
        setLoading(true);
        
        try {
            if (cache.current.has(cacheKey)) {
                const cached = cache.current.get(cacheKey)!;
                if (Date.now() - cached.timestamp < CACHE_TTL) {
                    setData(cached.data);
                    setLoading(false);
                    return;
                }
            }

            const response = await axios.get<T>(serviceURL, {
                params,
                withCredentials: true,
                signal: abortController.signal,
                ...config
            });

            setStatus(response.status);
            setStatusText(response.statusText);
            setData(response.data);
            setError(null);
            
            cache.current.set(cacheKey, { 
                data: response.data, 
                timestamp: Date.now() 
            });
        } catch (error) {
            if (abortController.signal.aborted) return;

            if (currentRetry < retries && shouldRetry(error)) {
                const delay = Math.min(1000 * 2 ** currentRetry, 30000);
                await new Promise(res => setTimeout(res, delay));
                return getData(currentRetry + 1);
            }

            const err = error as AxiosError | Error;
            setError(err);
            setStatus(axios.isAxiosError(err) ? err.response?.status || null : null);
        } finally {
            if (!abortController.signal.aborted) {
                setLoading(false);
            }
        }
    }, []);

    const refresh = () => {
        const cacheKey = `${serviceURL}-${JSON.stringify(params)}`;
        cache.current.delete(cacheKey);
        getData();
    };

    useEffect(() => {
        const handler = setTimeout(() => getData(), 300);
        return () => clearTimeout(handler);
    },[getData]);

    return { status, statusText, data, error, loading, refresh };
};

function shouldRetry(error: unknown): boolean {
    if (!axios.isAxiosError(error)) return false;
    
    // Retry on network errors or server errors
    return !error.response || error.response.status >= 500;
}