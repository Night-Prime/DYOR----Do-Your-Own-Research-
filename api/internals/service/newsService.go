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

func (n *NewsService) GetNews() (*models.News, error) {
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