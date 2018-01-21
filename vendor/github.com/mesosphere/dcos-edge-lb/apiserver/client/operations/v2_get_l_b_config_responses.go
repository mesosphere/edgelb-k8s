// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// V2GetLBConfigReader is a Reader for the V2GetLBConfig structure.
type V2GetLBConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetLBConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewV2GetLBConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewV2GetLBConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2GetLBConfigOK creates a V2GetLBConfigOK with default headers values
func NewV2GetLBConfigOK() *V2GetLBConfigOK {
	return &V2GetLBConfigOK{}
}

/*V2GetLBConfigOK handles this case with default header values.

Rendered lb config for pool.
*/
type V2GetLBConfigOK struct {
	Payload string
}

func (o *V2GetLBConfigOK) Error() string {
	return fmt.Sprintf("[GET /v2/pools/{name}/lbconfig][%d] v2GetLBConfigOK  %+v", 200, o.Payload)
}

func (o *V2GetLBConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetLBConfigDefault creates a V2GetLBConfigDefault with default headers values
func NewV2GetLBConfigDefault(code int) *V2GetLBConfigDefault {
	return &V2GetLBConfigDefault{
		_statusCode: code,
	}
}

/*V2GetLBConfigDefault handles this case with default header values.

Unexpected error.
*/
type V2GetLBConfigDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the v2 get l b config default response
func (o *V2GetLBConfigDefault) Code() int {
	return o._statusCode
}

func (o *V2GetLBConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v2/pools/{name}/lbconfig][%d] V2GetLBConfig default  %+v", o._statusCode, o.Payload)
}

func (o *V2GetLBConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
