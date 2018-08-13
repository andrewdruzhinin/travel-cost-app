package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/andrewdruzhinin/travel-cost-app/distance_service/trippb"

	"googlemaps.github.io/maps"
)

// DistanceInfo :
type DistanceInfo struct {
	Distance int
	Duration float64
}

// GetDistanceInfo :
func GetDistanceInfo(tp *pb.Trip) (*DistanceInfo, error) {
	APIKey := os.Getenv("GOOGLEMAPS_API_KEY")

	c, err := maps.NewClient(maps.WithAPIKey(APIKey))
	if err != nil {
		return nil, err
	}

	startPoint := tp.GetStartPoint()
	endPoint := tp.GetEndPoint()

	origin := fmt.Sprintf("%v,%v", startPoint.GetLat(), startPoint.GetLong())
	destination := fmt.Sprintf("%v,%v", endPoint.GetLat(), endPoint.GetLong())

	r := &maps.DistanceMatrixRequest{
		Origins:      []string{origin},
		Destinations: []string{destination},
		Units:        `UnitsMetric`,
		Language:     "Ru",
	}

	resp, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return nil, err
	}

	distance := resp.Rows[0].Elements[0].Distance.Meters
	duration := resp.Rows[0].Elements[0].Duration.Minutes()
	distInfo := &DistanceInfo{distance, duration}

	return distInfo, nil
}
