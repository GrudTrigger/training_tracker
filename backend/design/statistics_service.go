package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

var _ = Service("statistics", func() {
	Description("Получение статистики по тренировкам")

	Method("get-trainings-statistics", func() {
		Description("Получение статисики")
		Meta("openapi:summary", "Получение статисики")

		Result(TrainingsStatistics)

		HTTP(func() {
			GET("statistics")
			Response(StatusOK)
		})
	})
})
