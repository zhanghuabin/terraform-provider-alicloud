package dds

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

// DescribeActiveOperationTask invokes the dds.DescribeActiveOperationTask API synchronously
// api document: https://help.aliyun.com/api/dds/describeactiveoperationtask.html
func (client *Client) DescribeActiveOperationTask(request *DescribeActiveOperationTaskRequest) (response *DescribeActiveOperationTaskResponse, err error) {
	response = CreateDescribeActiveOperationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveOperationTaskWithChan invokes the dds.DescribeActiveOperationTask API asynchronously
// api document: https://help.aliyun.com/api/dds/describeactiveoperationtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskWithChan(request *DescribeActiveOperationTaskRequest) (<-chan *DescribeActiveOperationTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveOperationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveOperationTask(request)
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

// DescribeActiveOperationTaskWithCallback invokes the dds.DescribeActiveOperationTask API asynchronously
// api document: https://help.aliyun.com/api/dds/describeactiveoperationtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActiveOperationTaskWithCallback(request *DescribeActiveOperationTaskRequest, callback func(response *DescribeActiveOperationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveOperationTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveOperationTask(request)
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

// DescribeActiveOperationTaskRequest is the request struct for api DescribeActiveOperationTask
type DescribeActiveOperationTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskType             string           `position:"Query" name:"TaskType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	IsHistory            requests.Integer `position:"Query" name:"IsHistory"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Region               string           `position:"Query" name:"Region"`
}

// DescribeActiveOperationTaskResponse is the response struct for api DescribeActiveOperationTask
type DescribeActiveOperationTaskResponse struct {
	*responses.BaseResponse
	RequestId        string      `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageSize         int         `json:"PageSize" xml:"PageSize"`
	PageNumber       int         `json:"PageNumber" xml:"PageNumber"`
	Items            []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeActiveOperationTaskRequest creates a request to invoke DescribeActiveOperationTask API
func CreateDescribeActiveOperationTaskRequest() (request *DescribeActiveOperationTaskRequest) {
	request = &DescribeActiveOperationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeActiveOperationTask", "dds", "openAPI")
	return
}

// CreateDescribeActiveOperationTaskResponse creates a response to parse from DescribeActiveOperationTask response
func CreateDescribeActiveOperationTaskResponse() (response *DescribeActiveOperationTaskResponse) {
	response = &DescribeActiveOperationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}