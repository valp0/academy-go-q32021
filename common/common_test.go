package common

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonRes(t *testing.T) {
	got1 := JsonResponse("Testing")
	want1, _ := json.Marshal(body{message{"Testing"}})

	got2 := JsonResponse([]int{1, 2, 3})
	want2, _ := json.Marshal(body{items{[]int{1, 2, 3}}})

	if !reflect.DeepEqual(want1, got1) {
		t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", got1, want1)
	}

	if !reflect.DeepEqual(want2, got2) {
		t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", got1, want1)
	}
}

func TestPrettyJson(t *testing.T) {
	want1 := `{` + "\n" +
		`  "data": {` + "\n" +
		`    "items": [` + "\n" +
		`      {` + "\n" +
		`        "id": 51,` + "\n" +
		`        "name": "dugtrio"` + "\n" +
		`      },` + "\n" +
		`      {` + "\n" +
		`        "id": 737,` + "\n" +
		`        "name": "charjabug"` + "\n" +
		`      }` + "\n" +
		`    ]` + "\n" +
		`  }` + "\n" +
		`}`

	want2 := `{` + "\n" +
		`  "data": {` + "\n" +
		`    "message": "Testing"` + "\n" +
		`  }` + "\n" +
		`}`

	want3 := "Testing"

	elems := []Element{}
	elems = append(elems, Element{Key: 51, Value: "dugtrio"})
	elems = append(elems, Element{Key: 737, Value: "charjabug"})

	got1 := PrettyJsonRes(elems)
	got2 := PrettyJsonRes("Testing")
	got3 := PrettifyJson([]byte{84, 101, 115, 116, 105, 110, 103})

	if want1 != got1 {
		t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", got1, want1)
	}

	if want2 != got2 {
		t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", got2, want2)
	}

	if want3 != got3 {
		t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", got3, want3)
	}
}

func TestRandom(t *testing.T) {
	var rand int
	max := 898
	for i := 0; i < 100000; i++ {
		rand = RandInt(max)
		if !(rand < 899 && 0 < rand) {
			t.Fatalf("rand function returned a number out of range: %d", rand)
		}
	}

	for i := 0; i < 100000; i++ {
		rand = randInt(max)
		if !(rand < 899 && 0 < rand) {
			t.Fatalf("rand function returned a number out of range: %d", rand)
		}
	}
}
