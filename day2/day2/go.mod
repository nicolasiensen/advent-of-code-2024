module example.com/day2

go 1.23.2

replace example.com/level => ../level

require (
	example.com/level v0.0.0-00010101000000-000000000000
	example.com/report v0.0.0-00010101000000-000000000000
)

replace example.com/report => ../report
