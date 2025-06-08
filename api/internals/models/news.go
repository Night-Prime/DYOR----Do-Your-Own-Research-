package models

import (
    "time"
    "github.com/lib/pq"
)

type News struct {
    ID                   string         `json:"id" gorm:"primaryKey"`
    Title                string         `json:"title"`
    URL                  string         `json:"url"`
    TimePublished        time.Time      `json:"time_published"`
    Authors              pq.StringArray `gorm:"type:text[]" json:"authors"`
    Summary              string         `json:"summary"`
    BannerImage          string         `json:"banner_image"`
    Source               string         `json:"source"`
    CategoryWithinSource string         `json:"category_within_source"`
    SourceDomain         string         `json:"source_domain"`
    Topics               []Topic        `gorm:"foreignKey:NewsID" json:"topics"`
    OverallSentimentScore float64       `json:"overall_sentiment_score"`
    OverallSentimentLabel string        `json:"overall_sentiment_label"`
    TickerSentiment      []TickerSentiment `gorm:"foreignKey:NewsID" json:"ticker_sentiment"`
}

type Topic struct {
    ID             uint    `json:"-" gorm:"primaryKey"`
    NewsID         string  `json:"-"`
    Topic          string  `json:"topic"`
    RelevanceScore float64 `json:"relevance_score" gorm:"type:numeric(10,6)"`
}

type TickerSentiment struct {
    ID                   uint    `json:"-" gorm:"primaryKey"`
    NewsID               string  `json:"-"`
    Ticker               string  `json:"ticker"`
    RelevanceScore       float64 `json:"relevance_score" gorm:"type:numeric(10,6)"`
    TickerSentimentScore float64 `json:"ticker_sentiment_score" gorm:"type:numeric(10,6)"`
    TickerSentimentLabel string  `json:"ticker_sentiment_label"`
}