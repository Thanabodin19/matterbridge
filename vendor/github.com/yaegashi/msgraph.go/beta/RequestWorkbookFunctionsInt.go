// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsIntRequestBuilder struct{ BaseRequestBuilder }

// Int action undocumented
func (b *WorkbookFunctionsRequestBuilder) Int(reqObj *WorkbookFunctionsIntRequestParameter) *WorkbookFunctionsIntRequestBuilder {
	bb := &WorkbookFunctionsIntRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/int"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsIntRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsIntRequestBuilder) Request() *WorkbookFunctionsIntRequest {
	return &WorkbookFunctionsIntRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsIntRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsIntRateRequestBuilder struct{ BaseRequestBuilder }

// IntRate action undocumented
func (b *WorkbookFunctionsRequestBuilder) IntRate(reqObj *WorkbookFunctionsIntRateRequestParameter) *WorkbookFunctionsIntRateRequestBuilder {
	bb := &WorkbookFunctionsIntRateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/intRate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsIntRateRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsIntRateRequestBuilder) Request() *WorkbookFunctionsIntRateRequest {
	return &WorkbookFunctionsIntRateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsIntRateRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}