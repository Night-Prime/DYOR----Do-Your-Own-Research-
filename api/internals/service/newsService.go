package service

import (
	"fmt"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

type NewsService struct {
	newsAPIClient NewsAPIClient
}

func NewNewsService(newsAPIClient NewsAPIClient) *NewsService {
	return &NewsService{
		newsAPIClient: newsAPIClient,
	}
}

func (n *NewsService) GetNews() ([]*models.News, error) {
	fmt.Println("Fetching news from the Alpha Vantage Service")
	fmt.Println("---------------------------------------------\n")

	news, err := n.newsAPIClient.GetNewsData()
	if err != nil {
		return nil, &errors.DatabaseError{
            Message: "Error Fetching",
            Err: err,
        }
	}

	return news, nil
}

func (n *NewsService) CalculateNewsSentiment() (float64, error) {
    fmt.Println("Calculating news sentiment")
    fmt.Println("---------------------------------------------\n")

    news, err := n.newsAPIClient.GetNewsData()
    if err != nil {
        return 0, &errors.DatabaseError{
			Message: "Error Fetching News Data",
			Err:     err,
		}
    }

    if len(news) == 0 {
        return 0, &errors.ValidationError{
			Message: "No news data available",
		}
    }

    var (
        totalSentiment float64
        validScores    int
    )

    for _, item := range news {
        if item.OverallSentimentLabel == "" {
            continue
        }
        totalSentiment += item.OverallSentimentScore
        validScores++
        fmt.Println("Score: ", item.OverallSentimentScore, validScores)
    }

    if validScores == 0 {
        return 0, &errors.ValidationError{
			Message: "No valid sentiment scores found",
		}
    }

    averageSentiment := totalSentiment / float64(validScores)
    return averageSentiment, nil
}

func (n *NewsService) GetTopGainersLosers() (*models.TickerUpdates, error) {
    fmt.Println("Fetching Top Gainers & Losers Update from the Alpha Vantage Service")
	fmt.Println("---------------------------------------------\n")

    topGainersLosers, err := n.newsAPIClient.GetTopGainersLosers()
    if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Error Fetching",
            Err: err,
        }
    }

    return topGainersLosers, nil
}