package csb

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

// FindApproveServiceList invokes the csb.FindApproveServiceList API synchronously
// api document: https://help.aliyun.com/api/csb/findapproveservicelist.html
func (client *Client) FindApproveServiceList(request *FindApproveServiceListRequest) (response *FindApproveServiceListResponse, err error) {
	response = CreateFindApproveServiceListResponse()
	err = client.DoAction(request, response)
	return
}

// FindApproveServiceListWithChan invokes the csb.FindApproveServiceList API asynchronously
// api document: https://help.aliyun.com/api/csb/findapproveservicelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindApproveServiceListWithChan(request *FindApproveServiceListRequest) (<-chan *FindApproveServiceListResponse, <-chan error) {
	responseChan := make(chan *FindApproveServiceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindApproveServiceList(request)
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

// FindApproveServiceListWithCallback invokes the csb.FindApproveServiceList API asynchronously
// api document: https://help.aliyun.com/api/csb/findapproveservicelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindApproveServiceListWithCallback(request *FindApproveServiceListRequest, callback func(response *FindApproveServiceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindApproveServiceListResponse
		var err error
		defer close(result)
		response, err = client.FindApproveServiceList(request)
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

// FindApproveServiceListRequest is the request struct for api FindApproveServiceList
type FindApproveServiceListRequest struct {
	*requests.RpcRequest
	ProjectName    string           `position:"Query" name:"ProjectName"`
	ApproveLevel   string           `position:"Query" name:"ApproveLevel"`
	ShowDelService requests.Boolean `position:"Query" name:"ShowDelService"`
	CsbId          requests.Integer `position:"Query" name:"CsbId"`
	Alias          string           `position:"Query" name:"Alias"`
	ServiceName    string           `position:"Query" name:"ServiceName"`
}

// FindApproveServiceListResponse is the response struct for api FindApproveServiceList
type FindApproveServiceListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindApproveServiceListRequest creates a request to invoke FindApproveServiceList API
func CreateFindApproveServiceListRequest() (request *FindApproveServiceListRequest) {
	request = &FindApproveServiceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindApproveServiceList", "", "")
	return
}

// CreateFindApproveServiceListResponse creates a response to parse from FindApproveServiceList response
func CreateFindApproveServiceListResponse() (response *FindApproveServiceListResponse) {
	response = &FindApproveServiceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}