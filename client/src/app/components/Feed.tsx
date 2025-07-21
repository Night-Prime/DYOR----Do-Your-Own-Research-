/* eslint-disable @typescript-eslint/no-explicit-any */
'use client'

import React, { useEffect, useState } from 'react'
import {useFetch } from '../hooks/useFetch'
import Preloader from '../shared/Preloader';
import { DyorAlert } from '../shared/Alert';
import Welcome from './Welcome';
import { User } from '../data/models';
import InsightFeed from './InsightFeed';
import { useAppDispatch, useAppSelector } from '../hooks/hook';
import { setPortfolio } from '../core/portfolioSlice';
import ChartContainer from '../container/ChartContainer';

// NewsFeedList.tsx
interface NewsFeedListProps {
  news: any[];
}

const NewsFeedListComponent = ({ news }: NewsFeedListProps) => {
  if (!news?.length) {
    return (
      <div className="flex items-center justify-center h-full">
        <p className="text-gray-500">No news articles available at this time</p>
      </div>
    );
  }

  return (
    <>
      {news.map((feed) => (
        <InsightFeed
          key={feed.url}
          title={feed.title || 'No title available'}
          summary={feed.summary || 'No summary available'}
          url={feed.url || '#'}
          time={feed.time_published || ''}
          imgLink={feed.banner_image || ''}
        />
      ))}
    </>
  );
};

export const NewsFeedList = React.memo(NewsFeedListComponent);

// PortfolioContent.tsx
interface PortfolioContentProps {
  data: User;
  refresh: () => void;
}

const PortfolioContentComponent = ({ data, refresh }: PortfolioContentProps) => {
  const [news, setNews] = useState<any[]>([]);

  useEffect(() => {
    try {
      const storedNews = localStorage.getItem('user/news');
      setNews(storedNews ? JSON.parse(storedNews) : []);
    } catch (error) {
      console.error('Error parsing news:', error);
      setNews([]);
    }
  }, []);

  const emptyPortfolio = !data?.portfolios?.some(
    portfolio => portfolio?.assets?.length > 0
  );

  if (emptyPortfolio) {
    return <Welcome user={data} refresh={refresh} />;
  }

  return (
    <div className='w-full h-full grid grid-cols-[2fr_1fr] rounded-3xl overflow-hidden p-2'>
      <div className='w-full h-full flex flex-col'>
        <main className='flex-1 max-h-[88dvh] overflow-y-auto scroll-smooth p-4'>
          <NewsFeedList news={news} />
        </main>
      </div>
      <div className='w-full h-full flex flex-col'>
        <div className='flex-1 max-h-[88dvh] overflow-y-auto scroll-smooth p-4'>
          <ChartContainer />
        </div>
      </div>
    </div>
  );
};

export const PortfolioContent = React.memo(PortfolioContentComponent);

// Feed.tsx
const FeedComponent = () => {  
  const userId = useAppSelector((state) => state.auth.user?.id);
  const dispatch = useAppDispatch();

  const { data, loading, error, refresh } = useFetch<User>("user/portfolio", {
    id: userId
  });

  useEffect(() => {
    if (data?.portfolios?.[0]) {
      dispatch(setPortfolio(data.portfolios[0]));
    }
  }, [data?.portfolios, dispatch]);

  if (loading) return <Preloader />;
  if (error) {
    return <DyorAlert type="error" message={`${error}`} open autoClose />;
  }
  if (!data) return null;

  return <PortfolioContent data={data} refresh={refresh} />;
};

export const Feed = React.memo(FeedComponent);