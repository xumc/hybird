package main

import hybird "github.com/xumc/hybird"

type video struct {
	ID        int64
	Name      string
	ViewCount int // the value of this field are fetched from another service
}

type videoWithCount struct {
	ID        int64
	ViewCount int
}

func getViewCountsFromAnotherService(_ []int64) ([]videoWithCount, error) {
	return []videoWithCount{
		{1, 100},
		{2, 1000},
	}, nil
}

func getVideosFromMyOwnDatabase() ([]video, error) {
	return []video{
		{ID: 1, Name: "Titanic"},
		{ID: 2, Name: "Kill Bill"},
	}, nil
}

func GetAllVideosInfosPreviously() ([]video, error) {
	videos, err := getVideosFromMyOwnDatabase()
	if err != nil {
		return nil, err
	}

	videoIDs := make([]int64, len(videos))
	for i, v := range videos {
		videoIDs[i] = v.ID
	}

	viewCounts, err := getViewCountsFromAnotherService(videoIDs)
	if err != nil {
		return nil, err
	}

	videoCountsMap := make(map[int64]int)
	for _, vc := range viewCounts {
		videoCountsMap[vc.ID] = vc.ViewCount
	}

	for i := range videos {
		videos[i].ViewCount = videoCountsMap[videos[i].ID]
	}

	// Finally, we get what we want.
	return videos, nil
}

func GetAllVideosInfosNow() ([]video, error) {
	videos, err := getVideosFromMyOwnDatabase()
	if err != nil {
		return nil, err
	}
	videoIDs := hybird.ExtractInt64(videos, "ID")

	viewCounts, err := getViewCountsFromAnotherService(videoIDs)
	if err != nil {
		return nil, err
	}
	videoCountsMap := hybird.MapInt64Int(viewCounts, "ID", "ViewCount")

	for i := range videos {
		videos[i].ViewCount = videoCountsMap[videos[i].ID]
	}

	return videos, nil
}

func main() {
	GetAllVideosInfosPreviously()
	GetAllVideosInfosNow()
}
