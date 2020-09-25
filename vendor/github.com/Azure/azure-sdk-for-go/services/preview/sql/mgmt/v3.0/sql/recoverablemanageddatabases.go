package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecoverableManagedDatabasesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type RecoverableManagedDatabasesClient struct {
	BaseClient
}

// NewRecoverableManagedDatabasesClient creates an instance of the RecoverableManagedDatabasesClient client.
func NewRecoverableManagedDatabasesClient(subscriptionID string) RecoverableManagedDatabasesClient {
	return NewRecoverableManagedDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecoverableManagedDatabasesClientWithBaseURI creates an instance of the RecoverableManagedDatabasesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewRecoverableManagedDatabasesClientWithBaseURI(baseURI string, subscriptionID string) RecoverableManagedDatabasesClient {
	return RecoverableManagedDatabasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a recoverable managed database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client RecoverableManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string) (result RecoverableManagedDatabase, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoverableManagedDatabasesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, recoverableDatabaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecoverableManagedDatabasesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":     autorest.Encode("path", managedInstanceName),
		"recoverableDatabaseName": autorest.Encode("path", recoverableDatabaseName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases/{recoverableDatabaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecoverableManagedDatabasesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecoverableManagedDatabasesClient) GetResponder(resp *http.Response) (result RecoverableManagedDatabase, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInstance gets a list of recoverable managed databases.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client RecoverableManagedDatabasesClient) ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result RecoverableManagedDatabaseListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoverableManagedDatabasesClient.ListByInstance")
		defer func() {
			sc := -1
			if result.rmdlr.Response.Response != nil {
				sc = result.rmdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInstanceNextResults
	req, err := client.ListByInstancePreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "ListByInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.rmdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "ListByInstance", resp, "Failure sending request")
		return
	}

	result.rmdlr, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "ListByInstance", resp, "Failure responding to request")
	}
	if result.rmdlr.hasNextLink() && result.rmdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByInstancePreparer prepares the ListByInstance request.
func (client RecoverableManagedDatabasesClient) ListByInstancePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstanceSender sends the ListByInstance request. The method will close the
// http.Response Body if it receives an error.
func (client RecoverableManagedDatabasesClient) ListByInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByInstanceResponder handles the response to the ListByInstance request. The method always
// closes the http.Response Body.
func (client RecoverableManagedDatabasesClient) ListByInstanceResponder(resp *http.Response) (result RecoverableManagedDatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstanceNextResults retrieves the next set of results, if any.
func (client RecoverableManagedDatabasesClient) listByInstanceNextResults(ctx context.Context, lastResults RecoverableManagedDatabaseListResult) (result RecoverableManagedDatabaseListResult, err error) {
	req, err := lastResults.recoverableManagedDatabaseListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "listByInstanceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "listByInstanceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecoverableManagedDatabasesClient", "listByInstanceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstanceComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecoverableManagedDatabasesClient) ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result RecoverableManagedDatabaseListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoverableManagedDatabasesClient.ListByInstance")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInstance(ctx, resourceGroupName, managedInstanceName)
	return
}
