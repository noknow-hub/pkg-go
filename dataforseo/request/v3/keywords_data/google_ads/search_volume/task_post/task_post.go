//////////////////////////////////////////////////////////////////////
// task_post.go
//////////////////////////////////////////////////////////////////////
package task_post

import (
    "encoding/json"
    "regexp"
    "strings"
    "time"
    myAuthentication "github.com/noknow-hub/pkg-go/dataforseo/authentication"
    myConstant "github.com/noknow-hub/pkg-go/dataforseo/constant"
    myHttpClient "github.com/noknow-hub/pkg-go/http/client"
    myResult "github.com/noknow-hub/pkg-go/dataforseo/response/v3/keywords_data/google_ads/search_volume/task_post"
)

const (
    LIMIT_NUM_OF_TASKS_PER_REQ = myConstant.LIMIT_NUM_OF_TASKS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
)

var (
    NumOfApiCalls = 0
    LastCalledAt time.Time
    ApiCallQueues []time.Time
    Reg = regexp.MustCompile("[!！@＠%％^＾()（）=＝{};；：~〜`｀<>＜＞?？\\|｜,、，。…．Ⅱ⇨:ｦ-ﾝ○※＊「」［］【】｛｝‘’“”〈〉〔〕《》*🟡×]")
)

type Client struct {
    *myAuthentication.Authentication
    EndpointUrl string
}

type Data struct {
    DateFrom string             `json:"date_from,omitempty"`
    DateTo string               `json:"date_to,omitempty"`
    Keywords []string           `json:"keywords,omitempty"`
    LanguageName string         `json:"language_name,omitempty"`
    LanguageCode string         `json:"language_code,omitempty"`
    LocationName string         `json:"location_name,omitempty"`
    LocationCode int            `json:"location_code,omitempty"`
    LocationCoordinate string   `json:"location_coordinate,omitempty"`
    PingbackUrl string          `json:"pingback_url,omitempty"`
    PostbackUrl string          `json:"postback_url,omitempty"`
    SearchPartners bool         `json:"search_partners,omitempty"`
    SortBy string               `json:"sort_by,omitempty"`
    Tag string                  `json:"tag,omitempty"`
}


//////////////////////////////////////////////////////////////////////
// Clear NumOfApiCalls.
//////////////////////////////////////////////////////////////////////
func ClearNumOfApiCalls() {
    NumOfApiCalls = 0
}


//////////////////////////////////////////////////////////////////////
// Is over the number of tasks.
//////////////////////////////////////////////////////////////////////
func IsOverNumOfTasks(datas []*Data) bool {
    return len(datas) > myConstant.LIMIT_NUM_OF_TASKS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
}


