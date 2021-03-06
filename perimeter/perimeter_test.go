package shape

import (
  "testing"
  "reflect"
)

func TestGetArea(t *testing.T){
	var i1 struct {
		s1 Shape;
	}

	var c1 Circle
	var r1 Rectangle

	v := reflect.ValueOf(i1)
	testGetArea(t, reflect.TypeOf(c1), v.Field(0).Type())
	testGetArea(t, reflect.TypeOf(r1), v.Field(0).Type())
}

func TestGetPerimeter(t *testing.T){
	var i2 struct {
		s2 Shape;
	}

	var c2 Circle
	var r2 Rectangle
	v := reflect.ValueOf(i2)
	testGetArea(t, reflect.TypeOf(c2), v.Field(0).Type())
	testGetArea(t, reflect.TypeOf(r2), v.Field(0).Type())
}

func testGetArea(t *testing.T, t1, t2 reflect.Type){
	if !t1.Implements(t2){
		t.Errorf("%v does not implement %v", t1, t2)
	}
}
