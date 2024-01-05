package moonView

var part1 = `<!DOCTYPE html>
<html lang="en" xmlns="http://www.w3.org/1999/html">
`

/*
[
{"name":"Center", "lat":0.0,"lon":0.0},
{"name":"North Pole", "lat":90.0,"lon":0.0},
{"name":"South Pole", "lat":-90.0,"lon":0.0},
{"name":"Left edge", "lat":0.0,"lon":-90.0},
{"name":"Right edge", "lat":0.0,"lon":90.0},

{"name":"Anděl","lat":-10.4, "lon":12.4},
{"name":"Aristarchus", "lat":23.7,"lon":-47.4},
{"name":"Atlas",     "lat":46.7, "lon":44.4},
{"name":"Clavius", "lat":-58.4,"lon":-14.4},
{"name":"Cleomedes", "lat":27.7, "lon":55.5},
{"name":"Copernicus","lat": 9.62, "lon":-20.08},
{"name":"Cyrillus",  "lat":-13.2, "lon":24.0},
{"name":"Kepler","lat":  8.1, "lon":-38.00},
{"name":"Marius","lat":  11.9, "lon":-50.80},
{"name":"Plato", "lat": 51.6, "lon": -9.30},
{"name":"Posidonius","lat":31.7, "lon":30.1},
{"name":"Russel","lat":26.5, "lon":-75.4},
{"name":"Schickard","lat":-44.4, "lon":-55.1},
{"name":"Schomberger","lat":-76.7, "lon":24.9},
{"name":"Seleucus", "lat":21.0,"lon":-66.6},
{"name":"Tycho", "lat":-43.31,"lon":-11.36},

{"name":"Langrenus","lat":-8.9,"lon":60.9}
]
END;
*/

