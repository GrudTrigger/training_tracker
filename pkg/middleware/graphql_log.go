package middleware

import (
	"context"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type LoggingExtension struct{}

func (LoggingExtension) ExtensionName() string {
	return "LoggingExtension"
}

func (LoggingExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (LoggingExtension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()
	resp := next(ctx) //Передаем контекст запроса дальше на выполнение и распаршивание в graphql и ждем выполнение, записываем в переменную "resp" response чтобы потом передать ответ дальше по цепочки клиенту
	elapsed := time.Since(start)

	if op := graphql.GetOperationContext(ctx); op != nil {
		log.Printf("[GraphQL] Operation: %s %s  time=%s",op.Operation.Operation, op.OperationName, elapsed)
	} else {
		log.Printf("[GraphQL] Anonymous operation time=%s", elapsed)
	}

	return resp
}