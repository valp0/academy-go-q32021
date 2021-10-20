package services

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/valp0/academy-go-q32021/repo"
)

func TestAsyncSvc(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		params         map[string][]string
		expectedLength int
		err            error
	}{
		{
			"works fine, displays whole list",
			"../files/pokemock.csv",
			map[string][]string{},
			9,
			nil,
		},
		{
			"wrong id error",
			"../files/wrongId.csv",
			map[string][]string{},
			0,
			errors.New("parsing error in line 2: \"Hola\" is not a valid id, please use integers for id field"),
		},
		{
			"wrong path error",
			"../wrongPath/someFile.csv",
			map[string][]string{},
			0,
			errors.New("open ../wrongPath/someFile.csv: El sistema no puede encontrar la ruta especificada."),
		},
		{
			"unexisting file, creates it",
			"../files/deleteThisAsWell.csv",
			map[string][]string{},
			0,
			nil,
		},
		{
			"displays odd elements",
			"../files/pokemock.csv",
			map[string][]string{"type": {"odd"}},
			3,
			nil,
		},
		{
			"displays even elements",
			"../files/pokemock.csv",
			map[string][]string{"type": {"even"}},
			6,
			nil,
		},
		{
			"displays three elements",
			"../files/pokemock.csv",
			map[string][]string{"items": {"3"}},
			3,
			nil,
		},
		{
			"displays eight elements",
			"../files/pokemock.csv",
			map[string][]string{"items_per_worker": {"1"}},
			8,
			nil,
		},
	}

	r := repo.NewAsyncRepo()
	aSvc := NewAsyncSvc(r)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tc.params
			res, err := aSvc.Select(params, tc.path)
			if err == nil && tc.err != nil || err != nil && tc.err == nil {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if err == nil && err != tc.err {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if len(res) != tc.expectedLength {
				t.Fatalf("response has different length:\ngot: %v\nwant: %v\n", len(res), tc.expectedLength)
			}
		})
	}

	err := os.Remove("../files/deleteThisAsWell.csv")
	if err != nil {
		log.Print("Couldn't remove ../files/deleteThisAsWell.csv")
	}
}
