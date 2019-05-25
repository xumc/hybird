# hybird

As far as I know, the whole world is coming to MicroService. That means for each developer, we need to mash data together from multiple data sources. It might be our own database, or other services. If you are working on mash the data together using golang, it might be a nightmare. There are so many similar code. This is a utils for you to reduce repetitive code.
   
## Comparison

For example, we are working on a business to mesh video data together, related structs are:
```
type video struct {
	ID        int64
	Name      string
	ViewCount int // the value of this field are fetched from another service
}

type videoWithCount struct {
	ID        int64
	ViewCount int
}
```
And the view count comes from another service. We might write the below code previously.
```
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

```

With this utils library, we can write the code as below.
```
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
```