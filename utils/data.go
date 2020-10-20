package utils

import (
	"github.com/andrewGIS/apiserver/db"
	"sort"
	"time"
)

func MaxValueInArray(inArray []db.DataMeasure) db.DataMeasure {
	m := inArray[0]
	for i, e := range inArray {
		if i == 0 || e.T > m.T {
			m = e
		}
	}
	return m
}

func MinValueInArray(inArray []db.DataMeasure) db.DataMeasure {
	m := inArray[0]
	for i, e := range inArray {
		if i == 0 || e.T < m.T {
			m = e
		}
	}
	return m
}

// Instrument for sorting DataMeasure by time
// Special type for sorting measures for device
type DateMeasureSorted []db.DataMeasure

func (p DateMeasureSorted) Len() int {
	return len(p)
}
func (p DateMeasureSorted) Less(i, j int) bool {
	fTime, _ := time.Parse(time.RFC3339, p[i].Time)
	sTime, _ := time.Parse(time.RFC3339, p[j].Time)
	return sTime.Before(fTime)
}
func (p DateMeasureSorted) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Get newest measure in all DataMeasures
func GetLastTimeMeasure(dateSlice []db.DataMeasure) db.DataMeasure {
	dateSortedMeasures := make(DateMeasureSorted, 0, len(dateSlice))
	for _, d := range dateSlice {
		dateSortedMeasures = append(dateSortedMeasures, d)
	}
	sort.Sort(dateSortedMeasures)
	return dateSortedMeasures[0]
}
