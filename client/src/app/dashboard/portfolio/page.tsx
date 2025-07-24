"use client"

import AssetAISummary from '@/app/components/AssetAISummary';
import AssetContainers from '@/app/components/AssetContainers';
import { Asset, User } from '@/app/data/models';
import { useAppSelector } from '@/app/hooks/hook'
import { useFetch } from '@/app/hooks/useFetch';
import { DyorAlert } from '@/app/shared/Alert';
import PortfolioHealthGauge from '@/app/shared/charts/PortfolioHealthGauge';
import RiskRadarChart from '@/app/shared/charts/RadarChart';
import Preloader from '@/app/shared/Preloader';
import React, { useState } from 'react'

const Portfolio = () => {
  const [showAIModal, setShowAIModal] = useState<boolean>(false);
  const [selectedAsset, setSelectedAsset] = useState<Asset | null>(null);

  const toggleModal = () => {
    setShowAIModal(!showAIModal);
  };

  const handleAssetClick = (asset: Asset) => {
    setSelectedAsset(asset);
    setShowAIModal(true);
  };

  // TODO need to abstract the entire api call to a store
  const userId = useAppSelector((state) => state.auth.user?.id);
  const { data, loading, error } = useFetch<User>("user/portfolio", {
    id: userId
  });
  const assets = data?.portfolios[0]?.assets;

  if (loading) return <Preloader />;
  if (error) {
    return <DyorAlert type="error" message={`${error}`} open autoClose />;
  }

  const { sentiment_score } = JSON.parse(localStorage.getItem('user/sentiment-score') || '{}');
  const market_score = sentiment_score * 180;


  return (
    <>
      <div className="w-full h-full flex flex-col overflow-hidden">
        <div className="flex-1 min-h-0 grid grid-cols-[2fr_1fr] p-2">
          <div className="overflow-y-auto pr-2">
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-4 gap-4 p-4">
              {assets?.map((asset: Asset, index: number) => (
                <div
                  key={asset.id || index}
                  onClick={() => handleAssetClick(asset)}
                  className="w-full aspect-[5/3] bg-white dark:bg-gray-800 rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-all duration-200"
                >
                  <AssetContainers
                    name={asset?.name}
                    symbol={asset?.symbol}
                    type={asset?.type}
                    price={asset?.current_price}
                  />
                </div>
              ))}
            </div>
          </div>

          <div className="overflow-y-auto">
            <div className="w-full h-full gap-4 flex flex-col items-center p-2">
              <div className='w-[96%] h-full'>
                <PortfolioHealthGauge score={market_score} label={'Portfolio Health'} />
              </div>

              <div className='w-[96%] h-full'>
                <RiskRadarChart />
              </div>
            </div>
          </div>
        </div>
      </div>
      {showAIModal && selectedAsset && (
        <AssetAISummary close={toggleModal} asset={selectedAsset} />
      )}
    </>
  )
}

export default Portfolio
