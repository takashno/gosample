module github.com/takashno/gosample

go 1.12

replace github.com/takashno/gosample2 => ../gosample2

require (
	github.com/takashno/gosample2 v0.0.0-20200221022707-dbe481dd8466
	golang.org/x/text v0.3.2
)
