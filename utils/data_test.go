package utils

import (
	"github.com/andrewGIS/apiserver/db"
	"testing"
)

func TestMaxValueInArray(t *testing.T) {
	cases := []struct {
		data   []db.DataMeasure
		Result db.DataMeasure
		Err    error
	}{
		{
			data: []db.DataMeasure{
				db.DataMeasure{T: 32, Time: "2019-08-24T09:35:06.654Z"},
				db.DataMeasure{T: 33, Time: "2019-08-24T09:35:06.654Z"},
			},
			Result: db.DataMeasure{T: 33, Time: "2019-08-24T09:35:06.654Z"},
			Err:    nil,
		},
	}

	for _, testCase := range cases {
		res := MaxValueInArray(testCase.data)
		if res != testCase.Result {
			t.Errorf("Excepted %v, got is %v", testCase.Result, res)
		}
	}

}

func TestMinValueInArray(t *testing.T) {
	cases := []struct {
		data   []db.DataMeasure
		Result db.DataMeasure
		Err    error
	}{
		{
			data: []db.DataMeasure{
				db.DataMeasure{T: 32, Time: "2019-08-24T09:35:06.654Z"},
				db.DataMeasure{T: 33, Time: "2019-08-24T09:35:06.654Z"},
			},
			Result: db.DataMeasure{T: 32, Time: "2019-08-24T09:35:06.654Z"},
			Err:    nil,
		},
		{
			data: []db.DataMeasure{
				db.DataMeasure{T: 32, Time: "2019-08-24T09:35:06.654Z"},
				db.DataMeasure{T: 33, Time: "2019-08-24T09:35:06.654Z"},
			},
			Result: db.DataMeasure{T: 32, Time: "2019-08-24T09:35:06.654Z"},
			Err:    nil,
		},
	}

	for _, testCase := range cases {
		res := MinValueInArray(testCase.data)
		if res != testCase.Result {
			t.Errorf("Excepted %v, got is %v", testCase.Result, res)
		}
	}

}
