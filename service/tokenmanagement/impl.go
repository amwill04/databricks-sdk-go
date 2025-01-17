// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just TokenManagement API methods
type tokenManagementImpl struct {
	client *client.DatabricksClient
}

func (a *tokenManagementImpl) CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error) {
	var createOboTokenResponse CreateOboTokenResponse
	path := "/api/2.0/token-management/on-behalf-of/tokens"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createOboTokenResponse)
	return &createOboTokenResponse, err
}

func (a *tokenManagementImpl) Delete(ctx context.Context, request Delete) error {
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *tokenManagementImpl) Get(ctx context.Context, request Get) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	path := fmt.Sprintf("/api/2.0/token-management/tokens/%v", request.TokenId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &tokenInfo)
	return &tokenInfo, err
}

func (a *tokenManagementImpl) List(ctx context.Context, request List) (*ListTokensResponse, error) {
	var listTokensResponse ListTokensResponse
	path := "/api/2.0/token-management/tokens"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTokensResponse)
	return &listTokensResponse, err
}
