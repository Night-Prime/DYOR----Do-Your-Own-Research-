-- Migration script for the News model and related tables

-- Create the News table
CREATE TABLE news (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    time_published TIMESTAMP NOT NULL,
    authors TEXT[] NOT NULL,
    summary TEXT NOT NULL,
    banner_image TEXT NOT NULL,
    source TEXT NOT NULL,
    category_within_source TEXT NOT NULL,
    source_domain TEXT NOT NULL,
    overall_sentiment_score FLOAT NOT NULL,
    overall_sentiment_label TEXT NOT NULL
);

-- Create the Topic table
CREATE TABLE topics (
    id SERIAL PRIMARY KEY,
    news_id TEXT NOT NULL REFERENCES news(id) ON DELETE CASCADE,
    topic TEXT NOT NULL,
    relevance_score TEXT NOT NULL
);

-- Create the TickerSentiment table
CREATE TABLE ticker_sentiments (
    id SERIAL PRIMARY KEY,
    news_id TEXT NOT NULL REFERENCES news(id) ON DELETE CASCADE,
    ticker TEXT NOT NULL,
    relevance_score TEXT NOT NULL,
    ticker_sentiment_score TEXT NOT NULL,
    ticker_sentiment_label TEXT NOT NULL
);