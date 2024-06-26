package tour

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"
)

type person struct {
	name   string
	height float32
}
type typeStudent struct {
	Name             string `json:"first name"`
	Age              int    `json:"AGE"`
	ElectiveLanguage string `json:"elective language,omitempty"`
}
type typeClass struct {
	Name     string `json:"class"`
	Students []typeStudent
}
type Coordinate struct {
	Lat, Long float64
}
type Star struct {
	Name string
	RA   string
	DE   string
	Mag  float32
	Dist float32
}

const (
	B10   = true
	Num10 = 10
	F10   = 10.0
	S10   = "#10"
)

const (
	NoInnerPlanetNumber = iota
	Mercury
	Venus
	Earth
	Mars
)

var BrightStars = []Star{
	{Name: "Sirius", RA: "06h 45m 08.917s", DE: "−16° 42′ 58.02″", Mag: -1.46, Dist: 8.6},
	{Name: "Canopus", RA: "06h 23m 57.10988s", DE: "−52° 41′ 44.3810″", Mag: -0.74, Dist: 310},
	{Name: "Rigil Kentaurus", RA: "14h 39m 36.49400s", DE: "−60° 50′ 02.3737″", Mag: +0.01, Dist: 4.34},
	{Name: "Arcturus", RA: "14h 15m 39.7s", DE: "+19° 10′ 56″", Mag: -0.05, Dist: 37},
	{Name: "Vega", RA: "18h 36m 56.33635s", DE: "+38° 47′ 01.2802″", Mag: +0.026, Dist: 25},
	{Name: "Capella", RA: "05h 16m 41.35871s", DE: "+45° 59′ 52.7693″", Mag: +0.08, Dist: 43},
	{Name: "Rigel", RA: "05h 14m 32.27210s", DE: "−08° 12′ 05.8981″", Mag: +0.13, Dist: 860},
}

// map literal
var Coordinates = map[string]Coordinate{
	"Buenos Aires": {
		-34.6156548, -58.515698,
	},
	"New York": {
		40.6976312, -74.1444846,
	},
	"Prague": {
		50.0597507, 14.3832245,
	},
}

func CoordinatesByName(name string) {
	coord := Coordinates[name]
	fmt.Printf("%s coordinates: %v\n", name, coord)
}
func InnerPlanetsOrder() {
	fmt.Printf("Order of inner planets: Mercury:%d, Venus:%d, Earth:%d, Mars:%d \n", Mercury, Venus, Earth, Mars)
}
func unmutableSlice_rgbNames() []string {
	return []string{"red", "green", "blue"}
}

var (
	type_b          bool
	type_string     string
	type_int        int  //32 or 64 and types int8  int16  int32  int64
	type_uint       uint //32 or 64 and types uint8  uint16  uint32  uint64
	type_uintptr    uintptr
	type_byte       byte // alias for uint8
	type_rune       rune // alias for int32 represents a Unicode code point
	type_float32    float32
	type_float64    float64
	type_complex64  complex64
	type_complex128 complex128

	type_person person
	personJoe   = person{name: "Joe Doe", height: 180}

	primes10   = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	rgbNames   = []string{"red", "green", "blue"}    // slice
	rgbNames3  = [3]string{"red", "green", "blue"}   // array len 3 cap 3
	rgbNames3c = [...]string{"red", "green", "blue"} // array len 3 cap 3
)

func GlobalVariables() {
	type_b = true //false
	type_string = "ABC"
	type_int = -1
	type_uint = 1_000_000 // MILLION
	fmt.Printf("%v %v %v %v %v \n", type_b, type_string, type_int, type_uint, type_uintptr)
	type_byte = 255
	type_rune = 255
	type_float32 = 1e38
	type_float64 = 1e308
	type_complex64 = 1 + 1i
	type_complex128 = 1 + 1i
	fmt.Printf("%v %v %v %v %v %v \n", type_byte, type_rune, type_float32, type_float64, type_complex64, type_complex128)
	type_person = person{name: "Joe", height: 184.5}
	fmt.Printf("%v \n", type_person)
	fmt.Printf("%+v \n", type_person)
	fmt.Printf("%#v \n", type_person)
}

