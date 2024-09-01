package clockface_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"
	"time"

	"github.com/cativovo/learn-go-with-tests/math/clockface"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Line    []Line   `xml:"line"`
	Circle  Circle   `xml:"circle"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{
				X1: 150,
				Y1: 150,
				X2: 150,
				Y2: 60,
			},
		},
		{
			time: simpleTime(0, 0, 30),
			line: Line{
				X1: 150,
				Y1: 150,
				X2: 150,
				Y2: 240,
			},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			var b bytes.Buffer
			if err := clockface.SVGWriter(&b, test.time); err != nil {
				t.Fatalf("not expecting error %s", err)
			}

			var svg SVG
			if err := xml.Unmarshal(b.Bytes(), &svg); err != nil && err != io.EOF {
				t.Fatalf("not expecting error %s", err)
			}

			if !containsLine(test.line, svg.Line) {
				t.Errorf("expected to find %+v in %+v", test.line, b.String())
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{
				X1: 150,
				Y1: 150,
				X2: 150,
				Y2: 70,
			},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			var b bytes.Buffer
			if err := clockface.SVGWriter(&b, test.time); err != nil {
				t.Fatalf("not expecting error %s", err)
			}

			var svg SVG
			if err := xml.Unmarshal(b.Bytes(), &svg); err != nil && err != io.EOF {
				t.Fatalf("not expecting error %s", err)
			}

			if !containsLine(test.line, svg.Line) {
				t.Errorf("expected to find %+v in %+v", test.line, b.String())
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(6, 0, 0),
			line: Line{
				X1: 150,
				Y1: 150,
				X2: 150,
				Y2: 200,
			},
		},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			var b bytes.Buffer
			if err := clockface.SVGWriter(&b, test.time); err != nil {
				t.Fatalf("not expecting error %s", err)
			}

			var svg SVG
			if err := xml.Unmarshal(b.Bytes(), &svg); err != nil && err != io.EOF {
				t.Fatalf("not expecting error %s", err)
			}

			if !containsLine(test.line, svg.Line) {
				t.Errorf("expected to find %+v in %+v", test.line, b.String())
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05") // HH:MM:SS
}

func containsLine(line Line, lines []Line) bool {
	for _, v := range lines {
		if v == line {
			return true
		}
	}
	return false
}
