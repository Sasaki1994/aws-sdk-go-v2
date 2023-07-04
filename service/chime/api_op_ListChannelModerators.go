// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the moderators for a channel. The x-amz-chime-bearer request header
// is mandatory. Use the AppInstanceUserArn of the user that makes the API call as
// the value in the header. This API is is no longer supported and will not be
// updated. We recommend using the latest version, ListChannelModerators (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_ListChannelModerators.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by ListChannelModerators in the Amazon Chime SDK Messaging
// Namespace
func (c *Client) ListChannelModerators(ctx context.Context, params *ListChannelModeratorsInput, optFns ...func(*Options)) (*ListChannelModeratorsOutput, error) {
	if params == nil {
		params = &ListChannelModeratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelModerators", params, optFns, c.addOperationListChannelModeratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelModeratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelModeratorsInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	// The maximum number of moderators that you want returned.
	MaxResults *int32

	// The token passed by previous API calls until all requested moderators are
	// returned.
	NextToken *string

	noSmithyDocumentSerde
}

type ListChannelModeratorsOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The information about and names of each moderator.
	ChannelModerators []types.ChannelModeratorSummary

	// The token passed by previous API calls until all requested moderators are
	// returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelModeratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelModerators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelModerators{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListChannelModeratorsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListChannelModeratorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelModerators(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListChannelModeratorsMiddleware struct {
}

func (*endpointPrefix_opListChannelModeratorsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListChannelModeratorsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListChannelModeratorsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListChannelModeratorsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListChannelModeratorsAPIClient is a client that implements the
// ListChannelModerators operation.
type ListChannelModeratorsAPIClient interface {
	ListChannelModerators(context.Context, *ListChannelModeratorsInput, ...func(*Options)) (*ListChannelModeratorsOutput, error)
}

var _ ListChannelModeratorsAPIClient = (*Client)(nil)

// ListChannelModeratorsPaginatorOptions is the paginator options for
// ListChannelModerators
type ListChannelModeratorsPaginatorOptions struct {
	// The maximum number of moderators that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelModeratorsPaginator is a paginator for ListChannelModerators
type ListChannelModeratorsPaginator struct {
	options   ListChannelModeratorsPaginatorOptions
	client    ListChannelModeratorsAPIClient
	params    *ListChannelModeratorsInput
	nextToken *string
	firstPage bool
}

// NewListChannelModeratorsPaginator returns a new ListChannelModeratorsPaginator
func NewListChannelModeratorsPaginator(client ListChannelModeratorsAPIClient, params *ListChannelModeratorsInput, optFns ...func(*ListChannelModeratorsPaginatorOptions)) *ListChannelModeratorsPaginator {
	if params == nil {
		params = &ListChannelModeratorsInput{}
	}

	options := ListChannelModeratorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelModeratorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelModeratorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelModerators page.
func (p *ListChannelModeratorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelModeratorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListChannelModerators(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListChannelModerators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListChannelModerators",
	}
}
