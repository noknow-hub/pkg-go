//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package account

import (
    "time"
)

const (
    COL_ADDRESS = "address"
    COL_ADDRESS_CITY = "address_city"
    COL_ADDRESS_COUNTRY_CODE = "address_country_code"
    COL_ADDRESS_OPTION = "address_option"
    COL_ADDRESS_POSTAL_CODE = "address_postal_code"
    COL_AGE = "age"
    COL_BIOGRAPHY = "biography"
    COL_BIRTHDAY = "birthday"
    COL_CREATED_AT = "created_at"
    COL_EMAIL = "email"
    COL_FIRST_NAME = "first_name"
    COL_GENDER = "gender"
    COL_ID = "id"
    COL_LAST_LOGGED_IN = "last_logged_in"
    COL_LAST_NAME = "last_name"
    COL_MIDDLE_NAME = "middle_name"
    COL_NATIONALITY_CODE = "nationality_code"
    COL_NICKNAME = "nickname"
    COL_PASSWORD = "password"
    COL_PHONE_NUMBER = "phone_number"
    COL_PUBLISHABLE_TOKEN = "publishable_token"
    COL_SECRET_TOKEN = "secret_token"
    COL_STATUS = "status"
    COL_UPDATED_AT = "updated_at"
    NUM_COLS = 24
    TABLE_NAME = "accounts"
    VAL_STATUS_BUSINESS_ACTIVE = "21"
    VAL_STATUS_BUSINESS_DEACTIVE = "29"
    VAL_STATUS_BUSINESS_INITIAL_STAGE = "20"
    VAL_STATUS_DEFAULT = "0"
    VAL_STATUS_MANAGER_ACTIVE = "31"
    VAL_STATUS_MANAGER_DEACTIVE = "39"
    VAL_STATUS_MANAGER_INITIAL_STAGE = "30"
    VAL_STATUS_PERSONAL_ACTIVE = "11"
    VAL_STATUS_PERSONAL_DEACTIVE = "19"
    VAL_STATUS_PERSONAL_INITIAL_STAGE = "10"
)

type Account struct {
    Id string
    Email string
    Status string
    NickName string
    FirstName string
    LastName string
    MiddleName string
    NationalityCode string
    PhoneNumber string
    Age string
    Gender string
    Birthday time.Time
    Biography string
    Password string
    PublishableToken string
    SecretToken string
    AddressCountryCode string
    AddressPostalCode string
    AddressCity string
    Address string
    AddressOption string
    LastLoggedIn time.Time
    CreatedAt time.Time
    UpdatedAt time.Time
}
