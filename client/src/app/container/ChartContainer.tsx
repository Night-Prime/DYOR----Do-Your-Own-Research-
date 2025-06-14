"use client"
import React, { useMemo } from 'react'
import { useAppSelector } from '../hooks/hook'
import AnimatedTab from '../components/AnimatedTab';
import PortfolioHealthGauge from '../shared/charts/PortfolioHealthGauge';
import TopPerformingChart from '../shared/charts/TopPerformingChart';
import RiskRadarChart from '../shared/charts/RadarChart';
import NegativeAreaChart from '../shared/charts/NegativeCharts';
import VerticalBarChart from '../shared/charts/VerticalBarChart';
import { recommended } from '../data/asset';
import { getAssetData } from '../utils/service-calls';


// This is a container that wraps most of the chart & graph components, as it provides them data

const ChartContainer = () => {
    // const assets = useAppSelector((state) => state.portfolio.assets);
    // const assetDetails = assets.map(asset => {
    //     return {
    //         type: asset.type,
    //         symbol: asset.symbol,
    //     }
    // })
    // console.log("Assets: ", assetDetails);
    const score = useMemo(() => {
        const raw = localStorage.getItem('user/sentiment-score');
        const rawTickers = localStorage.getItem('user/top-gainers-losers');
        return {
            sentiment: JSON.parse(raw || '{}'),
            tickers : JSON.parse(rawTickers || '{}')
        }
    }, [])
    const market_score = (score.sentiment.sentiment_score ?? 0) * 180;
    const topGainers = score.tickers.top_gainers;
    const topLosers = score.tickers.top_losers


    return (
        <div className="w-full h-full">
            <AnimatedTab tab1Label='Market Sentiments' chartComponent1={<PortfolioHealthGauge score={market_score} />} tab2Label='Top 5 Performers' chartComponent2={<TopPerformingChart tickers={topGainers} />} />
            <AnimatedTab tab1Label='Risk Radar' chartComponent1={<RiskRadarChart />} tab2Label='Worst 5 Performers' chartComponent2={<NegativeAreaChart tickers={topLosers} />} />
            <AnimatedTab tab1Label='Portfolio Health' chartComponent1={<PortfolioHealthGauge score={market_score} />} tab2Label='Recommendations' chartComponent2={<VerticalBarChart />} />
        </div>
    )
};

export default ChartContainer
