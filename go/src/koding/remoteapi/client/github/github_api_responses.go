package github

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// GithubAPIReader is a Reader for the GithubAPI structure.
type GithubAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GithubAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGithubAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGithubAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGithubAPIOK creates a GithubAPIOK with default headers values
func NewGithubAPIOK() *GithubAPIOK {
	return &GithubAPIOK{}
}

/*GithubAPIOK handles this case with default header values.

Request processed successfully
*/
type GithubAPIOK struct {
	Payload *models.DefaultResponse
}

func (o *GithubAPIOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Github.api][%d] githubApiOK  %+v", 200, o.Payload)
}

func (o *GithubAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGithubAPIUnauthorized creates a GithubAPIUnauthorized with default headers values
func NewGithubAPIUnauthorized() *GithubAPIUnauthorized {
	return &GithubAPIUnauthorized{}
}

/*GithubAPIUnauthorized handles this case with default header values.

Unauthorized request
*/
type GithubAPIUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *GithubAPIUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Github.api][%d] githubApiUnauthorized  %+v", 401, o.Payload)
}

func (o *GithubAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
