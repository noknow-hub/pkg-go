//////////////////////////////////////////////////////////////////////
// constant.go
//////////////////////////////////////////////////////////////////////
package constant

const (
    API_BASE_URL_V3 = "https://api.dataforseo.com/v3"
    API_BASE_URL_V3_SANDBOX = "https://sandbox.dataforseo.com/v3"

    DEVICE_DESKTOP = "desktop"
    DEVICE_MOBILE = "mobile" 

    ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_GET_V3 = "https://api.dataforseo.com/v3/keywords_data/google_ads/search_volume/task_get/$id"
    ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_GET_V3_SANDBOX = "https://sandbox.dataforseo.com/v3/keywords_data/google_ads/search_volume/task_get/$id"
    ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 = "https://api.dataforseo.com/v3/keywords_data/google_ads/search_volume/task_post"
    ENDPOINT_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3_SANDBOX = "https://sandbox.dataforseo.com/v3/keywords_data/google_ads/search_volume/task_post"

    ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3 = "https://api.dataforseo.com/v3/serp/google/organic/tasks_ready"
    ENDPOINT_SERP_GOOGLE_ORGANIC_TASKS_READY_V3_SANDBOX = "https://sandbox.dataforseo.com/v3/serp/google/organic/tasks_ready"
    ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_GET_ADVANCED_V3 = "https://api.dataforseo.com/v3/serp/google/organic/task_get/advanced/$id"
    ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_GET_ADVANCED_V3_SANDBOX = "https://sandbox.dataforseo.com/v3/serp/google/organic/task_get/advanced/$id"
    ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3 = "https://api.dataforseo.com/v3/serp/google/organic/task_post"
    ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX = "https://sandbox.dataforseo.com/v3/serp/google/organic/task_post"

    LANG_CODE_EN = "en"
    LANG_CODE_DE = "de"
    LANG_CODE_IT = "it"
    LANG_CODE_JA = "ja"
    LANG_CODE_KO = "ko"
    LANG_CODE_ES = "es"

    LIMIT_NUM_OF_CALLS_PER_MIN_FOR_SERP_GOOGLE_ORGANIC_TASK_POST_V3 = 2000
    LIMIT_NUM_OF_TASKS_PER_REQ_FOR_SERP_GOOGLE_ORGANIC_TASK_POST_V3 = 100
    LIMIT_NUM_OF_CALLS_PER_MIN_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 = 2000
    LIMIT_NUM_OF_TASKS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 = 100
    LIMIT_NUM_OF_KEYWORDS_PER_TASK_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 = 1000
    LIMIT_NUM_OF_WARDS_PER_REQ_FOR_KEYWORDS_DATA_GOOGLE_ADS_SEARCH_VOLUME_TASK_POST_V3 = 10

    LOCATION_CODE_DE = 2276
    LOCATION_CODE_IT = 2380
    LOCATION_CODE_JP = 2392
    LOCATION_CODE_KR = 2410
    LOCATION_CODE_ES = 2724
    LOCATION_CODE_GB = 2826
    LOCATION_CODE_US = 2840

    POSTBACK_DATA_ADVANCED = "advanced"
)
