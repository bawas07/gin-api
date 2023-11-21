package loggers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RequestLogger(l *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		defer func() {
			l.
				Info().
				Str("method", c.Request.Method).
				Str("url", c.Request.URL.RequestURI()).
				Str("user_agent", c.Request.UserAgent()).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()

		c.Next()
	}
}

// func RequestLogger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		l := Get()

// 		defer func() {
// 			l.
// 				Info().
// 				Str("method", r.Method).
// 				Str("url", r.URL.RequestURI()).
// 				Str("user_agent", r.UserAgent()).
// 				Dur("elapsed_ms", time.Since(start)).
// 				Msg("incoming request")
// 		}()

// 		next.ServeHTTP(w, r)

// 	})
// }
