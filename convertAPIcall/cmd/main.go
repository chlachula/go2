package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	// export CONVERTAPI_SECRET=my-secret; echo $CONVERTAPI_SECRET
	var secret string
	if secret = os.Getenv("CONVERTAPI_SECRET"); secret == "" {
		fmt.Println("Error: Environment variable CONVERTAPI_SECRET not set")
		return
	}
	config.Default = config.NewDefault(secret) // Get your secret at https://www.convertapi.com/a

	//User information
	if user, err := convertapi.UserInfo(nil); err == nil {
		fmt.Println("User information: ")
		fmt.Printf("%+v\n", user)
	} else {
		fmt.Println(err)
	}

	//Conversion
	fromPath := "cen-cal-joe-doe.svg"
	toPath := "joe-doe-cal.png"
	if file, errs := convertapi.ConvertPath(fromPath, toPath); errs == nil {
		fmt.Println("Result of the file conversion saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}

	//Conversion notes 2024-12-05
	// SVG width="980.0" height="1180.0"
	// Compression: PNG ZIP
	// Resolution: 200x200 DPI
	// Current size: 980 x 1180  Pixels (1.16 MPixels) (0.83)
	// 12.4 x 15.0 cm; 4.90 x 5.90 inches
}
