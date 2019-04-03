package sum

import "testing"

var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{120, 1, 119},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{10, 20, 30},
	{40, 50, 90},
	{1200, 10, 1190},
}

func TestSumUInt32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{10, 20, 30},
	{40, 50, 90},
	{1200, 10, 1190},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{100, 200, 300},
	{400, 500, 900},
	{12000, 100, 11900},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1.1, 2.2, 3.3},
	{4.4, 5.5, 9.9},
	{120.1, 1.2, 119.3},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%v, %v) returned %v, expected %v", v.n1, v.n2, val, v.expected)
		}
	}
}
