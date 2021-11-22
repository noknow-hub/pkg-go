//////////////////////////////////////////////////////////////////////
// dataforseo.go
//////////////////////////////////////////////////////////////////////
package dataforseo

import (
    "encoding/base64"
//    myAuthentication "github.com/noknow-hub/pkg-go/vendor/dataforseo/client/authentication"
    myConstant "github.com/noknow-hub/pkg-go/vendor/dataforseo/constant"
//    myConstant "dataforseo/constant"
//    myGoogle "github.com/noknow-hub/pkg-go/vendor/dataforseo/client/v3/serp/google"
)


//////////////////////////////////////////////////////////////////////
// New OrganicTaskPostClient object.
//////////////////////////////////////////////////////////////////////
func NewOrganicTaskPostClient(login, password string, isSandbox bool) string {
    cred := base64.StdEncoding.EncodeToString([]byte(login + ":" + password))
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX
    }
    return cred
}
