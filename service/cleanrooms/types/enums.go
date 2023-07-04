// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedExceptionReason string

// Enum values for AccessDeniedExceptionReason
const (
	AccessDeniedExceptionReasonInsufficientPermissions AccessDeniedExceptionReason = "INSUFFICIENT_PERMISSIONS"
)

// Values returns all known values for AccessDeniedExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessDeniedExceptionReason) Values() []AccessDeniedExceptionReason {
	return []AccessDeniedExceptionReason{
		"INSUFFICIENT_PERMISSIONS",
	}
}

type AggregateFunctionName string

// Enum values for AggregateFunctionName
const (
	AggregateFunctionNameSum           AggregateFunctionName = "SUM"
	AggregateFunctionNameSumDistinct   AggregateFunctionName = "SUM_DISTINCT"
	AggregateFunctionNameCount         AggregateFunctionName = "COUNT"
	AggregateFunctionNameCountDistinct AggregateFunctionName = "COUNT_DISTINCT"
	AggregateFunctionNameAvg           AggregateFunctionName = "AVG"
)

// Values returns all known values for AggregateFunctionName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AggregateFunctionName) Values() []AggregateFunctionName {
	return []AggregateFunctionName{
		"SUM",
		"SUM_DISTINCT",
		"COUNT",
		"COUNT_DISTINCT",
		"AVG",
	}
}

type AggregationType string

// Enum values for AggregationType
const (
	AggregationTypeCountDistinct AggregationType = "COUNT_DISTINCT"
)

// Values returns all known values for AggregationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AggregationType) Values() []AggregationType {
	return []AggregationType{
		"COUNT_DISTINCT",
	}
}

type AnalysisMethod string

// Enum values for AnalysisMethod
const (
	AnalysisMethodDirectQuery AnalysisMethod = "DIRECT_QUERY"
)

// Values returns all known values for AnalysisMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisMethod) Values() []AnalysisMethod {
	return []AnalysisMethod{
		"DIRECT_QUERY",
	}
}

type AnalysisRuleType string

// Enum values for AnalysisRuleType
const (
	AnalysisRuleTypeAggregation AnalysisRuleType = "AGGREGATION"
	AnalysisRuleTypeList        AnalysisRuleType = "LIST"
)

// Values returns all known values for AnalysisRuleType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisRuleType) Values() []AnalysisRuleType {
	return []AnalysisRuleType{
		"AGGREGATION",
		"LIST",
	}
}

type CollaborationQueryLogStatus string

// Enum values for CollaborationQueryLogStatus
const (
	CollaborationQueryLogStatusEnabled  CollaborationQueryLogStatus = "ENABLED"
	CollaborationQueryLogStatusDisabled CollaborationQueryLogStatus = "DISABLED"
)

// Values returns all known values for CollaborationQueryLogStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CollaborationQueryLogStatus) Values() []CollaborationQueryLogStatus {
	return []CollaborationQueryLogStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ConfiguredTableAnalysisRuleType string

// Enum values for ConfiguredTableAnalysisRuleType
const (
	ConfiguredTableAnalysisRuleTypeAggregation ConfiguredTableAnalysisRuleType = "AGGREGATION"
	ConfiguredTableAnalysisRuleTypeList        ConfiguredTableAnalysisRuleType = "LIST"
)

// Values returns all known values for ConfiguredTableAnalysisRuleType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfiguredTableAnalysisRuleType) Values() []ConfiguredTableAnalysisRuleType {
	return []ConfiguredTableAnalysisRuleType{
		"AGGREGATION",
		"LIST",
	}
}

type ConflictExceptionReason string

// Enum values for ConflictExceptionReason
const (
	ConflictExceptionReasonAlreadyExists     ConflictExceptionReason = "ALREADY_EXISTS"
	ConflictExceptionReasonSubresourcesExist ConflictExceptionReason = "SUBRESOURCES_EXIST"
	ConflictExceptionReasonInvalidState      ConflictExceptionReason = "INVALID_STATE"
)

