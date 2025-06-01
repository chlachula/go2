package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
)

// SVGLine defines the properties for an SVG line element.
type SVGLine struct {
	X1          float64 `json:"x1"`
	Y1          float64 `json:"y1"`
	X2          float64 `json:"x2"`
	Y2          float64 `json:"y2"`
	Stroke      string  `json:"stroke"`
	StrokeWidth float64 `json:"strokeWidth"`
}

// SVGText defines the properties for an SVG text element.
type SVGText struct {
	X                float64
	Y                float64
	Content          string
	FontFamily       string
	FontSize         float64
	Fill             string
	DominantBaseline string
	TextAnchor       string
}

/*
type SVGText struct {
	X                int    `json:"x"`
	Y                int    `json:"y"`
	Content          string `json:"content"`
	FontFamily       string `json:"fontFamily"`
	FontSize         int    `json:"fontSize"`
	Fill             string `json:"fill"`
	DominantBaseline string `json:"dominantBaseline"`
	TextAnchor       string `json:"textAnchor"`
}
*/
// MoonFeature defines a feature on the moon, including its pointer lines and text label.
type MoonFeature struct {
	Name       string  `json:"name"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Side       string  `json:"side"`
	ShadowLine *SVGLine
	MainLine   *SVGLine
	Text       *SVGText
}

//  {"name":"Crater Kepler", "lat":8.1, "lon":-38.0, "side":"left"}

// SVGRect defines the properties for an SVG rect element.
type SVGRect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Fill   string  `json:"fill"`
}

// SVGCircle defines the properties for an SVG circle element.
type SVGCircle struct {
	CX   float64 `json:"cx"`
	CY   float64 `json:"cy"`
	R    float64 `json:"r"`
	Fill string  `json:"fill"`
}

// SVGImage defines the properties for an SVG image element.
type SVGImage struct {
	Href   string  `json:"href"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// PageData holds all the data needed to render the HTML page.
type PageData struct {
	PageTitle  string        `json:"pageTitle"`
	SVGWidth   float64       `json:"svgWidth"`
	SVGHeight  float64       `json:"svgHeight"`
	SVGViewBox string        `json:"svgViewBox"`
	SVGBgRect  SVGRect       `json:"svgBgRect"`
	MoonCircle SVGCircle     `json:"moonCircle"`
	MoonImage  SVGImage      `json:"moonImage"`
	Features   []MoonFeature `json:"features"`
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.PageTitle}}</title>
    <style>
      body {
          background-color: pink;
      }   
      .center {
          text-align: center;
      }   
    </style>
</head>
<body>
    <div class="center">
       <h1>{{.PageTitle}}</h1>

        <svg width="{{.SVGWidth}}" height="{{.SVGHeight}}" viewBox="{{.SVGViewBox}}" 
             xmlns="http://www.w3.org/2000/svg"
             xmlns:xlink="http://www.w3.org/1999/xlink">
            <rect x="{{.SVGBgRect.X}}" y="{{.SVGBgRect.Y}}" width="{{.SVGBgRect.Width}}" height="{{.SVGBgRect.Height}}" fill="{{.SVGBgRect.Fill}}" />

            <circle cx="{{.MoonCircle.CX}}" cy="{{.MoonCircle.CY}}" r="{{.MoonCircle.R}}" fill="{{.MoonCircle.Fill}}" />

            <image
                xlink:href="{{.MoonImage.Href}}"
                x="{{.MoonImage.X}}"
                y="{{.MoonImage.Y}}"
                width="{{.MoonImage.Width}}"
                height="{{.MoonImage.Height}}"
            />

            {{range .Features}}
            <!-- Feature: {{.Text.Content}} -->
            <!-- Shadow Line -->
            <line x1="{{.ShadowLine.X1}}" y1="{{.ShadowLine.Y1}}" x2="{{.ShadowLine.X2}}" y2="{{.ShadowLine.Y2}}" stroke="{{.ShadowLine.Stroke}}" stroke-width="{{.ShadowLine.StrokeWidth}}" />
            <!-- Main Line -->
            <line x1="{{.MainLine.X1}}" y1="{{.MainLine.Y1}}" x2="{{.MainLine.X2}}" y2="{{.MainLine.Y2}}" stroke="{{.MainLine.Stroke}}" stroke-width="{{.MainLine.StrokeWidth}}" />
            <!-- Text Label -->
            <text x="{{.Text.X}}" y="{{.Text.Y}}" font-family="{{.Text.FontFamily}}" font-size="{{.Text.FontSize}}" fill="{{.Text.Fill}}" dominant-baseline="{{.Text.DominantBaseline}}" text-anchor="{{.Text.TextAnchor}}">{{.Text.Content}}</text>
            {{end}}
       </svg>
</div>
</body>
</html>`

func computeFeatures(data *PageData) {
	var sideX, leftY, rightY, stepY float64
	var textAnchor string
	leftY = -1.9 * data.MoonCircle.R
	rightY = leftY
	stepY = 0.75 * data.MoonCircle.R

	for i, feature := range data.Features {
		fmt.Printf("%d. %s\n", i, feature.Name)
		sideX = 1.5 * data.MoonCircle.R
		var y2 float64
		if feature.Side == "left" {
			sideX = -sideX
			textAnchor = "end"
			leftY += stepY
			y2 = leftY
		} else {
			textAnchor = "start"
			rightY += stepY
			y2 = rightY
		}

		x1 := data.MoonCircle.R * math.Sin(feature.Lon*math.Pi/180.0)
		y1 := -data.MoonCircle.R * math.Sin(feature.Lat*math.Pi/180.0)
		feature.ShadowLine = &SVGLine{X1: x1, Y1: y1, X2: sideX, Y2: y2, Stroke: "lightblue", StrokeWidth: 2}
		feature.MainLine = &SVGLine{X1: x1, Y1: y1 + 3, X2: sideX, Y2: y2 + 3, Stroke: "black", StrokeWidth: 2}
		txt := &SVGText{X: 1.04 * sideX, Y: y2, Content: feature.Name, FontFamily: "Arial, sans-serif",
			FontSize: 30, Fill: "black", DominantBaseline: "middle", TextAnchor: textAnchor}
		feature.Text = txt
		data.Features[i] = feature

	}

}
func main() {
	jsonFilePath := "nakedEyeMoon.json"
	htmlOutputPath := "index.html"

	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file %s: %v", jsonFilePath, err)
	}

	var data PageData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	computeFeatures(&data)

	tmpl, err := template.New("html").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	outputFile, err := os.Create(htmlOutputPath)
	if err != nil {
		log.Fatalf("Failed to create output HTML file %s: %v", htmlOutputPath, err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Printf("Successfully generated %s from %s\n", htmlOutputPath, jsonFilePath)
}
