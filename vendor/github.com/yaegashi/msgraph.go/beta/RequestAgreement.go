// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AgreementRequestBuilder is request builder for Agreement
type AgreementRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementRequest
func (b *AgreementRequestBuilder) Request() *AgreementRequest {
	return &AgreementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementRequest is request for Agreement
type AgreementRequest struct{ BaseRequest }

// Get performs GET request for Agreement
func (r *AgreementRequest) Get(ctx context.Context) (resObj *Agreement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Agreement
func (r *AgreementRequest) Update(ctx context.Context, reqObj *Agreement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Agreement
func (r *AgreementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AgreementAcceptanceRequestBuilder is request builder for AgreementAcceptance
type AgreementAcceptanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementAcceptanceRequest
func (b *AgreementAcceptanceRequestBuilder) Request() *AgreementAcceptanceRequest {
	return &AgreementAcceptanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementAcceptanceRequest is request for AgreementAcceptance
type AgreementAcceptanceRequest struct{ BaseRequest }

// Get performs GET request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Get(ctx context.Context) (resObj *AgreementAcceptance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Update(ctx context.Context, reqObj *AgreementAcceptance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AgreementFileRequestBuilder is request builder for AgreementFile
type AgreementFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementFileRequest
func (b *AgreementFileRequestBuilder) Request() *AgreementFileRequest {
	return &AgreementFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementFileRequest is request for AgreementFile
type AgreementFileRequest struct{ BaseRequest }

// Get performs GET request for AgreementFile
func (r *AgreementFileRequest) Get(ctx context.Context) (resObj *AgreementFile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AgreementFile
func (r *AgreementFileRequest) Update(ctx context.Context, reqObj *AgreementFile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AgreementFile
func (r *AgreementFileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