// Values returns all known values for ConflictExceptionReason. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConflictExceptionReason) Values() []ConflictExceptionReason {
	return []ConflictExceptionReason{
		"ALREADY_EXISTS",
		"SUBRESOURCES_EXIST",
		"INVALID_STATE",
	}
}

type FilterableMemberStatus string

// Enum values for FilterableMemberStatus
const (
	FilterableMemberStatusInvited FilterableMemberStatus = "INVITED"
	FilterableMemberStatusActive  FilterableMemberStatus = "ACTIVE"
)

// Values returns all known values for FilterableMemberStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterableMemberStatus) Values() []FilterableMemberStatus {
	return []FilterableMemberStatus{
		"INVITED",
		"ACTIVE",
	}
}

type JoinOperator string

// Enum values for JoinOperator
const (
	JoinOperatorOr  JoinOperator = "OR"
	JoinOperatorAnd JoinOperator = "AND"
)

// Values returns all known values for JoinOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JoinOperator) Values() []JoinOperator {
	return []JoinOperator{
		"OR",
		"AND",
	}
}

type JoinRequiredOption string

// Enum values for JoinRequiredOption
const (
	JoinRequiredOptionQueryRunner JoinRequiredOption = "QUERY_RUNNER"
)

// Values returns all known values for JoinRequiredOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JoinRequiredOption) Values() []JoinRequiredOption {
	return []JoinRequiredOption{
		"QUERY_RUNNER",
	}
}

type MemberAbility string

// Enum values for MemberAbility
const (
	MemberAbilityCanQuery          MemberAbility = "CAN_QUERY"
	MemberAbilityCanReceiveResults MemberAbility = "CAN_RECEIVE_RESULTS"
)

// Values returns all known values for MemberAbility. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberAbility) Values() []MemberAbility {
	return []MemberAbility{
		"CAN_QUERY",
		"CAN_RECEIVE_RESULTS",
	}
}

type MembershipQueryLogStatus string

// Enum values for MembershipQueryLogStatus
const (
	MembershipQueryLogStatusEnabled  MembershipQueryLogStatus = "ENABLED"
	MembershipQueryLogStatusDisabled MembershipQueryLogStatus = "DISABLED"
)

// Values returns all known values for MembershipQueryLogStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (MembershipQueryLogStatus) Values() []MembershipQueryLogStatus {
	return []MembershipQueryLogStatus{
		"ENABLED",
		"DISABLED",
	}
}

type MembershipStatus string

// Enum values for MembershipStatus
const (
	MembershipStatusActive               MembershipStatus = "ACTIVE"
	MembershipStatusRemoved              MembershipStatus = "REMOVED"
	MembershipStatusCollaborationDeleted MembershipStatus = "COLLABORATION_DELETED"
)

// Values returns all known values for MembershipStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MembershipStatus) Values() []MembershipStatus {
	return []MembershipStatus{
		"ACTIVE",
		"REMOVED",
		"COLLABORATION_DELETED",
	}
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusInvited MemberStatus = "INVITED"
	MemberStatusActive  MemberStatus = "ACTIVE"
	MemberStatusLeft    MemberStatus = "LEFT"
	MemberStatusRemoved MemberStatus = "REMOVED"
)

// Values returns all known values for MemberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberStatus) Values() []MemberStatus {
	return []MemberStatus{
		"INVITED",
		"ACTIVE",
		"LEFT",
		"REMOVED",
	}
}

type ProtectedQueryStatus string

// Enum values for ProtectedQueryStatus
const (
	ProtectedQueryStatusSubmitted  ProtectedQueryStatus = "SUBMITTED"
	ProtectedQueryStatusStarted    ProtectedQueryStatus = "STARTED"
	ProtectedQueryStatusCancelled  ProtectedQueryStatus = "CANCELLED"
	ProtectedQueryStatusCancelling ProtectedQueryStatus = "CANCELLING"
	ProtectedQueryStatusFailed     ProtectedQueryStatus = "FAILED"
	ProtectedQueryStatusSuccess    ProtectedQueryStatus = "SUCCESS"
	ProtectedQueryStatusTimedOut   ProtectedQueryStatus = "TIMED_OUT"
)

