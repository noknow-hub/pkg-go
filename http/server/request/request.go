//////////////////////////////////////////////////////////////////////
// request.go
//////////////////////////////////////////////////////////////////////
package request

import (
    "net/http"
    "strings"
)


//////////////////////////////////////////////////////////////////////
// Get the lang code with 2 digits.
// Resolve a language code in the following priority order:
//   1. URL parameter (such as "lang=en")
//   2. Cookie (such as "lang")
//   3. HTTP Header ("Accept-Language")
// When it could not resolve a language code,
// will be returned defaultLangCode such as "en".
//////////////////////////////////////////////////////////////////////
func GetLangCode(r *http.Request, supportedLangCodes []string, defaultLangCode, urlQueryKey, cookieName string) string {
    // URL query parameter
    s := GetLangCodeFromUrlQuery(r, supportedLangCodes, urlQueryKey)
    if s != "" {
        return s
    }

    // Cookie
    s = GetLangCodeFromCookie(r, supportedLangCodes, cookieName)
    if s != "" {
        return s
    }

    // HTTP Header ("Accept-Language")
    s = GetLangCodeFromAcceptLanguage(r)
    if s != "" {
        return s
    }

    return defaultLangCode
}


//////////////////////////////////////////////////////////////////////
// Get the lang code with 2 digits from URL query.
//////////////////////////////////////////////////////////////////////
func GetLangCodeFromUrlQuery(r *http.Request, supportedLangCodes []string, urlQueryKey string) string {
    queryLang := r.URL.Query().Get(urlQueryKey)
    if queryLang != "" {
        for _, supportedLangCode := range supportedLangCodes {
            if supportedLangCode == queryLang {
                return supportedLangCode
            }
        }
    }
    return ""
}


//////////////////////////////////////////////////////////////////////
// Get the lang code with 2 digits from cookie.
//////////////////////////////////////////////////////////////////////
func GetLangCodeFromCookie(r *http.Request, supportedLangCodes []string, cookieName string) string {
    cookie, _ := r.Cookie(cookieName)
    if cookie != nil {
        cookieValTag := cookie.Value
        for _, supportedLangCode := range supportedLangCodes {
            if supportedLangCode == cookieValTag {
                return supportedLangCode
            }
        }
    }
    return ""
}


//////////////////////////////////////////////////////////////////////
// Get the lang code with 2 digits from HTTP Header ("Accept-Language").
//////////////////////////////////////////////////////////////////////
func GetLangCodeFromAcceptLanguage(r *http.Request) string {
    al := r.Header.Get("Accept-Language")
    if al != "" {
        return strings.Split(strings.Split(strings.Split(al, ",")[0], ";")[0], "-")[0]
    }
    return ""
}
