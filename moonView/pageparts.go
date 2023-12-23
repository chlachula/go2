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
v0.13 Moon at 2023-12-18 03:00 UTC, JD:2460296.625, 8428 hours since 2023-1-1
    <br/>
    <svg width="730" height="730" class="image" >
        <defs>
            <pattern id="img1" patternUnits="userSpaceOnUse" width="730" height="730">
			    <image xlink:href="%s"
                       x="0" y="0" width="730" height="730"/>
            </pattern>
            <!--?php moon_draw(); ?-->
        </defs>
        <rect x="0" y="0" width="730" height="730" fill="url(#img1)"/>
        <circle cx="365" cy="365" r="<?php echo $rad ?>" stroke="yellow" stroke-width="1" stroke-dasharray="2 10"
                fill="none"/>
        <!--use xlink:href="#moon_drawing" x="365" y="365"/-->
            <?php moon_draw(); ?>
    </svg>
    <div class="tooltip"></div>
    <br/>
    Source: <a href="%s">moon image</a>
    Credit: NASA Scientific Visualization Studio - <a href="https://svs.gsfc.nasa.gov/Gallery/moonphase.html">Moon Phase and Libration</a> Gallery
    <br/>

    <!--form action="/moon"-->
    <form action="./">
        <label for="birthday">UTC date:</label>
        <input type="button" onclick="yesterday()" value="&lt;">
        <input type="date" id="date" name="date" min="2011-01-01" max="<?php echo $ynext ?>-12-31"
               value="<?php echo $val_date ?>">
        <input type="button" onclick="tomorrow()" value="&gt;">
        <label for="birthday">UTC hour:</label>
        <input type="number" id="utc_hour" name="utc_hour" min="0" max="23" size="2" value="<?php echo $val_hour ?>">
        grid: <input type="checkbox" name="grid" id="grid" >
        <input type="submit" value="Submit">
        show info: <input type="checkbox" name="showinfo" id="showinfo"  onchange="toggleMoonHourResources()">
    </form>
`

var part4 = `
   <svg id="svg1" align="center" width="100" height="100" viewBox="-50 -50 125 125" xmlns="http://www.w3.org/2000/svg"
         fill-rule="evenodd" clip-rule="evenodd">
        <path d="M11.917 11.019c0-.507-.41-.918-.916-.918s-.917.411-.917.918c0 .507.411.918.917.918s.916-.411.916-.918m1.751 0c0 1.473-1.196 2.671-2.667 2.671-1.47
    0-2.667-1.198-2.667-2.671 0-1.473 1.197-2.671 2.667-2.671 1.471 0 2.667 1.198 2.667 2.671m5.125-2.679c-.827-2.397-2.722-4.29-5.117-5.113l-.118.936c1.981.741
    3.553 2.313 4.299 4.293l.936-.116zm-1.858.232c-.652-1.58-1.913-2.843-3.491-3.494l-.12.955c1.166.548 2.109 1.491 2.656 2.659l.955-.12zm-2.267
    2.447c0-2.028-1.643-3.673-3.667-3.673-2.025 0-3.667 1.645-3.667 3.673s1.642 3.673 3.667 3.673c2.024 0 3.667-1.645
    3.667-3.673m-5.991 4.987c-1.166-.549-2.107-1.492-2.654-2.66l-.954.119c.65 1.582 1.911 2.844 3.49 3.496l.118-.955zm-.238 1.906c-1.989-.747-3.569-2.332-4.308-4.329l-.935.118c.822
    2.412 2.721 4.318 5.126 5.147l.117-.936zm13.561-6.893c0 .264-.022.521-.04.78-.132-.033-.457-.114-.894-.021-.295-.486-.85-.799-1.503-.799-.685
    0-1.27.351-1.548.885-.946-.17-2.098.418-2.098 1.593v2.761c-.687-.72-2.916-.376-2.916 1.41 0 .275.062.549.185.82.066.158 1.393 2.805 1.467
    2.955-1.144.404-2.37.635-3.652.635-6.075 0-11.001-4.933-11.001-11.019 0-6.085 4.926-11.019 11-11.019s11 4.934 11 11.019m-6.302 6.286c.007.01.757
    1.39.872 1.607.124.228.494.179.494-.12v-5.335c0-.839 1.348-.814 1.348 0v4.311c0 .234.453.23.453 0l.002-5.131c0-.441.355-.656.714-.656.363
    0 .729.221.729.656v5.072c0 .235.437.244.437.006v-4.323c0-.862 1.475-.886 1.475 0v4.579c0 .233.472.234.472 0v-2.849c0-.778 1.304-.822 1.304.039l.002
    6.499c0 1.489-.831 2.34-2.406 2.34h-2.935c-1.497 0-2.022-.846-2.438-1.696-.395-.808-2.001-3.976-2.125-4.272-.066-.144-.095-.28-.095-.404 0-.809
    1.276-1.128 1.697-.323"/>
    </svg>
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
