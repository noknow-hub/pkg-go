//////////////////////////////////////////////////////////////////////
// dataforseo.go
//////////////////////////////////////////////////////////////////////
package dataforseo

import (
    myV3KeywordsDataGoogleAdsSearchVolumeTaskGet "github.com/noknow-hub/pkg-go/dataforseo/request/v3/keywords_data/google_ads/search_volume/task_get"
    myV3KeywordsDataGoogleAdsSearchVolumeTaskPost "github.com/noknow-hub/pkg-go/dataforseo/request/v3/keywords_data/google_ads/search_volume/task_post"
    myV3SerpGoogleOrganicStandard "github.com/noknow-hub/pkg-go/dataforseo/client/v3/serp/google/organic/standard"
)


//////////////////////////////////////////////////////////////////////
// New client for "v3 > Keywords Data API > Google Ads > Search Volume > Task GET"
//////////////////////////////////////////////////////////////////////
func NewClientV3KeywordsDataGoogleAdsSearchVolumeTaskGet(login, password string, isSandbox bool, id string) *myV3KeywordsDataGoogleAdsSearchVolumeTaskGet.Client {
    return myV3KeywordsDataGoogleAdsSearchVolumeTaskGet.NewClient(login, password, isSandbox, id)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3 > Keywords Data API > Google Ads > Search Volume > Task POST" 
//////////////////////////////////////////////////////////////////////
func NewClientV3KeywordsDataGoogleAdsSearchVolumeTaskPost(login, password string, isSandbox bool) *myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.Client {
    return myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.NewClient(login, password, isSandbox)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: task: get: advanced".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTaskGetAdvanced(login, password string, isSandbox bool, id string) *myV3SerpGoogleOrganicStandard.TaskGetAdvancedClient {
    return myV3SerpGoogleOrganicStandard.NewTaskGetAdvancedClient(login, password, isSandbox, id)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: task: post".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTaskPost(login, password string, isSandbox bool) *myV3SerpGoogleOrganicStandard.TaskPostClient {
    return myV3SerpGoogleOrganicStandard.NewTaskPostClient(login, password, isSandbox)
}


//////////////////////////////////////////////////////////////////////
// New client for "v3: serp: google: organic: tasks: ready".
//////////////////////////////////////////////////////////////////////
func NewClientV3SerpGoogleOrganicTasksReady(login, password string, isSandbox bool) *myV3SerpGoogleOrganicStandard.TasksReadyClient {
    return myV3SerpGoogleOrganicStandard.NewTasksReadyClient(login, password, isSandbox)
}


//////////////////////////////////////////////////////////////////////
// New data with "location_name" for "v3 > Keywords Data API > Google Ads > Search Volume > Task POST"
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationNameV3KeywordsDataGoogleAdsSearchVolumeTaskPost(keywords []string, locationName string) *myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.Data {
    return myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.NewDataWithLocationName(keywords, locationName)
}


//////////////////////////////////////////////////////////////////////
// New data with "location_code" for "v3 > Keywords Data API > Google Ads > Search Volume > Task POST"
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCodeV3KeywordsDataGoogleAdsSearchVolumeTaskPost(keywords []string, locationCode int) *myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.Data {
    return myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.NewDataWithLocationCode(keywords, locationCode)
}


//////////////////////////////////////////////////////////////////////
// New data with "location_coordinate" for "v3 > Keywords Data API > Google Ads > Search Volume > Task POST"
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCoordinateV3KeywordsDataGoogleAdsSearchVolumeTaskPost(keywords []string, locationCoordinate string) *myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.Data {
    return myV3KeywordsDataGoogleAdsSearchVolumeTaskPost.NewDataWithLocationCoordinate(keywords, locationCoordinate)
}





