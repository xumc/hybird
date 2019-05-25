package hybrid

import (
	"reflect"
	"testing"
)

type videoWithCount struct {
	ID        int64
	ViewCount int
}

func TestMapInt64Int(t *testing.T) {
	s := []videoWithCount{
		{1, 100},
		{2, 1000},
	}

	m := MapInt64Int(s, "ID", "ViewCount")

	expected := map[int64]int{
		1: 100,
		2: 1000,
	}
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("expected [1,2], but got %v", m)
	}
}
