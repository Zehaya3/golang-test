package design // The convention consists of naming the design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("message", func() { // API defines the microservice endpoint and
	Title("The Message API")               // other global properties. There should be one
	Description("This is the Message API") // and exactly one API definition appearing in
	Scheme("http")                         // the design.
	Host("localhost:8080")
})

var _ = Resource("message", func() { // Resources group related API endpoints
	BasePath("/messages")      // together. They map to REST resources for REST
	DefaultMedia(MessageMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get messages by Room ID") // with its path, parameters (both path
		Routing(GET("/:roomId"))               // parameters and querystring values) and payload
		Params(func() {                        // (shape of the request body).
			Param("roomId", Integer, "Room ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// MessageMedia defines the media type used to render messages.
var MessageMedia = MediaType("application/vnd.goa.example.message+json", func() {
	Description("A message of room")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique Message ID")
		Attribute("href", String, "API href for making requests on the message")
		Attribute("message", String, "Sentence of Message")
		Required("id", "href", "message")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("message")
	})
})
