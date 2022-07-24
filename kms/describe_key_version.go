package kms

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
	"github.com/zhangdapeng520/zdpgo_nacos/sdk/requests"
	"github.com/zhangdapeng520/zdpgo_nacos/sdk/responses"
)

// DescribeKeyVersion invokes the kms.DescribeKeyVersion API synchronously
// api document: https://help.aliyun.com/api/kms/describekeyversion.html
func (client *Client) DescribeKeyVersion(request *DescribeKeyVersionRequest) (response *DescribeKeyVersionResponse, err error) {
	response = CreateDescribeKeyVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKeyVersionWithChan invokes the kms.DescribeKeyVersion API asynchronously
// api document: https://help.aliyun.com/api/kms/describekeyversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKeyVersionWithChan(request *DescribeKeyVersionRequest) (<-chan *DescribeKeyVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeKeyVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKeyVersion(request)
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

// DescribeKeyVersionWithCallback invokes the kms.DescribeKeyVersion API asynchronously
// api document: https://help.aliyun.com/api/kms/describekeyversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeKeyVersionWithCallback(request *DescribeKeyVersionRequest, callback func(response *DescribeKeyVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKeyVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeKeyVersion(request)
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

// DescribeKeyVersionRequest is the request struct for api DescribeKeyVersion
type DescribeKeyVersionRequest struct {
	*requests.RpcRequest
	KeyVersionId string `position:"Query" name:"KeyVersionId"`
	KeyId        string `position:"Query" name:"KeyId"`
}

// DescribeKeyVersionResponse is the response struct for api DescribeKeyVersion
type DescribeKeyVersionResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	KeyVersion KeyVersion `json:"KeyVersion" xml:"KeyVersion"`
}

// CreateDescribeKeyVersionRequest creates a request to invoke DescribeKeyVersion API
func CreateDescribeKeyVersionRequest() (request *DescribeKeyVersionRequest) {
	request = &DescribeKeyVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DescribeKeyVersion", "kms", "openAPI")
	return
}

// CreateDescribeKeyVersionResponse creates a response to parse from DescribeKeyVersion response
func CreateDescribeKeyVersionResponse() (response *DescribeKeyVersionResponse) {
	response = &DescribeKeyVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
