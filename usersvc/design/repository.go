package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("repository", func() {
	Description("The service makes it possible to insert, delete or get users.")

	HTTP(func() {
		Path("/")
	})

	Method("insert", func() {
		Description("Add new user and return its ID.")
		Payload(User)
		Result(String)
		HTTP(func() {
			POST("/insert")
			Response(StatusCreated)
		})
	})

	Method("delete", func() {
		Description("Remove user from storage")
		Payload(func() {
			Attribute("id", String, "ID of user to remove")
			Required("id")
		})
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			DELETE("/delete/{id}")
			Response(StatusNoContent)
		})
	})

	Method("get", func() {
		Result(StoredUser)
		Error("not_found", NotFound, "User not found")
		Payload(func() {
			Attribute("id", String, "Get user by ID")
			Attribute("view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		HTTP(func() {
			GET("/get/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})
