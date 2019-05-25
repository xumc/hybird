# hybird

As far as I know, the whole world is coming to MicroService. That means for each developer, we need to mash data together from multiple data sources. It might be our own database, or other services. If you are working on mash the data together using golang, it might be a nightmare. There are so many similar code. This is a utils for you to reduce repetitive code.
   
## Comparison

For example, we are working on a business to mesh video data together, and the view count comes from another service. We might write the below code previously.
```
type video struct {
	ID        int64
	Name      string
	ViewCount int
}

func GetAllVideosInfos() []video {
    videos := getVideosFromMyOwnDatabase()

    videoIDs := make([]int64, len(videos)
    for i, v := range videos {
        videoIDs[i] = v.ID
    }

    viewCounts := getViewCountsFromAnotherService(videoIDs)

    videoCountsMap = make(map[int64]int)
    for _, vc := range videoCounts {
        videoCountsMap[vc.ID] = vc.ViewCount 
    }

    for i := videos {
      videos[i] = videoCountsMap[videos[i].ID]
    }

    // Finally, we get what we want.
    return videos
}

```

With this utils library, we can write the code as below.
```
type video struct {
	ID        int64
	Name      string
	ViewCount int
}

func GetAllVideosInfos() []video {
    videos := getVideosFromMyOwnDatabase()
    videoIDs := hybird.ExtractInt64(videos, "ID")
    
    viewCounts := getViewCountsFromAnotherService(videoIDs)
    videoCountsMap :=hybird.MapInt64Int(videoCounts, "ID", "ViewCount")
    
    for i := videos {
        videos[i].VideoCount = videoCountsMap[videos[i].ID]
    }
    
    return videos
}
```