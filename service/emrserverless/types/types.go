// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an application. Amazon EMR Serverless uses applications to
// run jobs.
type Application struct {

	// The ID of the application.
	//
	// This member is required.
	ApplicationId *string

	// The ARN of the application.
	//
	// This member is required.
	Arn *string

	// The date and time when the application run was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon EMR release associated with the application.
	//
	// This member is required.
	ReleaseLabel *string

	// The state of the application.
	//
	// This member is required.
	State ApplicationState

	// The type of application, such as Spark or Hive.
	//
	// This member is required.
	Type *string

	// The date and time when the application run was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The CPU architecture of an application.
	Architecture Architecture

	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration *AutoStartConfig

	// The configuration for an application to automatically stop after a certain
	// amount of time being idle.
	AutoStopConfiguration *AutoStopConfig

	// The image configuration applied to all worker types.
	ImageConfiguration *ImageConfiguration

	// The initial capacity of the application.
	InitialCapacity map[string]InitialCapacityConfig

	// The maximum capacity of the application. This is cumulative across all workers
	// at any given point in time during the lifespan of the application is created. No
	// new resources will be created once any one of the defined limits is hit.
	MaximumCapacity *MaximumAllowedResources

	// The name of the application.
	Name *string

	// The network configuration for customer VPC connectivity for the application.
	NetworkConfiguration *NetworkConfiguration

	// The state details of the application.
	StateDetails *string

	// The tags assigned to the application.
	Tags map[string]string

	// The specification applied to each worker type.
	WorkerTypeSpecifications map[string]WorkerTypeSpecification

	noSmithyDocumentSerde
}

// The summary of attributes associated with an application.
type ApplicationSummary struct {

	// The ARN of the application.
	//
	// This member is required.
	Arn *string

	// The date and time when the application was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ID of the application.
	//
	// This member is required.
	Id *string

	// The Amazon EMR release associated with the application.
	//
	// This member is required.
	ReleaseLabel *string

	// The state of the application.
	//
	// This member is required.
	State ApplicationState

	// The type of application, such as Spark or Hive.
	//
	// This member is required.
	Type *string

	// The date and time when the application was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The CPU architecture of an application.
	Architecture Architecture

	// The name of the application.
	Name *string

	// The state details of the application.
	StateDetails *string

	noSmithyDocumentSerde
}

// The configuration for an application to automatically start on job submission.
type AutoStartConfig struct {

	// Enables the application to automatically start on job submission. Defaults to
	// true.
	Enabled *bool

	noSmithyDocumentSerde
}

// The configuration for an application to automatically stop after a certain
// amount of time being idle.
type AutoStopConfig struct {

	// Enables the application to automatically stop after a certain amount of time
	// being idle. Defaults to true.
	Enabled *bool

	// The amount of idle time in minutes after which your application will
	// automatically stop. Defaults to 15 minutes.
	IdleTimeoutMinutes *int32

	noSmithyDocumentSerde
}

// A configuration specification to be used when provisioning an application. A
// configuration consists of a classification, properties, and optional nested
// configurations. A classification refers to an application-specific configuration
// file. Properties are the settings you want to change in that file.
type Configuration struct {

	// The classification within a configuration.
	//
	// This member is required.
	Classification *string

	// A list of additional configurations to apply within a configuration object.
	Configurations []Configuration

	// A set of properties specified within a configuration classification.
	Properties map[string]string

	noSmithyDocumentSerde
}

// A configuration specification to be used to override existing configurations.
type ConfigurationOverrides struct {

	// The override configurations for the application.
	ApplicationConfiguration []Configuration

	// The override configurations for monitoring.
	MonitoringConfiguration *MonitoringConfiguration

	noSmithyDocumentSerde
}

// The configurations for the Hive job driver.
type Hive struct {

	// The query for the Hive job run.
	//
	// This member is required.
	Query *string

	// The query file for the Hive job run.
	InitQueryFile *string

	// The parameters for the Hive job run.
	Parameters *string

	noSmithyDocumentSerde
}

// The applied image configuration.
type ImageConfiguration struct {

	// The image URI.
	//
	// This member is required.
	ImageUri *string

	// The SHA256 digest of the image URI. This indicates which specific image the
	// application is configured for. The image digest doesn't exist until an
	// application has started.
	ResolvedImageDigest *string

	noSmithyDocumentSerde
}

