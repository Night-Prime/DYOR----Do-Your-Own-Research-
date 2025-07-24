"use client"
import React, { useMemo } from 'react'
import AnimatedTab from '../components/AnimatedTab';
import PortfolioHealthGauge from '../shared/charts/PortfolioHealthGauge';
import RiskRadarChart from '../shared/charts/RadarChart';
import VerticalBarChart from '../shared/charts/VerticalBarChart';
import { formatAssetData } from '../utils/helper';
import AISummary from '../components/AISummary';


// This is a container that wraps most of the chart & graph components, as it provides them data
// TODO End this Prop Drilling Disaster below (use redux)
const ChartContainer = () => {
    const score = useMemo(() => {
        const raw = localStorage.getItem('user/sentiment-score');
        const rawRecommendedData = localStorage.getItem('user/assets-file?file=recommended');
        return {
            sentiment: JSON.parse(raw || '{}'),
            recommended: JSON.parse(rawRecommendedData || '{}')
        }
    }, [])
    const market_score = (score.sentiment.sentiment_score ?? 0) * 180;
    const recommended = formatAssetData(score.recommended.data, true) || [];


    return (
        <div className="w-full h-full flex flex-col">
            <AnimatedTab tab1Label='Market Sentiments' chartComponent1={<PortfolioHealthGauge score={market_score} label={''} />} tab2Label='Top Performers' chartComponent2={<VerticalBarChart recommended={recommended} />} />
            <AnimatedTab tab1Label='Insight & Analysis' chartComponent1={<AISummary />} tab2Label='Risk Radar' chartComponent2={<RiskRadarChart />} />
            {/* <AnimatedTab tab1Label='Portfolio Health' chartComponent1={<PortfolioHealthGauge score={market_score} />} tab2Label='Recommendations' chartComponent2={<VerticalBarChart recommended={recommended} />} /> */}
        </div>
    )
};

export default ChartContainer
