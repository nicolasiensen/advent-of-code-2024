package level

func Safe(level1 int, level2 int, increasing bool) bool {
	delta := level1 - level2

	if delta > 3 || delta < -3 || delta == 0 {
		return false
	}

	if level1 < level2 && !increasing {
		return false
	} else if level1 > level2 && increasing {
		return false
	}
	return true
}
