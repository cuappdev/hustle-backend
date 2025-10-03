package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type ctxKey string

const UIDKey ctxKey = "uid"

func RequireFirebaseUser(ac *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		const pref = "Bearer "
		authz := c.GetHeader("Authorization")
		if !strings.HasPrefix(authz, pref) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		idToken := strings.TrimPrefix(authz, pref)

		tok, err := ac.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// put uid into context for handlers
		ctx := context.WithValue(c.Request.Context(), UIDKey, tok.UID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Helper to read the uid in handlers:
func UIDFrom(c *gin.Context) string {
	v := c.Request.Context().Value(UIDKey)
	if s, ok := v.(string); ok { return s }
	return ""
}
