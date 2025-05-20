'use client'

import React from 'react'
import { useFetch } from '../hooks/useFetch'
import Preloader from '../shared/Preloader';
import { DyorAlert } from '../shared/Alert';
import Welcome from './Welcome';
import { User } from '../data/models';
import AnimatedTab from './AnimatedTab';
import InsightFeed from './InsightFeed';
import { insightSample } from '../data/Investment';
import PortfolioHealthGauge from '../shared/charts/PortfolioHealthGauge';
import VerticalBarChart from '../shared/charts/VerticalBarChart';
import RiskRadarChart from '../shared/charts/RadarChart';
import NegativeAreaChart from '../shared/charts/NegativeCharts';
import TopPerformingChart from '../shared/charts/TopPerformingChart';

const Feed = () => {
  // note: would resolve the localStorage reference error later
  const userJson = localStorage.getItem("user");
  const userDetails = userJson ? JSON.parse(userJson) : null;
  const { data, loading, error, refresh } = useFetch<User>("user/portfolio", { id: userDetails.id });
  const emptyPortfolio = data?.portfolios.length === 0

  if (loading) return <Preloader />
  if (error) {
    return <DyorAlert type="error" message={`${error}`} open={true} autoClose={true} />
  }

  return (
    <>
      {emptyPortfolio ? (
        <Welcome user={userDetails} refresh={refresh} />
      ) : (
        <div className='w-full h-full grid grid-cols-2 rounded-3xl overflow-y-hidden'>
          <div className='w-full h-full flex flex-col items-center justify-start'>
            <main className='max-h-[90dvh] overflow-y-scroll scroll-smooth pb-6'>
              {insightSample.map((feed, key) => (
                <span key={key}>
                  <InsightFeed title={feed.news} news={feed.insight} />
                </span>
              ))}
            </main>
          </div>
          <div className='w-full h-full max-h-max'>
            <div className=' w-full max-h-[90dvh] flex flex-col justify-start items-center gap-4 overflow-y-scroll scroll-smooth pb-8'>
            <AnimatedTab tab1Label='Market Sentiments' chartComponent1={<PortfolioHealthGauge />} tab2Label='Top 5 Performers' chartComponent2={<TopPerformingChart />}/>
            <AnimatedTab tab1Label='Risk Radar' chartComponent1={<RiskRadarChart/>} tab2Label='Worst 5 Performers' chartComponent2={<NegativeAreaChart />} />
            <AnimatedTab tab1Label='Portfolio Health' chartComponent1={<PortfolioHealthGauge />} tab2Label='Recommendations' chartComponent2={<VerticalBarChart />}  />
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default Feed;
