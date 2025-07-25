/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type VolumesAPI interface {

	/*
		CreateVolume Create a new volume

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return VolumesAPICreateVolumeRequest
	*/
	CreateVolume(ctx context.Context) VolumesAPICreateVolumeRequest

	// CreateVolumeExecute executes the request
	//  @return VolumeDto
	CreateVolumeExecute(r VolumesAPICreateVolumeRequest) (*VolumeDto, *http.Response, error)

	/*
		DeleteVolume Delete volume

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param volumeId ID of the volume
		@return VolumesAPIDeleteVolumeRequest
	*/
	DeleteVolume(ctx context.Context, volumeId string) VolumesAPIDeleteVolumeRequest

	// DeleteVolumeExecute executes the request
	DeleteVolumeExecute(r VolumesAPIDeleteVolumeRequest) (*http.Response, error)

	/*
		GetVolume Get volume details

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param volumeId ID of the volume
		@return VolumesAPIGetVolumeRequest
	*/
	GetVolume(ctx context.Context, volumeId string) VolumesAPIGetVolumeRequest

	// GetVolumeExecute executes the request
	//  @return VolumeDto
	GetVolumeExecute(r VolumesAPIGetVolumeRequest) (*VolumeDto, *http.Response, error)

	/*
		GetVolumeByName Get volume details by name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param name Name of the volume
		@return VolumesAPIGetVolumeByNameRequest
	*/
	GetVolumeByName(ctx context.Context, name string) VolumesAPIGetVolumeByNameRequest

	// GetVolumeByNameExecute executes the request
	//  @return VolumeDto
	GetVolumeByNameExecute(r VolumesAPIGetVolumeByNameRequest) (*VolumeDto, *http.Response, error)

	/*
		ListVolumes List all volumes

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return VolumesAPIListVolumesRequest
	*/
	ListVolumes(ctx context.Context) VolumesAPIListVolumesRequest

	// ListVolumesExecute executes the request
	//  @return []VolumeDto
	ListVolumesExecute(r VolumesAPIListVolumesRequest) ([]VolumeDto, *http.Response, error)
}

// VolumesAPIService VolumesAPI service
type VolumesAPIService service

type VolumesAPICreateVolumeRequest struct {
	ctx                    context.Context
	ApiService             VolumesAPI
	createVolume           *CreateVolume
	xDaytonaOrganizationID *string
}

func (r VolumesAPICreateVolumeRequest) CreateVolume(createVolume CreateVolume) VolumesAPICreateVolumeRequest {
	r.createVolume = &createVolume
	return r
}

// Use with JWT to specify the organization ID
func (r VolumesAPICreateVolumeRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) VolumesAPICreateVolumeRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r VolumesAPICreateVolumeRequest) Execute() (*VolumeDto, *http.Response, error) {
	return r.ApiService.CreateVolumeExecute(r)
}

/*
CreateVolume Create a new volume

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return VolumesAPICreateVolumeRequest
*/
func (a *VolumesAPIService) CreateVolume(ctx context.Context) VolumesAPICreateVolumeRequest {
	return VolumesAPICreateVolumeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VolumeDto
func (a *VolumesAPIService) CreateVolumeExecute(r VolumesAPICreateVolumeRequest) (*VolumeDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VolumeDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.CreateVolume")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createVolume == nil {
		return localVarReturnValue, nil, reportError("createVolume is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	// body params
	localVarPostBody = r.createVolume
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type VolumesAPIDeleteVolumeRequest struct {
	ctx                    context.Context
	ApiService             VolumesAPI
	volumeId               string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r VolumesAPIDeleteVolumeRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) VolumesAPIDeleteVolumeRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r VolumesAPIDeleteVolumeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVolumeExecute(r)
}

/*
DeleteVolume Delete volume

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param volumeId ID of the volume
	@return VolumesAPIDeleteVolumeRequest
*/
func (a *VolumesAPIService) DeleteVolume(ctx context.Context, volumeId string) VolumesAPIDeleteVolumeRequest {
	return VolumesAPIDeleteVolumeRequest{
		ApiService: a,
		ctx:        ctx,
		volumeId:   volumeId,
	}
}

// Execute executes the request
func (a *VolumesAPIService) DeleteVolumeExecute(r VolumesAPIDeleteVolumeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.DeleteVolume")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeId"+"}", url.PathEscape(parameterValueToString(r.volumeId, "volumeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type VolumesAPIGetVolumeRequest struct {
	ctx                    context.Context
	ApiService             VolumesAPI
	volumeId               string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r VolumesAPIGetVolumeRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) VolumesAPIGetVolumeRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r VolumesAPIGetVolumeRequest) Execute() (*VolumeDto, *http.Response, error) {
	return r.ApiService.GetVolumeExecute(r)
}

/*
GetVolume Get volume details

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param volumeId ID of the volume
	@return VolumesAPIGetVolumeRequest
*/
func (a *VolumesAPIService) GetVolume(ctx context.Context, volumeId string) VolumesAPIGetVolumeRequest {
	return VolumesAPIGetVolumeRequest{
		ApiService: a,
		ctx:        ctx,
		volumeId:   volumeId,
	}
}

// Execute executes the request
//
//	@return VolumeDto
func (a *VolumesAPIService) GetVolumeExecute(r VolumesAPIGetVolumeRequest) (*VolumeDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VolumeDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.GetVolume")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeId"+"}", url.PathEscape(parameterValueToString(r.volumeId, "volumeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type VolumesAPIGetVolumeByNameRequest struct {
	ctx                    context.Context
	ApiService             VolumesAPI
	name                   string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r VolumesAPIGetVolumeByNameRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) VolumesAPIGetVolumeByNameRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r VolumesAPIGetVolumeByNameRequest) Execute() (*VolumeDto, *http.Response, error) {
	return r.ApiService.GetVolumeByNameExecute(r)
}

/*
GetVolumeByName Get volume details by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of the volume
	@return VolumesAPIGetVolumeByNameRequest
*/
func (a *VolumesAPIService) GetVolumeByName(ctx context.Context, name string) VolumesAPIGetVolumeByNameRequest {
	return VolumesAPIGetVolumeByNameRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return VolumeDto
func (a *VolumesAPIService) GetVolumeByNameExecute(r VolumesAPIGetVolumeByNameRequest) (*VolumeDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VolumeDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.GetVolumeByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/by-name/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type VolumesAPIListVolumesRequest struct {
	ctx                    context.Context
	ApiService             VolumesAPI
	xDaytonaOrganizationID *string
	includeDeleted         *bool
}

// Use with JWT to specify the organization ID
func (r VolumesAPIListVolumesRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) VolumesAPIListVolumesRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

// Include deleted volumes in the response
func (r VolumesAPIListVolumesRequest) IncludeDeleted(includeDeleted bool) VolumesAPIListVolumesRequest {
	r.includeDeleted = &includeDeleted
	return r
}

func (r VolumesAPIListVolumesRequest) Execute() ([]VolumeDto, *http.Response, error) {
	return r.ApiService.ListVolumesExecute(r)
}

/*
ListVolumes List all volumes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return VolumesAPIListVolumesRequest
*/
func (a *VolumesAPIService) ListVolumes(ctx context.Context) VolumesAPIListVolumesRequest {
	return VolumesAPIListVolumesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []VolumeDto
func (a *VolumesAPIService) ListVolumesExecute(r VolumesAPIListVolumesRequest) ([]VolumeDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VolumeDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.ListVolumes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDeleted", r.includeDeleted, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
