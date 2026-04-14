package service

import (
	"context"
	"net/http"
	"time"

	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/client"
	"github.com/AkinbulejoSamson/samson-HNG-stage-zero.git/internal/dto"
)

type GenderService interface {
	ClassifyName(ctx context.Context, name string) (*dto.ProcessedData, int, error)
}

type genderService struct{}

func NewGenderService() GenderService {
	return &genderService{}
}

func (g genderService) ClassifyName(ctx context.Context, name string) (*dto.ProcessedData, int, error) {
	raw, err := client.FetchGenderizeRawData(ctx, name)
	if err != nil {
		if err.Error() == "no prediction available" {
			return nil, http.StatusOK, err
		}
		return nil, http.StatusBadGateway, err
	}

	isConfident := raw.Probability >= 0.7 && raw.Count >= 100

	result := &dto.ProcessedData{
		Name:        name,
		Gender:      raw.Gender,
		Probability: raw.Probability,
		SampleSize:  raw.Count,
		IsConfident: isConfident,
		ProcessedAt: time.Now().UTC().Format(time.RFC3339),
	}

	return result, http.StatusOK, nil
}
