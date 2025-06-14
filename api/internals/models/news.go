package models

import (
    "time"
    "strings"
    "github.com/lib/pq"
)

type News struct {
    ID                   string         `json:"id" gorm:"primaryKey"`
    Title                string         `json:"title"`
    URL                  string         `json:"url"`
    TimePublished        CustomTime      `json:"time_published"`
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
    RelevanceScore string `json:"relevance_score"`
}

type TickerSentiment struct {
    ID                   uint    `json:"-" gorm:"primaryKey"`
    NewsID               string  `json:"-"`
    Ticker               string  `json:"ticker"`
    RelevanceScore       string `json:"relevance_score"`
    TickerSentimentScore string `json:"ticker_sentiment_score"`
    TickerSentimentLabel string  `json:"ticker_sentiment_label"`
}

type CustomTime struct {
    time.Time
}

const ctLayout = "20060102T150405"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
    s := strings.Trim(string(b), `"`)
    if s == "null" || s == "" {
        ct.Time = time.Time{}
        return nil
    }
    ct.Time, err = time.Parse(ctLayout, s)
    return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
    if ct.Time.IsZero() {
        return []byte("null"), nil
    }
    return []byte(`"` + ct.Time.Format(ctLayout) + `"`), nil
}


// this would handle updates from other service for my charts
type TickerUpdates struct {
    TopGainers   []TickerAsset `gorm:"type:jsonb" json:"top_gainers"`
    TopLosers    []TickerAsset `gorm:"type:jsonb" json:"top_losers"`
    ActiveTraded []TickerAsset `gorm:"type:jsonb" json:"active_traded"`
}

type TickerAsset struct {
    Ticker           string `gorm:"type:varchar(10)" json:"ticker"`
    Price            string `gorm:"type:varchar(20)" json:"price"`
    ChangeAmount     string `gorm:"type:varchar(20)" json:"change_amount"`
    ChangePercentage string `gorm:"type:varchar(20)" json:"change_percentage"`
    Volume           string `gorm:"type:varchar(20)" json:"volume"`
}
