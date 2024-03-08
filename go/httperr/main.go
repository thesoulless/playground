package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func main() {
	mux := http.NewServeMux()
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		// randomize error
		bn, err := rand.Int(rand.Reader, big.NewInt(20))
		if err != nil {
			withError(r, err)
			return
		}

		n := bn.Int64()

		if n != 0 {
			if n > 5 && n < 10 {
				withError(r, &ValidationError{
					Fields: map[string]string{
						"coin": "heads",
					},
				})
				return
			}

			if n >= 10 {
				withError(r, fmt.Errorf("internal server error"))
				return
			}

			withError(r, &AuthError{"unauthorized"})
			return
		}

		hr := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, world!",
		}

		Respond(w, hr, http.StatusOK)
	}

	lm := LoggerMiddleware(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	em := ErrorsMiddleware()
	mux.HandleFunc("/hello", em(lm(helloHandler)))

	http.ListenAndServe(":8080", mux)
}

func withError(r *http.Request, err error) {
	cancelCause := r.Context().Value(ctxCancelCause)
	if cancelCause != nil {
		cancelCause.(context.CancelCauseFunc)(err)
	}
}

func LoggerMiddleware(log *slog.Logger) middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("started handling request")
			next.ServeHTTP(w, r)
		})
	}
}

type ctxKey int

const (
	ctxKeyError ctxKey = iota
	ctxCancelCause
)

type AuthError struct {
	Message string
}

func (e AuthError) Error() string {
	return e.Message
}

func IsAuthError(err error) bool {
	var ae *AuthError
	return errors.As(err, &ae)
}

type ValidationError struct {
	Fields map[string]string
}

func (e ValidationError) Error() string {
	return "validation error"
}

func IsValidationError(err error) bool {
	var ve *ValidationError
	return errors.As(err, &ve)
}

type ErrorDocument struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

func ErrorsMiddleware() middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "something went wrong", http.StatusInternalServerError)
				}
			}()

			ctx, cancel := context.WithCancelCause(context.Background())
			defer func() {
				cancel(nil)
			}()

			r = r.WithContext(context.WithValue(r.Context(), ctxCancelCause, cancel))

			next.ServeHTTP(w, r)

			ctxErr := ctx.Err()
			if err := context.Cause(ctx); ctxErr != nil && err != nil {
				var er ErrorDocument
				var status int

				switch {
				case IsAuthError(err):
					er = ErrorDocument{
						Error: http.StatusText(http.StatusUnauthorized),
					}
					status = http.StatusUnauthorized
				case IsValidationError(err):
					er = ErrorDocument{
						Error:  "validation error",
						Fields: err.(*ValidationError).Fields,
					}
					status = http.StatusBadRequest
				default:
					er = ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				Respond(w, er, status)
			}
		})
	}
}

func Respond(w http.ResponseWriter, data any, statusCode int) {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling data: %#v\n", err)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		fmt.Printf("error writing data: %#v\n", err)
		panic(err)
	}
}
