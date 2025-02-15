// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/files/s3.go

// Package mock_files is a generated GoMock package.
package mock_files

import (
	io "io"
	reflect "reflect"
	time "time"

	request "github.com/aws/aws-sdk-go/aws/request"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	gomock "github.com/golang/mock/gomock"
	files "github.com/redhatinsights/edge-api/pkg/services/files"
)

// MockS3ClientAPI is a mock of S3ClientAPI interface.
type MockS3ClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientAPIMockRecorder
}

// MockS3ClientAPIMockRecorder is the mock recorder for MockS3ClientAPI.
type MockS3ClientAPIMockRecorder struct {
	mock *MockS3ClientAPI
}

// NewMockS3ClientAPI creates a new mock instance.
func NewMockS3ClientAPI(ctrl *gomock.Controller) *MockS3ClientAPI {
	mock := &MockS3ClientAPI{ctrl: ctrl}
	mock.recorder = &MockS3ClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ClientAPI) EXPECT() *MockS3ClientAPIMockRecorder {
	return m.recorder
}

// GetObject mocks base method.
func (m *MockS3ClientAPI) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", input)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3ClientAPIMockRecorder) GetObject(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3ClientAPI)(nil).GetObject), input)
}

// GetObjectRequest mocks base method.
func (m *MockS3ClientAPI) GetObjectRequest(input *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectRequest", input)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*s3.GetObjectOutput)
	return ret0, ret1
}

// GetObjectRequest indicates an expected call of GetObjectRequest.
func (mr *MockS3ClientAPIMockRecorder) GetObjectRequest(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectRequest", reflect.TypeOf((*MockS3ClientAPI)(nil).GetObjectRequest), input)
}

// PutObject mocks base method.
func (m *MockS3ClientAPI) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", input)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3ClientAPIMockRecorder) PutObject(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3ClientAPI)(nil).PutObject), input)
}

// MockS3DownloaderAPI is a mock of S3DownloaderAPI interface.
type MockS3DownloaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3DownloaderAPIMockRecorder
}

// MockS3DownloaderAPIMockRecorder is the mock recorder for MockS3DownloaderAPI.
type MockS3DownloaderAPIMockRecorder struct {
	mock *MockS3DownloaderAPI
}

// NewMockS3DownloaderAPI creates a new mock instance.
func NewMockS3DownloaderAPI(ctrl *gomock.Controller) *MockS3DownloaderAPI {
	mock := &MockS3DownloaderAPI{ctrl: ctrl}
	mock.recorder = &MockS3DownloaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3DownloaderAPI) EXPECT() *MockS3DownloaderAPIMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockS3DownloaderAPI) Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{w, input}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Download", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockS3DownloaderAPIMockRecorder) Download(w, input interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{w, input}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockS3DownloaderAPI)(nil).Download), varargs...)
}

// MockS3UploaderAPI is a mock of S3UploaderAPI interface.
type MockS3UploaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3UploaderAPIMockRecorder
}

// MockS3UploaderAPIMockRecorder is the mock recorder for MockS3UploaderAPI.
type MockS3UploaderAPIMockRecorder struct {
	mock *MockS3UploaderAPI
}

// NewMockS3UploaderAPI creates a new mock instance.
func NewMockS3UploaderAPI(ctrl *gomock.Controller) *MockS3UploaderAPI {
	mock := &MockS3UploaderAPI{ctrl: ctrl}
	mock.recorder = &MockS3UploaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3UploaderAPI) EXPECT() *MockS3UploaderAPIMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *MockS3UploaderAPI) Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{input}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(*s3manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockS3UploaderAPIMockRecorder) Upload(input interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{input}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockS3UploaderAPI)(nil).Upload), varargs...)
}

// MockS3RequestAPI is a mock of S3RequestAPI interface.
type MockS3RequestAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3RequestAPIMockRecorder
}

// MockS3RequestAPIMockRecorder is the mock recorder for MockS3RequestAPI.
type MockS3RequestAPIMockRecorder struct {
	mock *MockS3RequestAPI
}

