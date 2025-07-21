/* eslint-disable @typescript-eslint/no-explicit-any */
'use client'

import React, { useEffect } from 'react'
import { useFetch } from '../hooks/useFetch'
import Preloader from '../shared/Preloader';
import { DyorAlert } from '../shared/Alert';
import Welcome from './Welcome';
import { User } from '../data/models';
import InsightFeed from './InsightFeed';
import { useAppDispatch, useAppSelector } from '../hooks/hook';
import { setPortfolio } from '../core/portfolioSlice';
import ChartContainer from '../container/ChartContainer';

const Feed = () => {
  const dispatch = useAppDispatch();
  const userDetails = useAppSelector((state) => state.auth.user);

  const { data, loading, error, refresh } = useFetch<User>("user/portfolio", {
    id: userDetails?.id
  });
  
  // const {data: score} = useFetch<any>("user/sentiment-score");
  // const {data: news, loading: newsLoading, error:newsError} = useFetch<any>("user/news");
  // const {data:ticker} = useFetch<any>("user/top-gainers-losers");

  const news = JSON.parse(localStorage.getItem('user/news') || '[]');

  // Dispatch portfolio data only when it's available
  useEffect(() => {
    if (data?.portfolios?.[0]) {
      dispatch(setPortfolio(data.portfolios[0]));
    }
  }, [data, dispatch]);

  if (loading) return <Preloader />;
  if (error) {
    return <DyorAlert type="error" message={`${error}`} open={true} autoClose={true} />;
  }

  const emptyPortfolio = !data?.portfolios?.some(
    portfolio => portfolio?.assets?.length > 0
  );

  return (
    <>
      {data && emptyPortfolio ? (
        <Welcome user={data} refresh={refresh} />
      ) : (
        <div className='w-full h-full grid grid-cols-[2fr_1fr] rounded-3xl overflow-hidden p-2'>
          <div className='w-full h-full flex flex-col'>
            <main className='flex-1 max-h-[88dvh] overflow-y-auto scroll-smooth p-4'>
              {news?.length ? (
                news.map((feed: any, index: number) => (
                  <span key={feed.id || index}>
                    <InsightFeed
                      title={feed.title || 'No title available'}
                      summary={feed.summary || 'No summary available'}
                      url={feed.url || '#'}
                      time={feed.time_published || ''}
                      imgLink={feed.banner_image || ''}
                    />
                  </span>
                ))
              ) : (
                <div className="flex items-center justify-center h-full">
                  <p className="text-gray-500">No news articles available at this time</p>
                </div>
              )}
            </main>
          </div>

          <div className='w-full h-full flex flex-col'>
            <div className='flex-1 max-h-[88dvh] overflow-y-auto scroll-smooth p-4'>
              <ChartContainer />
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default Feed;