// Values returns all known values for ProtectedQueryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProtectedQueryStatus) Values() []ProtectedQueryStatus {
	return []ProtectedQueryStatus{
		"SUBMITTED",
		"STARTED",
		"CANCELLED",
		"CANCELLING",
		"FAILED",
		"SUCCESS",
		"TIMED_OUT",
	}
}

type ProtectedQueryType string

// Enum values for ProtectedQueryType
const (
	ProtectedQueryTypeSql ProtectedQueryType = "SQL"
)

// Values returns all known values for ProtectedQueryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProtectedQueryType) Values() []ProtectedQueryType {
	return []ProtectedQueryType{
		"SQL",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeConfiguredTable            ResourceType = "CONFIGURED_TABLE"
	ResourceTypeCollaboration              ResourceType = "COLLABORATION"
	ResourceTypeMembership                 ResourceType = "MEMBERSHIP"
	ResourceTypeConfiguredTableAssociation ResourceType = "CONFIGURED_TABLE_ASSOCIATION"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"CONFIGURED_TABLE",
		"COLLABORATION",
		"MEMBERSHIP",
		"CONFIGURED_TABLE_ASSOCIATION",
	}
}

type ResultFormat string

// Enum values for ResultFormat
const (
	ResultFormatCsv     ResultFormat = "CSV"
	ResultFormatParquet ResultFormat = "PARQUET"
)

// Values returns all known values for ResultFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResultFormat) Values() []ResultFormat {
	return []ResultFormat{
		"CSV",
		"PARQUET",
	}
}

type ScalarFunctions string

// Enum values for ScalarFunctions
const (
	ScalarFunctionsTrunc    ScalarFunctions = "TRUNC"
	ScalarFunctionsAbs      ScalarFunctions = "ABS"
	ScalarFunctionsCeiling  ScalarFunctions = "CEILING"
	ScalarFunctionsFloor    ScalarFunctions = "FLOOR"
	ScalarFunctionsLn       ScalarFunctions = "LN"
	ScalarFunctionsLog      ScalarFunctions = "LOG"
	ScalarFunctionsRound    ScalarFunctions = "ROUND"
	ScalarFunctionsSqrt     ScalarFunctions = "SQRT"
	ScalarFunctionsCast     ScalarFunctions = "CAST"
	ScalarFunctionsLower    ScalarFunctions = "LOWER"
	ScalarFunctionsRtrim    ScalarFunctions = "RTRIM"
	ScalarFunctionsUpper    ScalarFunctions = "UPPER"
	ScalarFunctionsCoalesce ScalarFunctions = "COALESCE"
)

// Values returns all known values for ScalarFunctions. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScalarFunctions) Values() []ScalarFunctions {
	return []ScalarFunctions{
		"TRUNC",
		"ABS",
		"CEILING",
		"FLOOR",
		"LN",
		"LOG",
		"ROUND",
		"SQRT",
		"CAST",
		"LOWER",
		"RTRIM",
		"UPPER",
		"COALESCE",
	}
}

type SchemaType string

// Enum values for SchemaType
const (
	SchemaTypeTable SchemaType = "TABLE"
)

// Values returns all known values for SchemaType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SchemaType) Values() []SchemaType {
	return []SchemaType{
		"TABLE",
	}
}

type TargetProtectedQueryStatus string

// Enum values for TargetProtectedQueryStatus
const (
	TargetProtectedQueryStatusCancelled TargetProtectedQueryStatus = "CANCELLED"
)

// Values returns all known values for TargetProtectedQueryStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetProtectedQueryStatus) Values() []TargetProtectedQueryStatus {
	return []TargetProtectedQueryStatus{
		"CANCELLED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonInvalidConfiguration  ValidationExceptionReason = "INVALID_CONFIGURATION"
	ValidationExceptionReasonInvalidQuery          ValidationExceptionReason = "INVALID_QUERY"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"FIELD_VALIDATION_FAILED",
		"INVALID_CONFIGURATION",
		"INVALID_QUERY",
	}
}
