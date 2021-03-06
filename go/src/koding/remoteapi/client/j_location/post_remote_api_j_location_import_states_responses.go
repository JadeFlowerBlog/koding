package j_location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJLocationImportStatesReader is a Reader for the PostRemoteAPIJLocationImportStates structure.
type PostRemoteAPIJLocationImportStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJLocationImportStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJLocationImportStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJLocationImportStatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJLocationImportStatesOK creates a PostRemoteAPIJLocationImportStatesOK with default headers values
func NewPostRemoteAPIJLocationImportStatesOK() *PostRemoteAPIJLocationImportStatesOK {
	return &PostRemoteAPIJLocationImportStatesOK{}
}

/*PostRemoteAPIJLocationImportStatesOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJLocationImportStatesOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJLocationImportStatesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.importStates][%d] postRemoteApiJLocationImportStatesOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJLocationImportStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJLocationImportStatesUnauthorized creates a PostRemoteAPIJLocationImportStatesUnauthorized with default headers values
func NewPostRemoteAPIJLocationImportStatesUnauthorized() *PostRemoteAPIJLocationImportStatesUnauthorized {
	return &PostRemoteAPIJLocationImportStatesUnauthorized{}
}

/*PostRemoteAPIJLocationImportStatesUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJLocationImportStatesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJLocationImportStatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.importStates][%d] postRemoteApiJLocationImportStatesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJLocationImportStatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
