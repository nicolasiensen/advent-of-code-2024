package report

import (
	"testing"
)

func TestSafeReportWithLevelsDecreasing(t *testing.T) {
	rp := []int{7, 6, 4, 2, 1}
	expect := true
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestUnsafeReportWithIncreaseOf5(t *testing.T) {
	rp := []int{1, 2, 7, 8, 9}
	expect := false
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestUnsafeReportWithDecreaseOf4(t *testing.T) {
	rp := []int{9, 7, 6, 2, 1}
	expect := false
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestUnsafeReportWithIncreaseThenDecrease(t *testing.T) {
	rp := []int{1, 3, 2, 4, 5}
	expect := false
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestUnsafeReportWithNeitherIncreaseOrDescrease(t *testing.T) {
	rp := []int{8, 6, 4, 4, 1}
	expect := false
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestSafeReportWithLevelsIncreasing(t *testing.T) {
	rp := []int{1, 3, 6, 7, 9}
	expect := true
	result := Safe(rp)

	if expect != result {
		t.Fatalf(`Safe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantSafeWithoutRemovingAnyLevel(t *testing.T) {
	rp := []int{7, 6, 4, 2, 1}
	expect := true
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantUnsafe1(t *testing.T) {
	rp := []int{1, 2, 7, 8, 9}
	expect := false
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantUnsafe2(t *testing.T) {
	rp := []int{9, 7, 6, 2, 1}
	expect := false
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantSafeByRemovingLevel2(t *testing.T) {
	rp := []int{1, 3, 2, 4, 5}
	expect := true
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantSafeByRemovingLevel3(t *testing.T) {
	rp := []int{8, 6, 4, 4, 1}
	expect := true
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}

func TestTolerantSafeWithoutRemovingAnyLevel2(t *testing.T) {
	rp := []int{1, 3, 6, 7, 9}
	expect := true
	result := TolerantSafe(rp)

	if expect != result {
		t.Fatalf(`TolerantSafe(%v) = %v but it should be %v`, rp, result, expect)
	}
}