var part2 = `
<head>
    <meta charset="UTF-8">
    <link rel="icon" type="image/x-icon" href="./favicon.ico">
    <title>Moon</title>
    <style>
        #svg1 {
            border: 1px solid red
        }
        #moon_hour_resources {
            display: none;
        }
        .center {
            text-align: center
	}

        .container {
			position: relative;
			display: inline-block;
		}
		
	.tooltip {
			position: absolute;
			top: -30px;
			left: 0;
			display: none;
			padding: 5px;
			background-color: #000;
			color: #fff;
			font-size: 12px;
		}
    </style>
    <script>
        var startTime, endTime;
        function toggleMoonHourResources() {
            var x = document.getElementById("moon_hour_resources");
            console.log('x.style.display='+x.style.display)
            if (x.style.display === "" || x.style.display === "none") {
                x.style.display = "block";
            } else {
                x.style.display = "none";
            }
        }
        function yesterday(){
            var x = document.getElementById("date");
            var str = x.value;
            let d = new Date(str);
            d.setDate(d.getDate() - 1);
            let yesterday = d.toLocaleDateString('en-CA', {timeZone:'GMT'});
            x.value = yesterday;
            console.log("str=${str} Yesterday: ${yesterday}");
        }
        function tomorrow(){
            var x = document.getElementById("date");
            var str = x.value;
            let d = new Date(str);
            d.setDate(d.getDate() + 1);
            //let tomorrow  = d.toLocaleDateString('en-CA', {timeZone:'Asia/Shanghai'});
            let tomorrow  = d.toLocaleDateString('en-CA', {timeZone:'GMT'});
            x.value = tomorrow;
            console.log("str=${str} Tomorrow: ${tomorrow}");
	    }

	var craters = [

	{"x":10123.0,"y":10123.0, "r":1.0, "n":"remove me"}
	];
    </script>
</head>
<body>
<div class="center">
`
var part3 = `
{{.GetParams}} <br/>
v0.1 Moon at {{.TimeInfo}}
    <br/>
    <svg width="730" height="730" class="image" >
        <defs>
            <pattern id="img1" patternUnits="userSpaceOnUse" width="730" height="730">
			    <image xlink:href="{{.SVSframes}}730x730_1x1_{{.P36}}0p/moon{{.Hours}}jpg"
                       x="0" y="0" width="730" height="730"/>
            </pattern>
            <!--?php moon_draw(); ?-->
        </defs>
        <rect x="0" y="0" width="730" height="730" fill="url(#img1)"/>
        <circle cx="365" cy="365" r="{{.Radius}}" stroke="yellow" stroke-width="1" stroke-dasharray="2 10"
                fill="none"/>
        <!--use xlink:href="#moon_drawing" x="365" y="365"/-->
            <?php moon_draw(); ?>
    </svg>
    <div class="tooltip"></div>
    <br/>
    Source: <a href="{{.SVSframes}}730x730_1x1_{{.P36}}0p/moon{{.Hours}}jpg"  target="_blank" >moon image</a>
    Credit: NASA 
    <a href="https://svs.gsfc.nasa.gov/{{.SVSmagic1}}" target="_blank">Scientific Visualization Studio</a>
     - <a href="https://svs.gsfc.nasa.gov/Gallery/moonphase.html" target="_blank">Moon Phase and Libration</a> Gallery
    <br/>

    <!--form action="/moon"-->
    <form action="./">
        <label>UTC date:</label>
        <input type="button" onclick="yesterday()" value="&lt;">
        <input type="date" id="date" name="date" min="2011-01-01" max="{{.MaxYear}}-12-31"
               value="{{.CurrentDate}}">
        <input type="button" onclick="tomorrow()" value="&gt;">
        <label>UTC hour:</label>
        <input type="number" id="utc_hour" name="utc_hour" min="0" max="23" size="2" value="{{.UTChour}}">
        grid: <input type="checkbox" name="grid" id="grid" >
        <input type="submit" value="Submit">
        show info: <input type="checkbox" name="showinfo" id="showinfo"  onchange="toggleMoonHourResources()">
    </form>
`
var part_moon_hour_resources = `
<div id="moon_hour_resources">
	Year {{.YYYY}} images folder:
	<a href="{{.SVSframes}}"  target="_blank" >{{.SVSframes}}</a> <br/>
	<a href="{{.SVSframes}}1080x1080_1x1_{{.P36}}0p/orbit{{.Hours}}tif"  target="_blank" >1080</a>,
	1920x1080:
	<a href="{{.SVSframes}}1920x1080_16x9_{{.P36}}0p/distance/dist{{.Hours}}tif"  target="_blank" >d</a>
	-
	<a href="{{.SVSframes}}1920x1080_16x9_{{.P36}}0p/fancy/comp{{.Hours}}tif"  target="_blank" >comp</a>
	-
	<a href="{{.SVSframes}}1920x1080_16x9_{{.P36}}0p/labels/label{{.Hours}}png"  target="_blank" >l</a>
	-
	<a href="{{.SVSframes}}1920x1080_16x9_{{.P36}}0p/plain/moon{{.Hours}}tif"  target="_blank" >m</a>,
	hw abc,
	<a href="{{.SVSframes}}216x216_1x1_{{.P36}}0p/moon{{.Hours}}jpg"  target="_blank" >216</a>,
	<a href="{{.SVSframes}}320x320_1x1_{{.P36}}0p/globe{{.Hours}}tif"  target="_blank" >320</a>,
	3840x2160:
	<a href="{{.SVSframes}}3840x2160_16x9_{{.P36}}0p/distance/dist{{.Hours}}tif"  target="_blank" >d</a>
	-
	<a href="{{.SVSframes}}3840x2160_16x9_{{.P36}}0p/fancy/comp{{.Hours}}tif"  target="_blank" >comp</a>
	-
	<a href="{{.SVSframes}}3840x2160_16x9_{{.P36}}0p/labels/label{{.Hours}}tif"  target="_blank" >l</a>
	-
	<a href="{{.SVSframes}}3840x2160_16x9_{{.P36}}0p/plain/moon{{.Hours}}tif"  target="_blank" >m</a>,
	<a href="{{.SVSframes}}420x420_1x1_{{.P36}}0p/orbit{{.Hours}}tif"  target="_blank" >420</a>,
	5760x3240:
	<a href="{{.SVSframes}}5760x3240_16x9_{{.P36}}0p/distance/dist{{.Hours}}tif"  target="_blank" >d</a>
	-
	<a href="{{.SVSframes}}5760x3240_16x9_{{.P36}}0p/exr/moon{{.Hours}}exr"  target="_blank" >e</a>
	-
	<a href="{{.SVSframes}}5760x3240_16x9_{{.P36}}0p/fancy/comp{{.Hours}}tif"  target="_blank" ><b>comp</b></a>
	-
	<a href="{{.SVSframes}}5760x3240_16x9_{{.P36}}0p/labels/label{{.Hours}}png"  target="_blank" >l</a>
	-
	<a href="{{.SVSframes}}5760x3240_16x9_{{.P36}}0p/plain/moon{{.Hours}}tif"  target="_blank" >m</a>,
	<a href="{{.SVSframes}}640x640_1x1_{{.P36}}0p/globe{{.Hours}}tif"  target="_blank" >640</a>,
	<a href="{{.SVSframes}}730x730_1x1_{{.P36}}0p/moon{{.Hours}}jpg"  target="_blank" >730</a>,
	<a href="{{.SVSframes}}850x850_1x1_{{.P36}}0p/orbit{{.Hours}}tif"  target="_blank" >850</a>,
	<a href="{{.SVSframes}}960x960_1x1_{{.P36}}0p/globe{{.Hours}}tif"  target="_blank" >960</a>,
	mooninfo:
	<a href="{{.SVSframes}}../mooninfo_2024.txt"   target="_blank" >txt</a>,
	<a href="{{.SVSframes}}../mooninfo_2024.json"  target="_blank" >json</a><br/>
	{"time":"{{.Time}}","phase":{{.Phase}},"age":{{.Age}},"diameter":{{.Diameter}},"distance":{{.Distance}},<br/>
	"j2000":{"ra":{{.RA}},"dec":{{.Dec}}},<br/>
	"subsolar":{"lon":{{.Slon}},"lat":{{.Slat}}},<br/>
	"subearth":{"lon":{{.Elon}},"lat":{{.Elat}}},<br/>
	"posangle":{{.Posangle}} }<br/>
</div>

`