// The image configuration.
type ImageConfigurationInput struct {

	// The URI of an image in the Amazon ECR registry. This field is required when you
	// create a new application. If you leave this field blank in an update, Amazon EMR
	// will remove the image configuration.
	ImageUri *string

	noSmithyDocumentSerde
}

// The initial capacity configuration per worker.
type InitialCapacityConfig struct {

	// The number of workers in the initial capacity configuration.
	//
	// This member is required.
	WorkerCount int64

	// The resource configuration of the initial capacity configuration.
	WorkerConfiguration *WorkerResourceConfig

	noSmithyDocumentSerde
}

// The driver that the job runs on.
//
// The following types satisfy this interface:
//
//	JobDriverMemberHive
//	JobDriverMemberSparkSubmit
type JobDriver interface {
	isJobDriver()
}

// The job driver parameters specified for Hive.
type JobDriverMemberHive struct {
	Value Hive

	noSmithyDocumentSerde
}

func (*JobDriverMemberHive) isJobDriver() {}

// The job driver parameters specified for Spark.
type JobDriverMemberSparkSubmit struct {
	Value SparkSubmit

	noSmithyDocumentSerde
}

func (*JobDriverMemberSparkSubmit) isJobDriver() {}

// Information about a job run. A job run is a unit of work, such as a Spark JAR,
// Hive query, or SparkSQL query, that you submit to an Amazon EMR Serverless
// application.
type JobRun struct {

	// The ID of the application the job is running on.
	//
	// This member is required.
	ApplicationId *string

	// The execution role ARN of the job run.
	//
	// This member is required.
	Arn *string

	// The date and time when the job run was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user who created the job run.
	//
	// This member is required.
	CreatedBy *string

	// The execution role ARN of the job run.
	//
	// This member is required.
	ExecutionRole *string

	// The job driver for the job run.
	//
	// This member is required.
	JobDriver JobDriver

	// The ID of the job run.
	//
	// This member is required.
	JobRunId *string

	// The Amazon EMR release associated with the application your job is running on.
	//
	// This member is required.
	ReleaseLabel *string

	// The state of the job run.
	//
	// This member is required.
	State JobRunState

	// The state details of the job run.
	//
	// This member is required.
	StateDetails *string

	// The date and time when the job run was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The aggregate vCPU, memory, and storage that Amazon Web Services has billed for
	// the job run. The billed resources include a 1-minute minimum usage for workers,
	// plus additional storage over 20 GB per worker. Note that billed resources do not
	// include usage for idle pre-initialized workers.
	BilledResourceUtilization *ResourceUtilization

	// The configuration settings that are used to override default configuration.
	ConfigurationOverrides *ConfigurationOverrides

	// Returns the job run timeout value from the StartJobRun call. If no timeout was
	// specified, then it returns the default timeout of 720 minutes.
	ExecutionTimeoutMinutes *int64

	// The optional job run name. This doesn't have to be unique.
	Name *string

	// The network configuration for customer VPC connectivity.
	NetworkConfiguration *NetworkConfiguration

	// The tags assigned to the job run.
	Tags map[string]string

	// The job run total execution duration in seconds. This field is only available
	// for job runs in a COMPLETED , FAILED , or CANCELLED state.
	TotalExecutionDurationSeconds *int32

	// The aggregate vCPU, memory, and storage resources used from the time the job
	// starts to execute, until the time the job terminates, rounded up to the nearest
	// second.
	TotalResourceUtilization *TotalResourceUtilization

	noSmithyDocumentSerde
}

// The summary of attributes associated with a job run.
type JobRunSummary struct {

	// The ID of the application the job is running on.
	//
	// This member is required.
	ApplicationId *string

	// The ARN of the job run.
	//
	// This member is required.
	Arn *string

	// The date and time when the job run was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user who created the job run.
	//
	// This member is required.
	CreatedBy *string

	// The execution role ARN of the job run.
	//
	// This member is required.
	ExecutionRole *string

	// The ID of the job run.
	//
	// This member is required.
	Id *string

	// The Amazon EMR release associated with the application your job is running on.
	//
	// This member is required.
	ReleaseLabel *string

	// The state of the job run.
	//
	// This member is required.
	State JobRunState

	// The state details of the job run.
	//
	// This member is required.
	StateDetails *string

	// The date and time when the job run was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The optional job run name. This doesn't have to be unique.
	Name *string

	// The type of job run, such as Spark or Hive.
	Type *string

	noSmithyDocumentSerde
}

