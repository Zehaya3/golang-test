// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "message": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/golang-test/src/message/design
// --out=$(GOPATH)/src/github.com/golang-test
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// MessageHref returns the resource href.
func MessageHref(roomID interface{}) string {
	paramroomID := strings.TrimLeftFunc(fmt.Sprintf("%v", roomID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/messages/%v", paramroomID)
}
