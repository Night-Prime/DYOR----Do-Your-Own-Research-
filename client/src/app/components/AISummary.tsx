/* eslint-disable @typescript-eslint/no-unused-vars */
import { useState, useEffect } from 'react';
import { useAppSelector } from '../hooks/hook';
import { Asset } from '../data/models';
import { getAssetsInsights } from '../utils/service-calls';
import AIModal from "../shared/AIModal";


const AISummary = () => {
    const [currentAsset, setCurrentAsset] = useState<Asset | null>(null);
    const [lastShuffleTime, setLastShuffleTime] = useState(Date.now());
    const [summary, setSummary] = useState<string>("");
    const [showAIModal, setShowAIModal] = useState(false);

    const portfolio = useAppSelector((state) => state.portfolio);
    const assets = portfolio.assets ?? [];
    

    const getRandomAsset = () => {
        if (assets.length === 0) return null;

        const randomIndex = Math.floor(Math.random() * assets.length);
        return assets[randomIndex];
    };

    const shuffleAsset = async() => {
        const newAsset = getRandomAsset();
        setCurrentAsset(newAsset);
        setLastShuffleTime(Date.now());
        // const response = await getAssetsInsights(newAsset.symbol, "asset_insight");
        // if(response.success) setSummary(response.data.sentiment);
    };

    const processSummary = (summary: string): string => {
        const firstPeriodIndex = summary.indexOf('.');
        const firstSentence = firstPeriodIndex !== -1 
            ? summary.substring(0, firstPeriodIndex + 1) 
            : summary;
        
        return firstSentence;
    }

    const insight = processSummary(summary);

    const toggleModal = () => {
        setShowAIModal(!showAIModal);
    }

    useEffect(() => {
        if (assets.length > 0 && !currentAsset) {
            shuffleAsset();
        }
    }, [assets.length]);

    // Set up hourly interval
    useEffect(() => {
        const interval = setInterval(() => {
            shuffleAsset();
        }, 10 * 60 * 1000); // 1 hour in milliseconds

        return () => clearInterval(interval);
    }, [assets]);


    return (
        <>
            <div className="w-full h-full bg-gray-50 rounded-3xl shadow-sm">
            <div className="w-full h-full flex flex-col items-center justify-center gap-4 py-4 ">
                <div>
                    <h1 className='text-6xl font-bold'>
                        {currentAsset ? `$${currentAsset.symbol}` : ""}
                    </h1>
                </div>
                <div>
                    <h3>
                        {currentAsset ? `${currentAsset.name} (${currentAsset.type})` : "" }
                    </h3>
                </div>
                <div className='w-full'>
                    <p className='text-xs font-semibold m-auto w-[80%] text-center'>
                        {insight}
                    </p>
                </div>
                    {summary && (<div>
                        <button className='bg-lime-900 text-white font-medium text-sm rounded-3xl shadow-sm hover:bg-lime-950 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300 px-4 py-2 cursor-pointer' onClick={toggleModal}>
                            {/* {assets.length > 0 ? "See more" : "Add Assets"} */}
                            See more
                        </button>
                    </div>
                )}
            </div>
        </div>
        {showAIModal && <AIModal close={toggleModal} summary={summary} asset={currentAsset} />}
        </>
    )
}

export default AISummary