// The managed log persistence configuration for a job run.
type ManagedPersistenceMonitoringConfiguration struct {

	// Enables managed logging and defaults to true. If set to false, managed logging
	// will be turned off.
	Enabled *bool

	// The KMS key ARN to encrypt the logs stored in managed log persistence.
	EncryptionKeyArn *string

	noSmithyDocumentSerde
}

// The maximum allowed cumulative resources for an application. No new resources
// will be created once the limit is hit.
type MaximumAllowedResources struct {

	// The maximum allowed CPU for an application.
	//
	// This member is required.
	Cpu *string

	// The maximum allowed resources for an application.
	//
	// This member is required.
	Memory *string

	// The maximum allowed disk for an application.
	Disk *string

	noSmithyDocumentSerde
}

// The configuration setting for monitoring.
type MonitoringConfiguration struct {

	// The managed log persistence configuration for a job run.
	ManagedPersistenceMonitoringConfiguration *ManagedPersistenceMonitoringConfiguration

	// The Amazon S3 configuration for monitoring log publishing.
	S3MonitoringConfiguration *S3MonitoringConfiguration

	noSmithyDocumentSerde
}

// The network configuration for customer VPC connectivity.
type NetworkConfiguration struct {

	// The array of security group Ids for customer VPC connectivity.
	SecurityGroupIds []string

	// The array of subnet Ids for customer VPC connectivity.
	SubnetIds []string

	noSmithyDocumentSerde
}

// The resource utilization for memory, storage, and vCPU for jobs.
type ResourceUtilization struct {

	// The aggregated memory used per hour from the time the job starts executing
	// until the job is terminated.
	MemoryGBHour *float64

	// The aggregated storage used per hour from the time the job starts executing
	// until the job is terminated.
	StorageGBHour *float64

	// The aggregated vCPU used per hour from the time the job starts executing until
	// the job is terminated.
	VCPUHour *float64

	noSmithyDocumentSerde
}

// The Amazon S3 configuration for monitoring log publishing. You can configure
// your jobs to send log information to Amazon S3.
type S3MonitoringConfiguration struct {

	// The KMS key ARN to encrypt the logs published to the given Amazon S3
	// destination.
	EncryptionKeyArn *string

	// The Amazon S3 destination URI for log publishing.
	LogUri *string

	noSmithyDocumentSerde
}

// The configurations for the Spark submit job driver.
type SparkSubmit struct {

	// The entry point for the Spark submit job run.
	//
	// This member is required.
	EntryPoint *string

	// The arguments for the Spark submit job run.
	EntryPointArguments []string

	// The parameters for the Spark submit job run.
	SparkSubmitParameters *string

	noSmithyDocumentSerde
}

// The aggregate vCPU, memory, and storage resources used from the time job start
// executing till the time job is terminated, rounded up to the nearest second.
type TotalResourceUtilization struct {

	// The aggregated memory used per hour from the time job start executing till the
	// time job is terminated.
	MemoryGBHour *float64

	// The aggregated storage used per hour from the time job start executing till the
	// time job is terminated.
	StorageGBHour *float64

	// The aggregated vCPU used per hour from the time job start executing till the
	// time job is terminated.
	VCPUHour *float64

	noSmithyDocumentSerde
}

// The cumulative configuration requirements for every worker instance of the
// worker type.
type WorkerResourceConfig struct {

	// The CPU requirements for every worker instance of the worker type.
	//
	// This member is required.
	Cpu *string

	// The memory requirements for every worker instance of the worker type.
	//
	// This member is required.
	Memory *string

	// The disk requirements for every worker instance of the worker type.
	Disk *string

	noSmithyDocumentSerde
}

// The specifications for a worker type.
type WorkerTypeSpecification struct {

	// The image configuration for a worker type.
	ImageConfiguration *ImageConfiguration

	noSmithyDocumentSerde
}

// The specifications for a worker type.
type WorkerTypeSpecificationInput struct {

	// The image configuration for a worker type.
	ImageConfiguration *ImageConfigurationInput

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isJobDriver() {}
