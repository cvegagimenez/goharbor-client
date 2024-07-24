// Code generated by go-swagger; DO NOT EDIT.

package securityhub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// GetSecuritySummaryReader is a Reader for the GetSecuritySummary structure.
type GetSecuritySummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecuritySummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecuritySummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSecuritySummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecuritySummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSecuritySummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSecuritySummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecuritySummaryOK creates a GetSecuritySummaryOK with default headers values
func NewGetSecuritySummaryOK() *GetSecuritySummaryOK {
	return &GetSecuritySummaryOK{}
}

/*GetSecuritySummaryOK handles this case with default header values.

Success
*/
type GetSecuritySummaryOK struct {
	Payload *model.SecuritySummary
}

func (o *GetSecuritySummaryOK) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryOK  %+v", 200, o.Payload)
}

func (o *GetSecuritySummaryOK) GetPayload() *model.SecuritySummary {
	return o.Payload
}

func (o *GetSecuritySummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SecuritySummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecuritySummaryUnauthorized creates a GetSecuritySummaryUnauthorized with default headers values
func NewGetSecuritySummaryUnauthorized() *GetSecuritySummaryUnauthorized {
	return &GetSecuritySummaryUnauthorized{}
}

/*GetSecuritySummaryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSecuritySummaryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSecuritySummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSecuritySummaryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecuritySummaryForbidden creates a GetSecuritySummaryForbidden with default headers values
func NewGetSecuritySummaryForbidden() *GetSecuritySummaryForbidden {
	return &GetSecuritySummaryForbidden{}
}

/*GetSecuritySummaryForbidden handles this case with default header values.

Forbidden
*/
type GetSecuritySummaryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSecuritySummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryForbidden  %+v", 403, o.Payload)
}

func (o *GetSecuritySummaryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecuritySummaryNotFound creates a GetSecuritySummaryNotFound with default headers values
func NewGetSecuritySummaryNotFound() *GetSecuritySummaryNotFound {
	return &GetSecuritySummaryNotFound{}
}

/*GetSecuritySummaryNotFound handles this case with default header values.

Not found
*/
type GetSecuritySummaryNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSecuritySummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetSecuritySummaryNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecuritySummaryInternalServerError creates a GetSecuritySummaryInternalServerError with default headers values
func NewGetSecuritySummaryInternalServerError() *GetSecuritySummaryInternalServerError {
	return &GetSecuritySummaryInternalServerError{}
}

/*GetSecuritySummaryInternalServerError handles this case with default header values.

Internal server error
*/
type GetSecuritySummaryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSecuritySummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecuritySummaryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}