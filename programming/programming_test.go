package programming

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	programminglib "github.com/ricardo-mng/learning-go-lib/programming"
	"github.com/stretchr/testify/assert"
)

func setupGin(mockInterface *programminglib.MockInterface) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	SetRouterGroup(mockInterface, v1)

	return r
}

func TestPostUuid(t *testing.T) {
	// arrange
	mockInterface := programminglib.MockInterface{}
	mockCall := mockInterface.On("NewUuid", false)
	mockCall.Return("1ce44be5-fe68-46f7-a153-51c1c91a4ae4")

	r := setupGin(&mockInterface)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/programming/uuid", nil)

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, w.Code, http.StatusOK)

	output := postUuidOutput{}
	err := json.Unmarshal(w.Body.Bytes(), &output)

	assert.Nil(t, err)
	assert.Len(t, output.UUID, 36)
	assert.Contains(t, output.UUID, "-")

	mockInterface.AssertExpectations(t)
}

func TestPostUuidWithNoHyphen(t *testing.T) {
	// arrange
	mockInterface := programminglib.MockInterface{}
	mockCall := mockInterface.On("NewUuid", true)
	mockCall.Return("1ce44be5fe6846f7a15351c1c91a4ae4")

	r := setupGin(&mockInterface)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/programming/uuid?no-hyphens=true", nil)

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, w.Code, http.StatusOK)

	output := postUuidOutput{}
	err := json.Unmarshal(w.Body.Bytes(), &output)

	assert.Nil(t, err)
	assert.Len(t, output.UUID, 32)
	assert.NotContains(t, output.UUID, "-")

	mockInterface.AssertExpectations(t)
}
