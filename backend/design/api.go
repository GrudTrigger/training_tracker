package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("training_tracker", func() {
	Title("TRAINING TRACKER API")
	Description("Описание методов взаимодействия с API приложения Training Tracker")
	Version("1.0")

	Server("training tracker", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})
