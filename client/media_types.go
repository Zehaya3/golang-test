// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "message": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/golang-test/src/message/design
// --out=$(GOPATH)/src/github.com/golang-test
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A message of room (default view)
//
// Identifier: application/vnd.goa.example.message+json; view=default
type GoaExampleMessage struct {
	// API href for making requests on the message
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique Message ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Sentence of Message
	Message string `form:"message" json:"message" yaml:"message" xml:"message"`
}

// Validate validates the GoaExampleMessage media type instance.
func (mt *GoaExampleMessage) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// DecodeGoaExampleMessage decodes the GoaExampleMessage instance encoded in resp body.
func (c *Client) DecodeGoaExampleMessage(resp *http.Response) (*GoaExampleMessage, error) {
	var decoded GoaExampleMessage
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
