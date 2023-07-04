// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a monitor in Amazon CloudWatch Internet Monitor based on
// a monitor name. The information returned includes the Amazon Resource Name
// (ARN), create time, modified time, resources included in the monitor, and status
// information.
func (c *Client) GetMonitor(ctx context.Context, params *GetMonitorInput, optFns ...func(*Options)) (*GetMonitorOutput, error) {
	if params == nil {
		params = &GetMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMonitor", params, optFns, c.addOperationGetMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMonitorInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	noSmithyDocumentSerde
}

type GetMonitorOutput struct {

	// The time when the monitor was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The last time that the monitor was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	MonitorArn *string

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// The resources that have been added for the monitor. Resources are listed by
	// their Amazon Resource Names (ARNs).
	//
	// This member is required.
	Resources []string

	// The status of the monitor.
	//
	// This member is required.
	Status types.MonitorConfigState

	// The list of health event thresholds. A health event threshold percentage, for
	// performance and availability, determines the level of impact at which Amazon
	// CloudWatch Internet Monitor creates a health event when there's an internet
	// issue that affects your application end users.
	HealthEventsConfig *types.HealthEventsConfig

	// Publish internet measurements for Internet Monitor to another location, such as
	// an Amazon S3 bucket. The measurements are also published to Amazon CloudWatch
	// Logs.
	InternetMeasurementsLogDelivery *types.InternetMeasurementsLogDelivery

	// The maximum number of city-networks to monitor for your resources. A
	// city-network is the location (city) where clients access your application
	// resources from and the network or ASN, such as an internet service provider
	// (ISP), that clients access the resources through. This limit helps control
	// billing costs. To learn more, see Choosing a city-network maximum value  (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	MaxCityNetworksToMonitor int32

	// The health of the data processing for the monitor.
	ProcessingStatus types.MonitorProcessingStatusCode

	// Additional information about the health of the data processing for the monitor.
	ProcessingStatusInfo *string

	// The tags that have been added to monitor.
	Tags map[string]string

	// The percentage of the internet-facing traffic for your application that you
	// want to monitor with this monitor.
	TrafficPercentageToMonitor int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMonitor{}, middleware.After)
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
	if err = addOpGetMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMonitor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "internetmonitor",
		OperationName: "GetMonitor",
	}
}
