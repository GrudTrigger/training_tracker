package design

import (
	. "goa.design/goa/v3/dsl" //nolint:staticcheck
)

var _ = Service("statistics", func(){
	Description("Получение статистики по тренировкам")

	Method("get-all-exercise", func(){
		Description("Получение всех подходов")
		Meta("openapi:summary", "Получение всех подходов")
		
	})
})