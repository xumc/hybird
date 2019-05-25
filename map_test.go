package hybrid

import (
	"reflect"
	"testing"
)

type video struct {
	ID        int64
	Name      string
	ViewCount int
}

func Test_Extract_slice_normal_case(t *testing.T) {
	v1 := video{1, "Titanic", 100}
	v2 := video{2, "Kill Bill", 1000}

	videos := []video{v1, v2}

	ids := ExtractInt64(videos, "ID")
	if !reflect.DeepEqual(ids, []int64{1, 2}) {
		t.Fatalf("expected [1,2], but got %v", ids)
	}

	viewCounts := ExtractInt(videos, "ViewCount")
	if !reflect.DeepEqual(viewCounts, []int{100, 1000}) {
		t.Fatalf("expected [1,2], but got %v", viewCounts)
	}

	names := ExtractString(videos, "Name")
	if !reflect.DeepEqual(names, []string{"Titanic", "Kill Bill"}) {
		t.Fatalf("expected [1,2], but got %v", names)
	}
}

func Test_Extract_array_normal_case(t *testing.T) {
	v1 := video{1, "Titanic", 100}
	v2 := video{2, "Kill Bill", 1000}

	videos := [2]video{v1, v2}

	ids := ExtractInt64(videos, "ID")
	if !reflect.DeepEqual(ids, []int64{1, 2}) {
		t.Fatalf("expected [1,2], but got %v", ids)
	}

	viewCounts := ExtractInt(videos, "ViewCount")
	if !reflect.DeepEqual(viewCounts, []int{100, 1000}) {
		t.Fatalf("expected [1,2], but got %v", viewCounts)
	}

	names := ExtractString(videos, "Name")
	if !reflect.DeepEqual(names, []string{"Titanic", "Kill Bill"}) {
		t.Fatalf("expected [1,2], but got %v", names)
	}
}

func Test_Extract_empty_slice_normal_case(t *testing.T) {
	videos := []video{}

	ids := ExtractInt64(videos, "ID")
	if !reflect.DeepEqual(ids, []int64{}) {
		t.Fatalf("expected [1,2], but got %v", ids)
	}

	viewCounts := ExtractInt(videos, "ViewCount")
	if !reflect.DeepEqual(viewCounts, []int{}) {
		t.Fatalf("expected [1,2], but got %v", viewCounts)
	}

	names := ExtractString(videos, "Name")
	if !reflect.DeepEqual(names, []string{}) {
		t.Fatalf("expected [1,2], but got %v", names)
	}
}

func Test_Extract_pointer_element_slice_normal_case(t *testing.T) {
	v1 := &video{1, "Titanic", 100}
	v2 := &video{2, "Kill Bill", 1000}

	videos := []*video{v1, v2}

	ids := ExtractInt64(videos, "ID")
	if !reflect.DeepEqual(ids, []int64{1, 2}) {
		t.Fatalf("expected [1,2], but got %v", ids)
	}

	viewCounts := ExtractInt(videos, "ViewCount")
	if !reflect.DeepEqual(viewCounts, []int{100, 1000}) {
		t.Fatalf("expected [1,2], but got %v", viewCounts)
	}

	names := ExtractString(videos, "Name")
	if !reflect.DeepEqual(names, []string{"Titanic", "Kill Bill"}) {
		t.Fatalf("expected [1,2], but got %v", names)
	}
}

func Test_ExtractInt64_invalid_case(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != "only array and slice are supported for mapping" {
				t.Fatalf("got err: %v", err)
			}
		}
	}()

	v1 := video{1, "Titanic", 100}
	ExtractInt64(v1, "ID")
}

func Test_ExtractInt64_invalid_nil_case(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != "the first arg shouldn't be nil" {
				t.Fatalf("got err: %v", err)
			}
		}
	}()

	ExtractInt64(nil, "ID")
}

func Test_ExtractInt64_invalid_array_pointer_case(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != "only array and slice are supported for mapping" {
				t.Fatalf("got err: %v", err)
			}
		}
	}()

	v1 := video{1, "Titanic", 100}
	v2 := video{2, "Kill Bill", 1000}

	videos := &[2]video{v1, v2}

	ExtractInt64(videos, "ID")
}

func Test_ExtractInt64_invalid_key_case(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != "key doesn't exist in the struct" {
				t.Fatalf("got err: %v", err)
			}
		}
	}()

	v1 := video{1, "Titanic", 100}
	v2 := video{2, "Kill Bill", 1000}

	videos := [2]video{v1, v2}

	ExtractInt64(videos, "invalid_key")
}
