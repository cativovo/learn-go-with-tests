package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{
			time:  simpleTime(0, 0, 30),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 0),
			angle: 0,
		},
		{
			time:  simpleTime(0, 0, 45),
			angle: (math.Pi / 2) * 3,
		},
		{
			time:  simpleTime(0, 0, 7),
			angle: (math.Pi / 30) * 7,
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)
			want := test.angle
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{
			time: simpleTime(0, 0, 30),
			point: Point{
				X: 0,
				Y: -1,
			},
		},
		{
			time: simpleTime(0, 0, 45),
			point: Point{
				X: -1,
				Y: 0,
			},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondHandPoint(test.time)
			want := test.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{
			time:  simpleTime(0, 30, 0),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 7),
			angle: 7 * (math.Pi / (30 * 60)),
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := minutesInRadians(test.time)
			want := test.angle
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{
			time: simpleTime(0, 30, 0),
			point: Point{
				X: 0,
				Y: -1,
			},
		},
		{
			time: simpleTime(0, 45, 0),
			point: Point{
				X: -1,
				Y: 0,
			},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := minuteHandPoint(test.time)
			want := test.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{
			time:  simpleTime(6, 0, 0),
			angle: math.Pi,
		},
		{
			time:  simpleTime(0, 0, 0),
			angle: 0,
		},
		{
			time:  simpleTime(21, 0, 0),
			angle: math.Pi * 1.5,
		},
		{
			time:  simpleTime(0, 1, 30),
			angle: math.Pi / ((6 * 60 * 60) / 90),
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := hoursInRadians(test.time)
			want := test.angle
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHourhandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{
			time: simpleTime(6, 0, 0),
			point: Point{
				X: 0,
				Y: -1,
			},
		},
		{
			time: simpleTime(21, 0, 0),
			point: Point{
				X: -1,
				Y: 0,
			},
		},
	}

	for _, test := range cases {
		got := hourHandPoint(test.time)
		want := test.point
		if !roughlyEqualPoint(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05") // HH:MM:SS
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
