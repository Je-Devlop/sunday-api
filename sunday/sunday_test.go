package sunday

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestCreateScoops(t *testing.T) {

	t.Run("it should error when request is empty", func(t *testing.T) {
		h := NewSundayHandler(&gorm.DB{})

		w := httptest.NewRecorder()
		payload := bytes.NewBufferString(`{name:"", image_path:""}`)
		req, _ := http.NewRequest("POST", "http://0.0.0.0:8080/create-scoops", payload)

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		h.CreateScoops(c)

		want := 400

		if want != w.Result().StatusCode {
			t.Errorf("want %d but get %s\n", want, w.Result().Status)
		}

	})

}
