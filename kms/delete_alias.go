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

// DeleteAlias invokes the kms.DeleteAlias API synchronously
// api document: https://help.aliyun.com/api/kms/deletealias.html
func (client *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
	response = CreateDeleteAliasResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAliasWithChan invokes the kms.DeleteAlias API asynchronously
// api document: https://help.aliyun.com/api/kms/deletealias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAliasWithChan(request *DeleteAliasRequest) (<-chan *DeleteAliasResponse, <-chan error) {
	responseChan := make(chan *DeleteAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAlias(request)
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

// DeleteAliasWithCallback invokes the kms.DeleteAlias API asynchronously
// api document: https://help.aliyun.com/api/kms/deletealias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAliasWithCallback(request *DeleteAliasRequest, callback func(response *DeleteAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAliasResponse
		var err error
		defer close(result)
		response, err = client.DeleteAlias(request)
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

// DeleteAliasRequest is the request struct for api DeleteAlias
type DeleteAliasRequest struct {
	*requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
}

// DeleteAliasResponse is the response struct for api DeleteAlias
type DeleteAliasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAliasRequest creates a request to invoke DeleteAlias API
func CreateDeleteAliasRequest() (request *DeleteAliasRequest) {
	request = &DeleteAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DeleteAlias", "kms", "openAPI")
	return
}

// CreateDeleteAliasResponse creates a response to parse from DeleteAlias response
func CreateDeleteAliasResponse() (response *DeleteAliasResponse) {
	response = &DeleteAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
