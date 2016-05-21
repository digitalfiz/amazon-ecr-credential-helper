// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/ecr/ecriface (interfaces: ECRAPI)

package mock_ecriface

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	ecr "github.com/aws/aws-sdk-go/service/ecr"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ECRAPI interface
type MockECRAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockECRAPIRecorder
}

// Recorder for MockECRAPI (not exported)
type _MockECRAPIRecorder struct {
	mock *MockECRAPI
}

func NewMockECRAPI(ctrl *gomock.Controller) *MockECRAPI {
	mock := &MockECRAPI{ctrl: ctrl}
	mock.recorder = &_MockECRAPIRecorder{mock}
	return mock
}

func (_m *MockECRAPI) EXPECT() *_MockECRAPIRecorder {
	return _m.recorder
}

func (_m *MockECRAPI) BatchCheckLayerAvailability(_param0 *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	ret := _m.ctrl.Call(_m, "BatchCheckLayerAvailability", _param0)
	ret0, _ := ret[0].(*ecr.BatchCheckLayerAvailabilityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchCheckLayerAvailability(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchCheckLayerAvailability", arg0)
}

func (_m *MockECRAPI) BatchCheckLayerAvailabilityRequest(_param0 *ecr.BatchCheckLayerAvailabilityInput) (*request.Request, *ecr.BatchCheckLayerAvailabilityOutput) {
	ret := _m.ctrl.Call(_m, "BatchCheckLayerAvailabilityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.BatchCheckLayerAvailabilityOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchCheckLayerAvailabilityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchCheckLayerAvailabilityRequest", arg0)
}

func (_m *MockECRAPI) BatchDeleteImage(_param0 *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
	ret := _m.ctrl.Call(_m, "BatchDeleteImage", _param0)
	ret0, _ := ret[0].(*ecr.BatchDeleteImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchDeleteImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchDeleteImage", arg0)
}

func (_m *MockECRAPI) BatchDeleteImageRequest(_param0 *ecr.BatchDeleteImageInput) (*request.Request, *ecr.BatchDeleteImageOutput) {
	ret := _m.ctrl.Call(_m, "BatchDeleteImageRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.BatchDeleteImageOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchDeleteImageRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchDeleteImageRequest", arg0)
}

func (_m *MockECRAPI) BatchGetImage(_param0 *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
	ret := _m.ctrl.Call(_m, "BatchGetImage", _param0)
	ret0, _ := ret[0].(*ecr.BatchGetImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchGetImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchGetImage", arg0)
}

func (_m *MockECRAPI) BatchGetImageRequest(_param0 *ecr.BatchGetImageInput) (*request.Request, *ecr.BatchGetImageOutput) {
	ret := _m.ctrl.Call(_m, "BatchGetImageRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.BatchGetImageOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) BatchGetImageRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchGetImageRequest", arg0)
}

func (_m *MockECRAPI) CompleteLayerUpload(_param0 *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
	ret := _m.ctrl.Call(_m, "CompleteLayerUpload", _param0)
	ret0, _ := ret[0].(*ecr.CompleteLayerUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) CompleteLayerUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompleteLayerUpload", arg0)
}

func (_m *MockECRAPI) CompleteLayerUploadRequest(_param0 *ecr.CompleteLayerUploadInput) (*request.Request, *ecr.CompleteLayerUploadOutput) {
	ret := _m.ctrl.Call(_m, "CompleteLayerUploadRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.CompleteLayerUploadOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) CompleteLayerUploadRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompleteLayerUploadRequest", arg0)
}

func (_m *MockECRAPI) CreateRepository(_param0 *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateRepository", _param0)
	ret0, _ := ret[0].(*ecr.CreateRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) CreateRepository(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRepository", arg0)
}

func (_m *MockECRAPI) CreateRepositoryRequest(_param0 *ecr.CreateRepositoryInput) (*request.Request, *ecr.CreateRepositoryOutput) {
	ret := _m.ctrl.Call(_m, "CreateRepositoryRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.CreateRepositoryOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) CreateRepositoryRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRepositoryRequest", arg0)
}

func (_m *MockECRAPI) DeleteRepository(_param0 *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteRepository", _param0)
	ret0, _ := ret[0].(*ecr.DeleteRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DeleteRepository(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRepository", arg0)
}

func (_m *MockECRAPI) DeleteRepositoryPolicy(_param0 *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteRepositoryPolicy", _param0)
	ret0, _ := ret[0].(*ecr.DeleteRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DeleteRepositoryPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRepositoryPolicy", arg0)
}

func (_m *MockECRAPI) DeleteRepositoryPolicyRequest(_param0 *ecr.DeleteRepositoryPolicyInput) (*request.Request, *ecr.DeleteRepositoryPolicyOutput) {
	ret := _m.ctrl.Call(_m, "DeleteRepositoryPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.DeleteRepositoryPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DeleteRepositoryPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRepositoryPolicyRequest", arg0)
}

func (_m *MockECRAPI) DeleteRepositoryRequest(_param0 *ecr.DeleteRepositoryInput) (*request.Request, *ecr.DeleteRepositoryOutput) {
	ret := _m.ctrl.Call(_m, "DeleteRepositoryRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.DeleteRepositoryOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DeleteRepositoryRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRepositoryRequest", arg0)
}

func (_m *MockECRAPI) DescribeRepositories(_param0 *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeRepositories", _param0)
	ret0, _ := ret[0].(*ecr.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DescribeRepositories(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeRepositories", arg0)
}

func (_m *MockECRAPI) DescribeRepositoriesRequest(_param0 *ecr.DescribeRepositoriesInput) (*request.Request, *ecr.DescribeRepositoriesOutput) {
	ret := _m.ctrl.Call(_m, "DescribeRepositoriesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.DescribeRepositoriesOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) DescribeRepositoriesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeRepositoriesRequest", arg0)
}

func (_m *MockECRAPI) GetAuthorizationToken(_param0 *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizationToken", _param0)
	ret0, _ := ret[0].(*ecr.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetAuthorizationToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizationToken", arg0)
}

func (_m *MockECRAPI) GetAuthorizationTokenRequest(_param0 *ecr.GetAuthorizationTokenInput) (*request.Request, *ecr.GetAuthorizationTokenOutput) {
	ret := _m.ctrl.Call(_m, "GetAuthorizationTokenRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.GetAuthorizationTokenOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetAuthorizationTokenRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizationTokenRequest", arg0)
}

func (_m *MockECRAPI) GetDownloadUrlForLayer(_param0 *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
	ret := _m.ctrl.Call(_m, "GetDownloadUrlForLayer", _param0)
	ret0, _ := ret[0].(*ecr.GetDownloadUrlForLayerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetDownloadUrlForLayer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDownloadUrlForLayer", arg0)
}

func (_m *MockECRAPI) GetDownloadUrlForLayerRequest(_param0 *ecr.GetDownloadUrlForLayerInput) (*request.Request, *ecr.GetDownloadUrlForLayerOutput) {
	ret := _m.ctrl.Call(_m, "GetDownloadUrlForLayerRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.GetDownloadUrlForLayerOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetDownloadUrlForLayerRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDownloadUrlForLayerRequest", arg0)
}

func (_m *MockECRAPI) GetRepositoryPolicy(_param0 *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "GetRepositoryPolicy", _param0)
	ret0, _ := ret[0].(*ecr.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetRepositoryPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRepositoryPolicy", arg0)
}

func (_m *MockECRAPI) GetRepositoryPolicyRequest(_param0 *ecr.GetRepositoryPolicyInput) (*request.Request, *ecr.GetRepositoryPolicyOutput) {
	ret := _m.ctrl.Call(_m, "GetRepositoryPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.GetRepositoryPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) GetRepositoryPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRepositoryPolicyRequest", arg0)
}

func (_m *MockECRAPI) InitiateLayerUpload(_param0 *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
	ret := _m.ctrl.Call(_m, "InitiateLayerUpload", _param0)
	ret0, _ := ret[0].(*ecr.InitiateLayerUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) InitiateLayerUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InitiateLayerUpload", arg0)
}

func (_m *MockECRAPI) InitiateLayerUploadRequest(_param0 *ecr.InitiateLayerUploadInput) (*request.Request, *ecr.InitiateLayerUploadOutput) {
	ret := _m.ctrl.Call(_m, "InitiateLayerUploadRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.InitiateLayerUploadOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) InitiateLayerUploadRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InitiateLayerUploadRequest", arg0)
}

func (_m *MockECRAPI) ListImages(_param0 *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListImages", _param0)
	ret0, _ := ret[0].(*ecr.ListImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) ListImages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListImages", arg0)
}

func (_m *MockECRAPI) ListImagesRequest(_param0 *ecr.ListImagesInput) (*request.Request, *ecr.ListImagesOutput) {
	ret := _m.ctrl.Call(_m, "ListImagesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.ListImagesOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) ListImagesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListImagesRequest", arg0)
}

func (_m *MockECRAPI) PutImage(_param0 *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
	ret := _m.ctrl.Call(_m, "PutImage", _param0)
	ret0, _ := ret[0].(*ecr.PutImageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) PutImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutImage", arg0)
}

func (_m *MockECRAPI) PutImageRequest(_param0 *ecr.PutImageInput) (*request.Request, *ecr.PutImageOutput) {
	ret := _m.ctrl.Call(_m, "PutImageRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.PutImageOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) PutImageRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutImageRequest", arg0)
}

func (_m *MockECRAPI) SetRepositoryPolicy(_param0 *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "SetRepositoryPolicy", _param0)
	ret0, _ := ret[0].(*ecr.SetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) SetRepositoryPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRepositoryPolicy", arg0)
}

func (_m *MockECRAPI) SetRepositoryPolicyRequest(_param0 *ecr.SetRepositoryPolicyInput) (*request.Request, *ecr.SetRepositoryPolicyOutput) {
	ret := _m.ctrl.Call(_m, "SetRepositoryPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.SetRepositoryPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) SetRepositoryPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRepositoryPolicyRequest", arg0)
}

func (_m *MockECRAPI) UploadLayerPart(_param0 *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
	ret := _m.ctrl.Call(_m, "UploadLayerPart", _param0)
	ret0, _ := ret[0].(*ecr.UploadLayerPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) UploadLayerPart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UploadLayerPart", arg0)
}

func (_m *MockECRAPI) UploadLayerPartRequest(_param0 *ecr.UploadLayerPartInput) (*request.Request, *ecr.UploadLayerPartOutput) {
	ret := _m.ctrl.Call(_m, "UploadLayerPartRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ecr.UploadLayerPartOutput)
	return ret0, ret1
}

func (_mr *_MockECRAPIRecorder) UploadLayerPartRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UploadLayerPartRequest", arg0)
}
