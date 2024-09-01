package clockface

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"
)

type ClockFace struct {
	SecondHand Point
	MinuteHand Point
	HourHand   Point
}

var svgTemplate = `
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" viewBox="0 0 300 300" version="2.0">
	<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>
	<line x1="150" y1="150" x2="{{printf "%.3f" .SecondHand.X}}" y2="{{printf "%.3f" .SecondHand.Y}}" style="fill:none;stroke:#f00;stroke-width:3px;"/>
	<line x1="150" y1="150" x2="{{printf "%.3f" .MinuteHand.X}}" y2="{{printf "%.3f" .MinuteHand.Y}}" style="fill:none;stroke:#f00;stroke-width:3px;"/>
	<line x1="150" y1="150" x2="{{printf "%.3f" .HourHand.X}}" y2="{{printf "%.3f" .HourHand.Y}}" style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
`

var tmpl, tmplErr = template.New("clockface").Parse(strings.TrimSpace(svgTemplate))

func SVGWriter(w io.Writer, t time.Time) error {
	if tmplErr != nil {
		return fmt.Errorf("SVGWriter: unable to parse the template: %w", tmplErr)
	}

	mh := minuteHand(t)
	sh := secondHand(t)
	hh := hourHand(t)

	err := tmpl.Execute(w, ClockFace{
		MinuteHand: mh,
		SecondHand: sh,
		HourHand:   hh,
	})
	if err != nil {
		return fmt.Errorf("SVGWriter: unable to execute the template: %w", err)
	}

	return nil
}

func minuteHand(t time.Time) Point {
	p := minuteHandPoint(t)
	return makeHand(p, minuteHandLength)
}

func secondHand(t time.Time) Point {
	p := secondHandPoint(t)
	return makeHand(p, secondHandLength)
}

func hourHand(t time.Time) Point {
	p := hourHandPoint(t)
	return makeHand(p, hourHandLength)
}

func makeHand(p Point, l float64) Point {
	// scale
	p.X *= l
	p.Y *= l
	// flip
	p.Y = -p.Y
	// translate
	p.X += clockCenterX
	p.Y += clockCenterY
	return p
}
