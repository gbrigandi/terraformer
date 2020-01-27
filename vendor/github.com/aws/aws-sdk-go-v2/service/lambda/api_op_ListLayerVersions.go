// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListLayerVersionsInput struct {
	_ struct{} `type:"structure"`

	// A runtime identifier. For example, go1.x.
	CompatibleRuntime Runtime `location:"querystring" locationName:"CompatibleRuntime" type:"string" enum:"true"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// A pagination token returned by a previous call.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of versions to return.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListLayerVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListLayerVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListLayerVersionsInput"}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLayerVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LayerName != nil {
		v := *s.LayerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LayerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.CompatibleRuntime) > 0 {
		v := s.CompatibleRuntime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "CompatibleRuntime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListLayerVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of versions.
	LayerVersions []LayerVersionsListItem `type:"list"`

	// A pagination token returned when the response doesn't contain all versions.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListLayerVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLayerVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LayerVersions != nil {
		v := s.LayerVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "LayerVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListLayerVersions = "ListLayerVersions"

// ListLayerVersionsRequest returns a request value for making API operation for
// AWS Lambda.
//
// Lists the versions of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// Versions that have been deleted aren't listed. Specify a runtime identifier
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list
// only versions that indicate that they're compatible with that runtime.
//
//    // Example sending a request using ListLayerVersionsRequest.
//    req := client.ListLayerVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListLayerVersions
func (c *Client) ListLayerVersionsRequest(input *ListLayerVersionsInput) ListLayerVersionsRequest {
	op := &aws.Operation{
		Name:       opListLayerVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListLayerVersionsInput{}
	}

	req := c.newRequest(op, input, &ListLayerVersionsOutput{})
	return ListLayerVersionsRequest{Request: req, Input: input, Copy: c.ListLayerVersionsRequest}
}

// ListLayerVersionsRequest is the request type for the
// ListLayerVersions API operation.
type ListLayerVersionsRequest struct {
	*aws.Request
	Input *ListLayerVersionsInput
	Copy  func(*ListLayerVersionsInput) ListLayerVersionsRequest
}

// Send marshals and sends the ListLayerVersions API request.
func (r ListLayerVersionsRequest) Send(ctx context.Context) (*ListLayerVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLayerVersionsResponse{
		ListLayerVersionsOutput: r.Request.Data.(*ListLayerVersionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListLayerVersionsRequestPaginator returns a paginator for ListLayerVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListLayerVersionsRequest(input)
//   p := lambda.NewListLayerVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListLayerVersionsPaginator(req ListLayerVersionsRequest) ListLayerVersionsPaginator {
	return ListLayerVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListLayerVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListLayerVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListLayerVersionsPaginator struct {
	aws.Pager
}

func (p *ListLayerVersionsPaginator) CurrentPage() *ListLayerVersionsOutput {
	return p.Pager.CurrentPage().(*ListLayerVersionsOutput)
}

// ListLayerVersionsResponse is the response type for the
// ListLayerVersions API operation.
type ListLayerVersionsResponse struct {
	*ListLayerVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLayerVersions request.
func (r *ListLayerVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}