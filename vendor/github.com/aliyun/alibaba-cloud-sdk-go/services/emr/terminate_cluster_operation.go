package emr

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

// TerminateClusterOperation invokes the emr.TerminateClusterOperation API synchronously
// api document: https://help.aliyun.com/api/emr/terminateclusteroperation.html
func (client *Client) TerminateClusterOperation(request *TerminateClusterOperationRequest) (response *TerminateClusterOperationResponse, err error) {
	response = CreateTerminateClusterOperationResponse()
	err = client.DoAction(request, response)
	return
}

// TerminateClusterOperationWithChan invokes the emr.TerminateClusterOperation API asynchronously
// api document: https://help.aliyun.com/api/emr/terminateclusteroperation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateClusterOperationWithChan(request *TerminateClusterOperationRequest) (<-chan *TerminateClusterOperationResponse, <-chan error) {
	responseChan := make(chan *TerminateClusterOperationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TerminateClusterOperation(request)
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

// TerminateClusterOperationWithCallback invokes the emr.TerminateClusterOperation API asynchronously
// api document: https://help.aliyun.com/api/emr/terminateclusteroperation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateClusterOperationWithCallback(request *TerminateClusterOperationRequest, callback func(response *TerminateClusterOperationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TerminateClusterOperationResponse
		var err error
		defer close(result)
		response, err = client.TerminateClusterOperation(request)
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

// TerminateClusterOperationRequest is the request struct for api TerminateClusterOperation
type TerminateClusterOperationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OperationId     string           `position:"Query" name:"OperationId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// TerminateClusterOperationResponse is the response struct for api TerminateClusterOperation
type TerminateClusterOperationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTerminateClusterOperationRequest creates a request to invoke TerminateClusterOperation API
func CreateTerminateClusterOperationRequest() (request *TerminateClusterOperationRequest) {
	request = &TerminateClusterOperationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "TerminateClusterOperation", "emr", "openAPI")
	return
}

// CreateTerminateClusterOperationResponse creates a response to parse from TerminateClusterOperation response
func CreateTerminateClusterOperationResponse() (response *TerminateClusterOperationResponse) {
	response = &TerminateClusterOperationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}