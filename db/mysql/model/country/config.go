//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package country

import (
)

const (
    COL_AR = "ar"
    COL_CONTINENT = "continent"
    COL_COUNTRY_CODE = "country_code"
    COL_DE = "de"
    COL_EN = "en"
    COL_ES = "es"
    COL_FR = "fr"
    COL_JA = "ja"
    COL_PT = "pt"
    COL_RU = "ru"
    COL_STATUS = "status"
    COL_ZH_CN = "zh_cn"
    COL_ZH_TW = "zh_tw"
    COUNTRY_CODE_AD = "ad"  // Andorra
    COUNTRY_CODE_AE = "ae"  // United Arab Emirates
    COUNTRY_CODE_AF = "af"  // Afghanistan
    COUNTRY_CODE_AG = "ag"  // Antigua and Barbuda
    COUNTRY_CODE_AI = "ai"  // Anguilla
    COUNTRY_CODE_AL = "al"  // Albania
    COUNTRY_CODE_AM = "am"  // Armenia
    COUNTRY_CODE_AO = "ao"  // Angola
    COUNTRY_CODE_AQ = "aq"  // Antarctica
    COUNTRY_CODE_AR = "ar"  // Argentina
    COUNTRY_CODE_AS = "as"  // American Samoa
    COUNTRY_CODE_AT = "at"  // Austria
    COUNTRY_CODE_AU = "au"  // Australia
    COUNTRY_CODE_AW = "aw"  // Aruba
    COUNTRY_CODE_AX = "ax"  // Åland Islands
    COUNTRY_CODE_AZ = "az"  // Azerbaijan
    COUNTRY_CODE_BA = "ba"  // Bosnia and Herzegovina
    COUNTRY_CODE_BB = "bb"  // Barbados
    COUNTRY_CODE_BD = "bd"  // Bangladesh
    COUNTRY_CODE_BE = "be"  // Belgium
    COUNTRY_CODE_BF = "bf"  // Burkina Faso
    COUNTRY_CODE_BG = "bg"  // Bulgaria
    COUNTRY_CODE_BH = "bh"  // Bahrain
    COUNTRY_CODE_BI = "bi"  // Burundi
    COUNTRY_CODE_BJ = "bj"  // Benin
    COUNTRY_CODE_BL = "bl"  // Saint Barthélemy
    COUNTRY_CODE_BM = "bm"  // Bermuda
    COUNTRY_CODE_BN = "bn"  // Brunei Darussalam
    COUNTRY_CODE_BO = "bo"  // Bolivia
    COUNTRY_CODE_BQ = "bq"  // Bonaire
    COUNTRY_CODE_BR = "br"  // Brazil
    COUNTRY_CODE_BS = "bs"  // Bahamas
    COUNTRY_CODE_BT = "bt"  // Bhutan
    COUNTRY_CODE_BV = "bv"  // Bouvet Island
    COUNTRY_CODE_BW = "bw"  // Botswana
    COUNTRY_CODE_BY = "by"  // Belarus
    COUNTRY_CODE_BZ = "bz"  // Belize
    COUNTRY_CODE_CA = "ca"  // Canada
    COUNTRY_CODE_FR = "fr"  // France
    COUNTRY_CODE_IE = "ie"  // Ireland
    COUNTRY_CODE_IN = "in"  // India
    COUNTRY_CODE_IT = "it"  // Italy
    COUNTRY_CODE_DE = "de"  // Germany
    COUNTRY_CODE_JP = "jp"  // Japan
    COUNTRY_CODE_KR = "kr"  // Korean
    COUNTRY_CODE_ES = "es"  // Spain
    COUNTRY_CODE_NZ = "nz"  // NewZealand
    COUNTRY_CODE_SG = "sg"  // Republic of Singapore
    COUNTRY_CODE_UK = "uk"  // United Kingdom
    COUNTRY_CODE_US = "us"  // United States
    COUNTRY_CODE_ZA = "za"  // Republic of South Africa
    LANG_CODE_AR = "ar"
    LANG_CODE_DE = "de"
    LANG_CODE_EN = "en"
    LANG_CODE_ES = "es"
    LANG_CODE_FR = "fr"
    LANG_CODE_JA = "ja"
    LANG_CODE_PT = "pt"
    LANG_CODE_RU = "ru"
    LANG_CODE_ZHCN = "zhcn"
    LANG_CODE_ZHTW = "zhtw"
    NUM_COLS = 13
    TABLE_NAME = "countries"
    VAL_CONTINENT_AFRICA = "1"             // Africa
    VAL_CONTINENT_ASIA = "2"               // Asia
    VAL_CONTINENT_EUROPE = "3"             // Europe
    VAL_CONTINENT_NORTH_AMERICA = "4"      // North America
    VAL_CONTINENT_SOUTH_AMERICA = "5"      // South America
    VAL_CONTINENT_AUSTRALIA_OCEANIA = "6"  // Australia / Oceania
    VAL_CONTINENT_ANTARCTICA = "7"         // Antarctica
    VAL_STATUS_INACTIVE = "0"  // inactive
    VAL_STATUS_ACTIVE = "1"    // active
)

type Country struct {
    CountryCode string
    Ar string
    De string
    En string
    Es string
    Fr string
    Ja string
    Pt string
    Ru string
    ZhCn string
    ZhTw string
    Continent string
    Status string
    Name string
}
