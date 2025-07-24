import React, { useEffect, useState } from 'react'
import { ClearRounded } from '@mui/icons-material';
import { Asset } from '../data/models';
import { getAssetsInsights } from '../utils/service-calls';

interface AssetAIProps {
    close: () => void;
    asset: Asset
}

const AssetAISummary: React.FC<AssetAIProps> = ({ close, asset }) => {
    const [summary, setSummary] = useState<string>('');
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const getAssetInsight = async () => {
        try {
            setIsLoading(true);
            const response = await getAssetsInsights(asset.symbol, "asset_insight");
            if (response.success) {
                setSummary(response.data.sentiment);
            } else {
                setError('Error Occurred!, Try again');
            }
        } catch (err) {
            setError('An error occurred while fetching data');
            console.error("Error fetching insights:", err);
        } finally {
            setIsLoading(false);
        }
    };

    useEffect(() => {
        getAssetInsight();
    }, [asset.symbol]);

    const handleClose = () => {
        setSummary('');
        setError(null);
        setIsLoading(false);
        close?.();
    };

    return (
        <div className="fixed inset-0 z-[9999] bg-black/50 backdrop-blur-xs flex items-center justify-center p-4">
            <div className="w-full max-w-2xl bg-white rounded-lg shadow-xl flex flex-col max-h-[90vh]">
                <div className="flex justify-end p-4">
                    <button
                        className="text-gray-500 hover:text-gray-700 focus:outline-none"
                        onClick={handleClose}
                    >
                        <ClearRounded />
                    </button>
                </div>

                <div className="px-6 pb-6 flex-1 overflow-y-auto">
                    <div className="mb-4">
                        <h1 className="text-6xl font-bold">
                            ${asset.symbol}
                        </h1>
                    </div>

                    <div className="mb-6">
                        <h3 className="text-xl font-semibold text-gray-600">
                            {asset.name}
                        </h3>
                    </div>

                    <div className="bg-gray-50 rounded-lg p-4 min-h-[100px]">
                        {isLoading ? (
                            <div className="space-y-3">
                            <div className="flex space-x-2">
                                <div className="h-3 w-3 bg-lime-800 rounded-full animate-bounce"></div>
                                <div className="h-3 w-3 bg-lime-800 rounded-full animate-bounce" style={{ animationDelay: '0.2s' }}></div>
                                <div className="h-3 w-3 bg-lime-800 rounded-full animate-bounce" style={{ animationDelay: '0.4s' }}></div>
                                <div className="h-3 w-3 bg-lime-800 rounded-full animate-bounce" style={{ animationDelay: '0.6s' }}></div>
                            </div>
                            <p className="text-sm text-gray-400 italic">Gathering Information & Insights just for you...</p>
                        </div>
                        ) : error ? (
                            <p className="text-red-500">{error}</p>
                        ) : (
                            <p className="text-sm leading-relaxed overflow-y-auto max-h-[50vh]">
                                {summary}
                            </p>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AssetAISummary
