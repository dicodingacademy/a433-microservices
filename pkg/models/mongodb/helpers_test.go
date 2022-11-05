package mongodb

import (
	"reflect"
	"testing"
)

func TestToObjID(t *testing.T) {
	objid := ToObjID("5fff0b9175eec7e8b33a0612")

	if reflect.TypeOf(objid).String() != "primitive.ObjectID" {
		t.Errorf("got %v, want %v", reflect.TypeOf(objid).String(), "primitive.ObjectID")
	}
}