var part4 = `
</div>
<script>
		const image = document.querySelector('.image');
		const tooltip = document.querySelector('.tooltip');
         
        craters.sort(function(a, b){return a.r - b.r});

        function positionMessage(e){       
            tooltip.style.left = e.clientX + 'px';
		    tooltip.style.top = e.clientY - tooltip.offsetHeight - 10 + 'px';
            x = e.offsetX;
            y = e.offsetY;
            console.log("positionMessage x="+x+",y="+y)

            x1 = x - 365
            y1 = y - 365
            clen = craters.length
            for (let i = 0; i < clen; i++) {             
              if (craters[i].x < 9123.0) {
                //console.log(i, craters[i].n)
                 dx = craters[i].x - x1
                 dy = craters[i].y - y1
                 r1 = Math.sqrt(dx*dx + dy*dy)
                 if (r1 < craters[i].r) {                   
                    tooltip.textContent = "X: ${e.offsetX}, Y: ${e.offsetY} "+craters[i].n;
                 }
              }
  		    }
            tooltip.textContent = "X: ${e.offsetX}, Y: ${e.offsetY} len=${clen}"
        }
		image.addEventListener('mousemove', positionMessage)
        
		image.addEventListener('xmousemove11', (e) => {
		  tooltip.style.left = e.clientX + 'px';
		  tooltip.style.top = e.clientY - tooltip.offsetHeight - 10 + 'px';
		  tooltip.textContent = "X: ${e.offsetX}, Y: ${e.offsetY} v2";
        });

		image.addEventListener('mouseover', () => {
		  tooltip.style.display = 'block';
		});

		image.addEventListener('mouseout', () => {
		  tooltip.style.display = 'none';
		});

        image.addEventListener("click", function(e){
			x = e.offsetX
        	y = e.offsetY;
            x1 = x - 365
            y1 = y - 365
            clen = craters.length
            for (let i = 0; i < clen; i++) {             
              ///if (craters[i].x < 9123.0) {
                //console.log(i, craters[i].n)
                 dx = craters[i].x - x1
                 dy = craters[i].y - y1
                 r1 = Math.sqrt(dx*dx + dy*dy)
                 if (r1 < craters[i].r) {
                    //alert("x="+x+",y="+y+"\ncrater="+craters[i].n);
                    tooltip.textContent = "X: ${e.offsetX}, Y: ${e.offsetY} "+craters[i].n;
                 }
              ///}
}
			//alert("Hello Click! x="+x+",y="+y+" craters len="+craters.length);
		});
	</script>
</body>
</html>
`
