//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package serp

import (
    "time"
)

const (
    COL_CREATED_AT = "created_at"
    COL_COUNTRY_CODE = "country_code"
    COL_DEVICE = "device"
    COL_ID = "id"
    COL_KEYWORD = "keyword"
    COL_LANG_CODE = "lang_code"
    COL_NUM_OF_SEARCHES_FOR_KEYWORD = "num_of_searches_for_keyword"
    COL_SEARCH_ENGINE = "search_engine"
    COL_SEARCH_ENGINE_TYPE = "search_engine_type"
    COL_TOTAL_RESULTS = "total_results"
    COL_UPDATED_AT = "updated_at"
    NUM_COLS = 11
    TABLE_NAME = "serps"
    VAL_DEVICE_DESKTOP = "1"
    VAL_DEVICE_MOBILE = "2"
    VAL_DIFFICULTY_EASY = "1"
    VAL_DIFFICULTY_MEDIUM = "2"
    VAL_DIFFICULTY_HARD = "3"
    VAL_DIFFICULTY_VERY_HARD = "4"
)

type Serp struct {
    Id string
    Keyword string
    CountryCode string
    LangCode string
    Device string
    TotalResults string
    SearchEngine string
    SearchEngineType string
    NumOfSearchesForKeyword string
    CreatedAt time.Time
    UpdatedAt time.Time
}
