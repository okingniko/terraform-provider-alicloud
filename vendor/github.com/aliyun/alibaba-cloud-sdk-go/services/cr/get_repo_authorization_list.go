package cr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetRepoAuthorizationList invokes the cr.GetRepoAuthorizationList API synchronously
// api document: https://help.aliyun.com/api/cr/getrepoauthorizationlist.html
func (client *Client) GetRepoAuthorizationList(request *GetRepoAuthorizationListRequest) (response *GetRepoAuthorizationListResponse, err error) {
	response = CreateGetRepoAuthorizationListResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoAuthorizationListWithChan invokes the cr.GetRepoAuthorizationList API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepoauthorizationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoAuthorizationListWithChan(request *GetRepoAuthorizationListRequest) (<-chan *GetRepoAuthorizationListResponse, <-chan error) {
	responseChan := make(chan *GetRepoAuthorizationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoAuthorizationList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetRepoAuthorizationListWithCallback invokes the cr.GetRepoAuthorizationList API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepoauthorizationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoAuthorizationListWithCallback(request *GetRepoAuthorizationListRequest, callback func(response *GetRepoAuthorizationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoAuthorizationListResponse
		var err error
		defer close(result)
		response, err = client.GetRepoAuthorizationList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetRepoAuthorizationListRequest is the request struct for api GetRepoAuthorizationList
type GetRepoAuthorizationListRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Authorize     string `position:"Query" name:"Authorize"`
}

// GetRepoAuthorizationListResponse is the response struct for api GetRepoAuthorizationList
type GetRepoAuthorizationListResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoAuthorizationListRequest creates a request to invoke GetRepoAuthorizationList API
func CreateGetRepoAuthorizationListRequest() (request *GetRepoAuthorizationListRequest) {
	request = &GetRepoAuthorizationListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoAuthorizationList", "/repos/[RepoNamespace]/[RepoName]/authorizations", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoAuthorizationListResponse creates a response to parse from GetRepoAuthorizationList response
func CreateGetRepoAuthorizationListResponse() (response *GetRepoAuthorizationListResponse) {
	response = &GetRepoAuthorizationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