// NewMockS3RequestAPI creates a new mock instance.
func NewMockS3RequestAPI(ctrl *gomock.Controller) *MockS3RequestAPI {
	mock := &MockS3RequestAPI{ctrl: ctrl}
	mock.recorder = &MockS3RequestAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3RequestAPI) EXPECT() *MockS3RequestAPIMockRecorder {
	return m.recorder
}

// Presign mocks base method.
func (m *MockS3RequestAPI) Presign(expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Presign", expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Presign indicates an expected call of Presign.
func (mr *MockS3RequestAPIMockRecorder) Presign(expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Presign", reflect.TypeOf((*MockS3RequestAPI)(nil).Presign), expire)
}

// MockRequestPreSignerAPI is a mock of RequestPreSignerAPI interface.
type MockRequestPreSignerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRequestPreSignerAPIMockRecorder
}

// MockRequestPreSignerAPIMockRecorder is the mock recorder for MockRequestPreSignerAPI.
type MockRequestPreSignerAPIMockRecorder struct {
	mock *MockRequestPreSignerAPI
}

// NewMockRequestPreSignerAPI creates a new mock instance.
func NewMockRequestPreSignerAPI(ctrl *gomock.Controller) *MockRequestPreSignerAPI {
	mock := &MockRequestPreSignerAPI{ctrl: ctrl}
	mock.recorder = &MockRequestPreSignerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestPreSignerAPI) EXPECT() *MockRequestPreSignerAPIMockRecorder {
	return m.recorder
}

// Presign mocks base method.
func (m *MockRequestPreSignerAPI) Presign(req files.S3RequestAPI, expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Presign", req, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Presign indicates an expected call of Presign.
func (mr *MockRequestPreSignerAPIMockRecorder) Presign(req, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Presign", reflect.TypeOf((*MockRequestPreSignerAPI)(nil).Presign), req, expire)
}

// MockS3ClientInterface is a mock of S3ClientInterface interface.
type MockS3ClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientInterfaceMockRecorder
}

// MockS3ClientInterfaceMockRecorder is the mock recorder for MockS3ClientInterface.
type MockS3ClientInterfaceMockRecorder struct {
	mock *MockS3ClientInterface
}

// NewMockS3ClientInterface creates a new mock instance.
func NewMockS3ClientInterface(ctrl *gomock.Controller) *MockS3ClientInterface {
	mock := &MockS3ClientInterface{ctrl: ctrl}
	mock.recorder = &MockS3ClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ClientInterface) EXPECT() *MockS3ClientInterfaceMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockS3ClientInterface) Download(file io.WriterAt, bucket, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", file, bucket, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockS3ClientInterfaceMockRecorder) Download(file, bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockS3ClientInterface)(nil).Download), file, bucket, key)
}

// GetObject mocks base method.
func (m *MockS3ClientInterface) GetObject(bucket, key string) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", bucket, key)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3ClientInterfaceMockRecorder) GetObject(bucket, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3ClientInterface)(nil).GetObject), bucket, key)
}

// GetSignedURL mocks base method.
func (m *MockS3ClientInterface) GetSignedURL(bucket, key string, expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignedURL", bucket, key, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignedURL indicates an expected call of GetSignedURL.
func (mr *MockS3ClientInterfaceMockRecorder) GetSignedURL(bucket, key, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignedURL", reflect.TypeOf((*MockS3ClientInterface)(nil).GetSignedURL), bucket, key, expire)
}

// PutObject mocks base method.
func (m *MockS3ClientInterface) PutObject(file io.ReadSeeker, bucket, key, acl string) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", file, bucket, key, acl)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3ClientInterfaceMockRecorder) PutObject(file, bucket, key, acl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3ClientInterface)(nil).PutObject), file, bucket, key, acl)
}

// Upload mocks base method.
func (m *MockS3ClientInterface) Upload(file io.Reader, bucket, key, acl string) (*s3manager.UploadOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", file, bucket, key, acl)
	ret0, _ := ret[0].(*s3manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockS3ClientInterfaceMockRecorder) Upload(file, bucket, key, acl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockS3ClientInterface)(nil).Upload), file, bucket, key, acl)
}
