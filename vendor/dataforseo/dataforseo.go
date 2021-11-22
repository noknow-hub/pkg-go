//////////////////////////////////////////////////////////////////////
// dataforseo.go
//////////////////////////////////////////////////////////////////////
package dataforseo

import (
    myAuthentication "github.com/noknow-hub/pkg-go/vendor/dataforseo/client/authentication"
    myConstant "github.com/noknow-hub/pkg-go/vendor/dataforseo/constant"

)


//////////////////////////////////////////////////////////////////////
// New OrganicTaskPostClient object.
//////////////////////////////////////////////////////////////////////
func NewSerpGoogleOrganicTaskPostClient(login, password string, isSandbox bool) *OrganicTaskPostClient {
    endpointUrl := myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3
    if isSandbox {
        endpointUrl = myConstant.ENDPOINT_SERP_GOOGLE_ORGANIC_TASK_POST_V3_SANDBOX
    }
    return &OrganicTaskPostClient{
        Authentication: myAuthentication.NewAuthentication(login, password),
        EndpointUrl: endpointUrl,
    }
}
