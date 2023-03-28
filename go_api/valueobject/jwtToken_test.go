package valueobject_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/valueobject"
)

func TestJwtToken_Decode(t *testing.T) {
	userId := uuid.New()
	tokenString, _ := valueobject.NewJwtToken(userId)
	token, _ := tokenString.Encode()
	claims, _ := valueobject.Decode(token)
	if claims.UserId != userId {
		t.Errorf("failed: expected %d, got %d", userId, claims.UserId)
	}
}
