package sunday

import (
	"testing"
)

type TestDB struct{}

func (t *TestDB) CreateICreamScoop(scoop Scoop) error {
	return nil
}

func (t *TestDB) GetAllIceCreamScoops() ([]Scoop, error) {
	return []Scoop{}, nil
}

type TestContext struct {
	httpStatus int
}

func (t *TestContext) Bind(v interface{}) error {
	t.httpStatus = 400
	return nil
}

func (t *TestContext) JSON(status int, v interface{}) {
}

func TestCreateScoops(t *testing.T) {

	t.Run("it should error when request is empty", func(t *testing.T) {
		h := NewSundayHandler(&TestDB{})

		c := &TestContext{}

		h.CreateScoops(c)

		want := 400

		if want != c.httpStatus {
			t.Errorf("want %d but get %d\n", want, c.httpStatus)
		}

	})

}