func IntMinMax() {
	i32 := "int"
	u32 := "uint"
	i64 := ""
	u64 := ""
	if math.MaxInt == math.MaxInt64 {
		i32 = ""
		u32 = ""
		i64 = "int"
		u64 = "uint"
	}
	// int
	fmt.Printf("Integer types min and max:\n")
	// int
	fmt.Printf("  int8: %21d .. %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf(" int16: %21d .. %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf(" int32: %21d .. %d %s alias rune\n", math.MinInt32, math.MaxInt32, i32)
	fmt.Printf(" int64: %21d .. %d %s\n", math.MinInt64, math.MaxInt64, i64)
	// unsigned
	fmt.Printf(" uint8: %21d .. %d alias byte\n", 0, math.MaxUint8)
	fmt.Printf("uint16: %21d .. %d\n", 0, math.MaxUint16)
	fmt.Printf("uint32: %21d .. %d %s\n", 0, math.MaxUint32, u32)
	fmt.Printf("uint64: %21d .. %d %s\n", 0, uint64(math.MaxUint64), u64)
}

func FloatMinMax() {
	fmt.Printf("Float types min and max:\n")
	fmt.Printf("float32: %24.7e .. %12.7e\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64: %24.15e .. %22.15e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}

func QuotientAndModulo(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	modulo := dividend % divisor // % is modulus operator (remainder operator), result is integer reminder
	return quotient, modulo
}

func SystemInt() {
	systemInt := 32
	if math.MaxInt == math.MaxInt64 {
		systemInt = 64
	}
	fmt.Printf("This is %d bit system.\n", systemInt)
}

func Loops() {
	n := 1
	for n < 4 {
		fmt.Printf("%d: square:%d\n", n, n*n)
		n += 1
	}
	for i := 1; i < 4; i++ {
		fmt.Printf("%d: square:%d\n", i, i*i)
	}
	fmt.Println()
	for i, s := range []string{"a", "b"} {
		fmt.Printf("%d:%s\n", i, s)
	}
	fmt.Print("int range example: ")
	// since GO 1.22
	iMax := 5
	for i := range iMax {
		fmt.Printf("0<=%d<%d; ", i, iMax)
	}
	fmt.Println()
}

func GotoFunc() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
LOOP1:
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP1
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func Switches() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Expected OS Win: ", os)
	case "darwin":
		fmt.Println("Expected OS X: ", os)
	case "linux":
		fmt.Println("Expected OS Linux: ", os)
	default:
		// freebsd, openbsd,plan9, windows...
		fmt.Printf("Unexpected OS %s.\n", os)
	}

	now := time.Now()
	var d2s string
	switch m2 := now.Minute() % 2; m2 {
	case 0:
		d2s = "even"
	default:
		d2s = "odd"
	}
	fmt.Printf("Time %s has %s minute\n", now.String(), d2s)

	//Switch without a condition is the same as switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func PrintStringLiterals() {
	fmt.Printf("rgbNames %%v: %v \n", rgbNames) // [red green blue]
	fmt.Printf("rgbNames %%q: %q \n", rgbNames) // ["red" "green" "blue"]
	if len(rgbNames3) == len(rgbNames3c) {
		fmt.Println("[...]T is syntax sugar for [3]T")
	}
	fmt.Printf("unmutableSlice len=%d cap=%d\n", len(unmutableSlice_rgbNames()), cap(unmutableSlice_rgbNames()))

	fmt.Println("Person:", personJoe)
}

func PrintIntLiterals() {
	fmt.Print("The first 10 prime numbers: ")
	for i, n := range primes10 {
		fmt.Printf("%d:%d ", i+1, n)
	}
	fmt.Println()
}

func PrintJsonStudents() {
	studentsJson := `[{"name":"Chuck","age": 17,"elective language": "French"},{"name":"Dan","age":19}]`
	var students []typeStudent
	json.Unmarshal([]byte(studentsJson), &students)
	fmt.Printf("Students : %+v\n\n", students)

	var newClassOfStudents = typeClass{Name: "Language class", Students: []typeStudent{{Name: "Ann", Age: 17, ElectiveLanguage: "French"}, {Name: "Betty", Age: 16, ElectiveLanguage: "German"}}}
	bytes, _ := json.Marshal(newClassOfStudents)
	fmt.Printf("New class students in json format: %s\n\n", string(bytes))
}

func SortSliceExample() {
	var people = make([]person, 0)
	people = append(people, person{name: "Bob", height: 195.0})
	people = append(people, person{name: "Ann", height: 150.0})
	people = append(people, person{name: "Chuck", height: 182.0})
	sort.SliceStable(people, func(i, j int) bool { return people[i].height < people[j].height })
	fmt.Printf("Sorted by height: %v\n", people)
}
func PrintCurrentTime() {
	t := time.Now()

	fmt.Println(time.Layout + " ... time.Layout - The reference time, in numerical order")
	fmt.Println(t.Format(time.Layout))
	fmt.Println()

	fmt.Println(time.UnixDate + " ... time.UnixDate")
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println()

	userFormat := "2006-01-02 15:04 MST ~ 1/2/06 3/4 PM ~ Mon Jan 2 UTC-0700"
	fmt.Println(userFormat + " ... user format example")
	fmt.Println(t.Format(userFormat))
}
func LoadTextFile(filename string) string {
	bytes, err := os.ReadFile(filename) //Read entire file content. No need to close
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(bytes)
}
func VariadicSum(numbers ...int) {
	s := fmt.Sprintf("%d numbers sum: ", len(numbers))
	total := 0
	plus := ""
	for _, num := range numbers {
		total += num
		s += plus + strconv.Itoa(num)
		plus = " + "
	}
	fmt.Printf("%s = %d\n", s, total)
}
func PrintBrightStars() {
	for _, star := range BrightStars {
		fmt.Printf("  %-20s %-18s %-18s %5.2f  %5.1f  \n", star.Name, star.RA, star.DE, star.Mag, star.Dist)
	}
	fmt.Println()
}
