// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetOperationResultsInput
type ListStackSetOperationResultsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous request didn't return all of the remaining results, the response
	// object's NextToken parameter value is set to a token. To retrieve the next
	// set of results, call ListStackSetOperationResults again and assign that token
	// to the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The ID of the stack set operation.
	//
	// OperationId is a required field
	OperationId *string `min:"1" type:"string" required:"true"`

	// The name or unique ID of the stack set that you want to get operation results
	// for.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListStackSetOperationResultsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackSetOperationResultsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackSetOperationResultsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.OperationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OperationId"))
	}
	if s.OperationId != nil && len(*s.OperationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationId", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetOperationResultsOutput
type ListStackSetOperationResultsOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all results, NextToken is set to a token. To
	// retrieve the next set of results, call ListOperationResults again and assign
	// that token to the request object's NextToken parameter. If there are no remaining
	// results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// A list of StackSetOperationResultSummary structures that contain information
	// about the specified operation results, for accounts and regions that are
	// included in the operation.
	Summaries []StackSetOperationResultSummary `type:"list"`
}

// String returns the string representation
func (s ListStackSetOperationResultsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStackSetOperationResults = "ListStackSetOperationResults"

// ListStackSetOperationResultsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about the results of a stack set operation.
//
//    // Example sending a request using ListStackSetOperationResultsRequest.
//    req := client.ListStackSetOperationResultsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackSetOperationResults
func (c *Client) ListStackSetOperationResultsRequest(input *ListStackSetOperationResultsInput) ListStackSetOperationResultsRequest {
	op := &aws.Operation{
		Name:       opListStackSetOperationResults,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListStackSetOperationResultsInput{}
	}

	req := c.newRequest(op, input, &ListStackSetOperationResultsOutput{})
	return ListStackSetOperationResultsRequest{Request: req, Input: input, Copy: c.ListStackSetOperationResultsRequest}
}

// ListStackSetOperationResultsRequest is the request type for the
// ListStackSetOperationResults API operation.
type ListStackSetOperationResultsRequest struct {
	*aws.Request
	Input *ListStackSetOperationResultsInput
	Copy  func(*ListStackSetOperationResultsInput) ListStackSetOperationResultsRequest
}

// Send marshals and sends the ListStackSetOperationResults API request.
func (r ListStackSetOperationResultsRequest) Send(ctx context.Context) (*ListStackSetOperationResultsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStackSetOperationResultsResponse{
		ListStackSetOperationResultsOutput: r.Request.Data.(*ListStackSetOperationResultsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListStackSetOperationResultsResponse is the response type for the
// ListStackSetOperationResults API operation.
type ListStackSetOperationResultsResponse struct {
	*ListStackSetOperationResultsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStackSetOperationResults request.
func (r *ListStackSetOperationResultsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
