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
	resp := next(ctx)
	elapsed := time.Since(start)

	if op := graphql.GetOperationContext(ctx); op != nil {
		log.Printf("[GraphQL] Operation: %s %s  time=%s",op.Operation.Operation, op.OperationName, elapsed)
	} else {
		log.Printf("[GraphQL] Anonymous operation took %s", elapsed)
	}

	return resp
}