//////////////////////////////////////////////////////////////////////
// New Client object.
//////////////////////////////////////////////////////////////////////
func NewClient(login, password string, isSandbox bool) *Client {
    endpointUrl := myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3_SANDBOX
    }
    return &Client{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_name".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationName(keywords []string, locationName string) *Data {
    return &Data{
        Keywords: keywords,
        LocationName: locationName,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_code".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCode(keywords []string, locationCode int) *Data {
    return &Data{
        Keywords: keywords,
        LocationCode: locationCode,
    }
}


//////////////////////////////////////////////////////////////////////
// New Data with "location_coordinate".
//////////////////////////////////////////////////////////////////////
func NewDataWithLocationCoordinate(keywords []string, locationCoordinate string) *Data {
    return &Data{
        Keywords: keywords,
        LocationCoordinate: locationCoordinate,
    }
}


//////////////////////////////////////////////////////////////////////
// Optimize data.
//////////////////////////////////////////////////////////////////////
func OptData(datas []*Data) []*Data {
    var optDatas []*Data
    for _, data := range datas {
        if len(data.Keywords) <= myConstant.LIMIT_NUM_OF_KEYWORDS_PER_TASK_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 {
            optDatas = append(optDatas, &Data{
                DateFrom: data.DateFrom,
                DateTo: data.DateTo,
                Keywords: data.Keywords,
                LanguageName: data.LanguageName,
                LanguageCode: data.LanguageCode,
                LocationName: data.LocationName,
                LocationCode: data.LocationCode,
                LocationCoordinate: data.LocationCoordinate,
                PingbackUrl: data.PingbackUrl,
                PostbackUrl: data.PostbackUrl,
                SearchPartners: data.SearchPartners,
                SortBy: data.SortBy,
                Tag: data.Tag,
            })
        } else {
            var index int
            for {
                if len(data.Keywords) < myConstant.LIMIT_NUM_OF_KEYWORDS_PER_TASK_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 {
                    index = len(data.Keywords)
                } else {
                    index = myConstant.LIMIT_NUM_OF_KEYWORDS_PER_TASK_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3
                }
                optDatas = append(optDatas, &Data{
                    DateFrom: data.DateFrom,
                    DateTo: data.DateTo,
                    Keywords: data.Keywords[:index],
                    LanguageName: data.LanguageName,
                    LanguageCode: data.LanguageCode,
                    LocationName: data.LocationName,
                    LocationCode: data.LocationCode,
                    LocationCoordinate: data.LocationCoordinate,
                    PingbackUrl: data.PingbackUrl,
                    PostbackUrl: data.PostbackUrl,
                    SearchPartners: data.SearchPartners,
                    SortBy: data.SortBy,
                    Tag: data.Tag,
                })
                data.Keywords = data.Keywords[index:]
                if len(data.Keywords) == 0 {
                    break
                }
            }
        }
    }
    return optDatas
}


//////////////////////////////////////////////////////////////////////
// Optimize keyword.
//////////////////////////////////////////////////////////////////////
func OptimazeKeyword(keyword string) string {
    s := Reg.ReplaceAllString(keyword, " ")
    ss := strings.Split(s, " ")
    if len(ss) > myConstant.LIMIT_NUM_OF_WARDS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 {
        s = strings.Join(ss[:myConstant.LIMIT_NUM_OF_WARDS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3], " ")
    } else {
        s = strings.Join(ss, " ")
    }
    return s
}


//////////////////////////////////////////////////////////////////////
// Valid keyword.
//////////////////////////////////////////////////////////////////////
func ValidKeyword(keyword string) bool {
    return !Reg.MatchString(keyword)
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *Client) Run(datas []*Data) (int, *myResult.Response, error) {
    minAgo := time.Now().Add(time.Minute * -1)
    for i, ApiCallQueue := range ApiCallQueues {
        if ApiCallQueue.After(minAgo) {
            ApiCallQueues = ApiCallQueues[i:]
            break
        }
    }

    if len(ApiCallQueues) > myConstant.LIMIT_NUM_OF_CALLS_PER_MIN_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 {
        time.Sleep(60 * time.Second)
        ApiCallQueues = nil
    }
    ApiCallQueues = append(ApiCallQueues, time.Now())

    jsonData, err := json.Marshal(datas)
    if err != nil {
        return 0, nil, err
    }

    httpClient := myHttpClient.NewClient(c.EndpointUrl)
    httpClient.Config.
        AddHeaderContentType(myHttpClient.CONTENT_TYPE_JSON).
        AddHeaderAuthorization("Basic " + c.Cred).
        SetJsonData(jsonData)
    resp, err := httpClient.Post()
    if err != nil {
        return 0, nil, err
    }
    var result *myResult.Response
    if err := json.Unmarshal(resp.Body, &result); err != nil {
        return 0, nil, err
    }
    result.Raw = string(resp.Body)

    return resp.StatusCode, result, nil
}


//////////////////////////////////////////////////////////////////////
// Set "date_from".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetDateFrom(dateFrom string) *Data {
    d.DateFrom = dateFrom
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "date_to".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetDateTo(dateTo string) *Data {
    d.DateTo = dateTo
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_code".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetLanguageCode(languageCode string) *Data {
    d.LanguageCode = languageCode
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "language_name".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetLanguageName(languageName string) *Data {
    d.LanguageName = languageName
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "pingback_url".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetPingbackUrl(pingbackUrl string) *Data {
    d.PingbackUrl = pingbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "postback_url".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetPostbackUrl(postbackUrl string) *Data {
    d.PostbackUrl = postbackUrl
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "search_partners".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetSearchPartners(searchPartners bool) *Data {
    d.SearchPartners = searchPartners
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "sort_by".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetSortBy(sortBy string) *Data {
    d.SortBy = sortBy
    return d
}


//////////////////////////////////////////////////////////////////////
// Set "tag".
//////////////////////////////////////////////////////////////////////
func (d *Data) SetTag(tag string) *Data {
    d.Tag = tag
    return d
}


