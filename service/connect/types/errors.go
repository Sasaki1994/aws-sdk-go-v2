// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have sufficient permissions to perform this action.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The flow has not been published.
type ContactFlowNotPublishedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ContactFlowNotPublishedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContactFlowNotPublishedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContactFlowNotPublishedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ContactFlowNotPublishedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ContactFlowNotPublishedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The contact with the specified ID is not active or does not exist. Applies to
// Voice calls only, not to Chat, Task, or Voice Callback.
type ContactNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ContactNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContactNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContactNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ContactNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ContactNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Outbound calls to the destination number are not allowed.
type DestinationNotAllowedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DestinationNotAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DestinationNotAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DestinationNotAllowedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DestinationNotAllowedException"
	}
	return *e.ErrorCodeOverride
}
func (e *DestinationNotAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the specified name already exists.
type DuplicateResourceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicateResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateResourceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateResourceException"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateResourceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An entity with the same name already exists.
type IdempotencyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *IdempotencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotencyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IdempotencyException"
	}
	return *e.ErrorCodeOverride
}
func (e *IdempotencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request processing failed because of an error or failure with the service.
type InternalServiceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The flow is not valid.
type InvalidContactFlowException struct {
	Message *string

	ErrorCodeOverride *string

	Problems []ProblemDetail

	noSmithyDocumentSerde
}

func (e *InvalidContactFlowException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidContactFlowException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidContactFlowException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidContactFlowException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidContactFlowException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The problems with the module. Please fix before trying again.
type InvalidContactFlowModuleException struct {
	Message *string

	ErrorCodeOverride *string

	Problems []ProblemDetail

	noSmithyDocumentSerde
}

func (e *InvalidContactFlowModuleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidContactFlowModuleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidContactFlowModuleException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidContactFlowModuleException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidContactFlowModuleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified parameters are not valid.
type InvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request is not valid.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The allowed limit for the resource has been exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Maximum number (1000) of tags have been returned with current request. Consider
// changing request parameters to get more tags.
type MaximumResultReturnedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MaximumResultReturnedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MaximumResultReturnedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MaximumResultReturnedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MaximumResultReturnedException"
	}
	return *e.ErrorCodeOverride
}
func (e *MaximumResultReturnedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The contact is not permitted.
type OutboundContactNotPermittedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *OutboundContactNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OutboundContactNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OutboundContactNotPermittedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OutboundContactNotPermittedException"
	}
	return *e.ErrorCodeOverride
}
func (e *OutboundContactNotPermittedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The property is not valid.
type PropertyValidationException struct {
	Message *string

	ErrorCodeOverride *string

	PropertyList []PropertyValidationExceptionProperty

	noSmithyDocumentSerde
}

func (e *PropertyValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PropertyValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PropertyValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PropertyValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *PropertyValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource already has that name.
type ResourceConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// That resource is already in use. Please try another.
type ResourceInUseException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType ResourceType
	ResourceId   *string

	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource was not found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource is not ready.
type ResourceNotReadyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotReadyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotReadyException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service quota has been exceeded.
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The throttling limit has been exceeded.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No user with the specified credentials was found in the Amazon Connect instance.
type UserNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UserNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *UserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
