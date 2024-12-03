package level

import "testing"

func TestSafeWithBigNegativeDelta(t *testing.T) {
	level1 := 1
	level2 := 5
	increasing := true

	safe := Safe(level1, level2, increasing)

	expect := false

	if safe != expect {
		t.Fatalf(`Safe(%v, %v, %v) = %v but it should be %v`, level1, level2, increasing, safe, expect)
	}
}

func TestSafeWithBigPositiveDelta(t *testing.T) {
	level1 := 5
	level2 := 1
	increasing := true

	safe := Safe(level1, level2, increasing)

	expect := false

	if safe != expect {
		t.Fatalf(`Safe(%v, %v, %v) = %v but it should be %v`, level1, level2, increasing, safe, expect)
	}
}

func TestSafeWithNoDelta(t *testing.T) {
	level1 := 5
	level2 := 5
	increasing := true

	safe := Safe(level1, level2, increasing)

	expect := false

	if safe != expect {
		t.Fatalf(`Safe(%v, %v, %v) = %v but it should be %v`, level1, level2, increasing, safe, expect)
	}
}

func TestSafeWithIncreasingLevelButDecreasingReport(t *testing.T) {
	level1 := 1
	level2 := 3
	increasing := false

	safe := Safe(level1, level2, increasing)

	expect := false

	if safe != expect {
		t.Fatalf(`Safe(%v, %v, %v) = %v but it should be %v`, level1, level2, increasing, safe, expect)
	}
}

func TestSafeWithDecreasingLevelButIncreasingReport(t *testing.T) {
	level1 := 3
	level2 := 1
	increasing := true

	safe := Safe(level1, level2, increasing)

	expect := false

	if safe != expect {
		t.Fatalf(`Safe(%v, %v, %v) = %v but it should be %v`, level1, level2, increasing, safe, expect)
	}
}
