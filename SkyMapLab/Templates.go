package SkyMapLab

const (
	svgTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
    xmlns:xlink="http://www.w3.org/1999/xlink" 
	width="{{.WidthMM}}mm" height="{{.HeightMM}}mm" viewBox="{{.VBminX}} {{.VBminY}} {{.VBwidth}} {{.VBheight}}" 
	style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd;background:beige">
    <title>Sky Map Lab</title>
 <defs>
    <style>
	 .font1 { 
		font-size: {{.FontSize}}px;
		font-family: Franklin Gothic, sans-serif;
		font-weight: 90; 		
		letter-spacing: 2px;
	 }
	 .fontLegend { 
		font-size: {{.FontSizeLegend}}px;
		font-family: Franklin Gothic, sans-serif;
    font-weight: 90; 
    letter-spacing: 2px;
	 }
	 .upFont { 
		fill: {{.TopColor}};
	 }
	 .downFont { 
		fill: {{.BottomColor}};
	 }
	 .fontAxis{
		font-size: {{.FontSizeAxis}}px;
		font-family: Franklin Gothic, sans-serif;
	 }
	 .fontObj{
		font-size: {{.FontSizeObj}}px;
		font-family: Franklin Gothic, sans-serif;
	 }
	 .board {
		stroke:orange;
		stroke-width:0.5
		fill:pink;
	 }
	 .cross {
		stroke:black;
		stroke-width: {{.CrossStrokeWidth}};
		fill:none
	 }
    </style>
	<marker 
      id='arrow_head' 
      orient="auto" 
      markerWidth='15' 
      markerHeight='20' 
      refX='0.1' 
      refY='4'
     >
     <path d='M0,0 V8 L4,4 Z' fill="black" />
    </marker>
    <pattern id="PatternOpenCluster" width="1" height="1" patternContentUnits="objectBoundingBox">
<path style="stroke:{{.LegendColor}};stroke-width:0.05" d="M0.50,0.50 h0.06
M0.71,0.40 h0.06
M0.96,0.50 h0.06
M0.71,0.60 h0.06
M0.78,0.86 h0.06
M0.55,0.72 h0.06
M0.39,0.95 h0.06
M0.35,0.68 h0.06
M0.08,0.69 h0.06
M0.27,0.49 h0.06
M0.09,0.28 h0.06
M0.36,0.31 h0.06
M0.42,0.05 h0.06
M0.56,0.28 h0.06
M0.81,0.16 h0.06
" />
    </pattern>	
    <pattern id="PatternDiffuseNebula" width="0.11" height="0.11" patternContentUnits="objectBoundingBox">
      <rect x="0" y="0" width="0.125" height=".0085" fill="{{.LegendColor}}"/>
    </pattern>
    <pattern id="PatternPlanetaryNebula" width="1" height="1" patternContentUnits="objectBoundingBox">
      <g style="fill:none;stroke:{{.LegendColor}};stroke-width:0.02">
       <circle cx="0.50" cy="0.50" r="0.17" />
       <path d="M0.65,0.57 L0.95,0.71  A1,1 0 0,0 0.95,0.29  L0.65,0.43     
                M0.35,0.57 L0.05,0.71  A1,1 0 0,1 0.05,0.29  L0.35,0.43 " />
     </g>
    </pattern>
    <pattern id="PatternSupernovaRemnant" width="1" height="1" patternContentUnits="objectBoundingBox">
	<path style="stroke:{{.LegendColor}};stroke-width:0.02" d="M0.67,0.50 L1.00,0.50
    M0.66,0.56 L0.97,0.67
    M0.63,0.61 L0.88,0.82
    M0.58,0.64 L0.75,0.93
    M0.53,0.66 L0.59,0.99
    M0.47,0.66 L0.41,0.99
    M0.42,0.64 L0.25,0.93
    M0.37,0.61 L0.12,0.82
    M0.34,0.56 L0.03,0.67
    M0.33,0.50 L0.00,0.50
    M0.34,0.44 L0.03,0.33
    M0.37,0.39 L0.12,0.18
    M0.42,0.36 L0.25,0.07
    M0.47,0.34 L0.41,0.01
    M0.53,0.34 L0.59,0.01
    M0.58,0.36 L0.75,0.07
    M0.63,0.39 L0.88,0.18
    M0.66,0.44 L0.97,0.33
    " />
    </pattern>

    {{.Defs}}

	<g id="draw_AZ_grid" transform="rotate(180)">
	  <circle cx="0" cy="0" r="{{.RLat}}" stroke="black" stroke-width="0.5" fill="none" />
      <use xlink:href="#plotHorizon" />
      <use xlink:href="#plotAlmucantarats" />
	  <use xlink:href="#plotMeridians" />	
    </g>
  <g id="draw_map">
    <use xlink:href="#plotConstellations" />
    <use xlink:href="#plotZenithBelts" />
    <use xlink:href="#plotConstellationNames" />
    <use xlink:href="#plotOuterCircle" />
    <use xlink:href="#plotEcliptic" />
    <use xlink:href="#plotStars" />
    <use xlink:href="#plotObjects" />
    <use xlink:href="#plotDateRoundScale" />
    <use xlink:href="#plotRaHourScale" />
    <use xlink:href="#plotRaCross" />
    <use xlink:href="#plotAxisDeclinations" />
    <use xlink:href="#plotDirectionsOfTheApparentRotationOfTheSky" />	
    <use xlink:href="#plotObjectsLegend" />	
  </g>
  <g id="draw_platonYear_map">
    <use xlink:href="#plotPlatonYear" />
    <use xlink:href="#draw_map" />
  </g>
  <g id="draw_all">
    <use xlink:href="#plotPlatonYear" />
    <use xlink:href="#draw_map" />
    <use xlink:href="#draw_AZ_grid" />
  </g>
 </defs> 
  <rect width="500" height="{{.Height}}" x="-250" y="-{{.HeightHalf}}" stroke="blue" stroke-width="1" fill="azure" />
  <text x="-244" y="-{{.HeightHalf}}" fill="blue" font-size="8"><tspan dy="10">{{.PaperName}} ({{.WidthMM}}mm by {{.HeightMM}}mm) - Latitude: {{.Latitude}}</tspan></text>
  
  <use xlink:href="#{{.Draw}}" />
</svg>
`
	html1 = `<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>SkyMap</title>
  <script>
	function openSkyMap(f){
	  // f == this form
	  // http://localhost:8080/img/svg/skymap/co/n44/2/x
	  let u = '/img/svg/skymap/';
	  u +=  f.color_style.value + '/'
	  u +=  f.hemisphere.value + f.latitude.value + '/'
	  u +=  f.paper.value + '/'
	  u +=  f.draw.value 
	  
	  //alert('Hello from SkyMap submit! url='+u);		
	  //window.location.href = u; // in the same tab

	  const anchor = document.createElement("a");
	  anchor.href = u;
	  anchor.target = "_blank"; // Open in a new tab
	  anchor.click();
	}
  </script>	
</head>
<body style=\"text-align: center;\">
<h1>SkyMap Lab select page</h1>
  <form action="javascript:;" onsubmit=" openSkyMap( this ) ">  
   <select name="hemisphere" id="hemisphere">
	  <option value="s" >S</option>
	  <option value="n" selected="selected">N</option>
   </select>

   <label for="latitude">Latitude:</label>
   <input type="number" id="latitude" name="latitude" value="44" step="1"  min="0" max="90" size="2">
   
   <br/>
   <label for="color">Color:</label>
   <input type="radio" id="co" name="color_style" value="co" checked="checked">
   <label for="bw">Black &amp; White</label>   
   <input type="radio" id="bw" name="color_style" value="bw">
   
  <br/>
   <input type="checkbox" id="starnames" name="starnames" value="starnames" >
   <label for="starnames">Star names</label> 
  <br/>
   <input type="checkbox" id="zb" name="zenith_belt" value="zb" >
   <label for="zb">Zenith belt</label> 
  <br/>
   <input type="checkbox" id="ply" name="platon_year" value="ply" >
   <label for="ply">Platon year</label> 
  <br/>
  <br/>
   <input type="checkbox" id="az" name="AZ grid" value="az" >
   <label for="az">AZ grid</label> 
  <br/>

   <select name="paper" id="paper" title="paper">
	  <option value="0" title="297x210">A4</option>
	  <option value="1" title="420x297" >A3</option>
	  <option value="2" title="215.9x279.4" selected="selected">Letter 8.5x11</option>
	  <option value="3" title="215.9x355.6"  >Legal 8.5x14</option>
	  <option value="4" title="279.4x431.8" >Ledger 11x17</option>
   </select>
   <select name="draw" id="draw" title="draw">
	  <option value="0" selected="selected">Map + Platon Year</option>
	  <option value="1">AZ grid</option>
	  <option value="2">Map only</option>
	  <option value="3">All</option>
   </select>
   <br/>
   <br/>
   <input type="reset" value="RESET">
   <input type="submit" value="SUBMIT">

  </form>
</body>
</html>`
)
