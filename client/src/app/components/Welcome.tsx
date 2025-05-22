/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useState } from 'react'
import { Asset, User } from '../data/models';
import AssetBtn from '../shared/AssetBtn';
import { stocks } from '../data/asset';
import { saveAssets } from '../utils/api';
import { DyorAlert } from '../shared/Alert';

interface WelcomeProps {
    user: User | null,
    refresh: () => void
}


const Welcome: React.FC<WelcomeProps> = ({ user, refresh }) => {
    const [assets, setAssets] = useState<Asset[]>([]);
    const [selectedAssets, setSelectedAssets] = useState<{ [key: string]: boolean }>({});
    const [alert, setAlert] = useState<{
        open: boolean;
        type: 'success' | 'error';
        message: string;
    }>({
        open: false,
        type: 'success',
        message: ''
    });

    const addAsset = (stock: { name: string, symbol: string }) => {
        setAssets(prevAssets => {
            const exists = prevAssets.some(asset => asset.symbol === stock.symbol);
            return exists
                ? prevAssets.filter(asset => asset.symbol !== stock.symbol)
                : [...prevAssets, stock];
        });

        setSelectedAssets(prev => ({
            ...prev,
            [stock.symbol]: !prev[stock.symbol]
        }));
    };

    const saveAsset = async () => {
        const newPayload : any = {
            type: 'stock',
            symbols: assets,
            portfolioID: user?.portfolios[0]?.id
        };
        const response = await saveAssets(newPayload);

        setAlert({
            open: true,
            type: response.success ? 'success' : 'error',
            message: response.success
                ? 'Stocks Added Successfully'
                : 'Stocks couldn\'t be added'
        });

        if (response.success) {
            setAssets([]);
        }

        refresh();
    }

    return (
        <div className="w-full h-full fixed inset-0 flex backdrop-blur-sm items-center justify-center overflow-hidden z-40">
            <DyorAlert
                type={alert.type}
                message={alert.message}
                open={alert.open}
                autoClose={true}
                onClose={() => setAlert(prev => ({ ...prev, open: false }))}
            />
            <div className="absolute inset-0 flex items-start justify-center py-20">
                <div
                    className="relative bg-white w-full max-w-screen-lg max-h-[80vh] rounded-3xl p-6 shadow-lg z-10 overflow-y-auto"
                    style={{ scrollbarWidth: "none", msOverflowStyle: "none" }}
                >
                    <div className='w-full h-full flex flex-col gap-4'>
                        <div className='w-full border-b-[0.5px] border-gray'>
                            <h3 className='text-4xl font-extrabold'>Hello {user ? user?.first_name : 'User'}, Welcome to DYOR </h3>
                        </div>
                        <div className='w-full'>
                            <h3 className='text-lg'>Select the financial investments you wish to add to your portfolio below.</h3>
                        </div>
                        <div className="w-full mt-4 flex flex-col gap-4">
                            <div className='w-full h-full px-4'>
                                <h5 className='text-xl font-semibold border-b-[.5px]'>Stocks</h5>
                                <div className='w-full h-full py-4'>
                                    <div className='flex flex-wrap gap-3'>
                                        {stocks && stocks.map((stock, i) => (
                                            <span key={i} onClick={() => addAsset(stock)}>
                                                <AssetBtn name={stock.name} select={!!selectedAssets[stock.symbol]} />
                                            </span>
                                        ))}
                                    </div>
                                </div>
                            </div>

                            {/* <div className='w-full h-full px-4'>
                                <h5 className='text-xl font-semibold border-b-[.5px]'>Crypto</h5>
                                <div className='w-full h-full py-4'>
                                    <div className='flex flex-wrap gap-3'>
                                        {crypto && crypto.map((stock, i) => (
                                            <span key={i} onClick={() => addAsset(stock)}>
                                                <AssetBtn name={stock.name} select={!!selectedAssets[stock.symbol]} />
                                            </span>
                                        ))}
                                    </div>
                                </div>
                            </div> */}

                            <div className='mt-4 w-full flex flex-row justify-end items-end'>
                                <button
                                    onClick={saveAsset}
                                    className="w-1/6 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-3xl shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
                                >
                                    Save Portfolio
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    )
}

export default Welcome
