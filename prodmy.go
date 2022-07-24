/*This is a program to solve the following problem:
Implementation generic machine algorithm code for one of my scenario implementations in
my logistic business, preferably in MATLAB. The scenario is, I have four EV types of vehicles
with their battery capacity and range. As stated, 1st EV has a 900-kWh battery with max range
of 489 km , 2nd EV has a 350-kWh battery with max range of 240 km, 3rd EV has a 150-kWh
battery with a max range of 240 km, and 4th EV has an 80-kWh battery with a max range of 240
km. These EVs will need an overall charging station which includes level 1, 2, and 3 chargers. As
you go up by charger the faster power output (charge time) it has, however, the more
expensive it gets which will be accounted for. For example: level 1 states a 7-kW power output
with a price of 210 euros , level 2 states a 17-kW power output with a price of 775 euros, and
level 3 states a 203-kW power output with a price of 23,500 euros. Not including any losses
when charging in this scenario. The charging station must accommodate for 10 of the 1st EV, 5
of the 2nd EV, 5 of the 3rd EV, and 3 of the 4th EV. However, in my business each EV can have a
standby load time where they can be charging also a preferred minimum range that it needs.
We will face this scenario, as the worst case, where all the EVs don’t have any charge, starting
from zero. The 1st EV has can only have a maximum of 160 minutes for each EV and minimum
range target of 250 km, the 2nd EV can have a maximum of 120 minutes for each EV and a
minimum range target of 150 km, 3r EV can have a maximum of 60 minutes for each EV and a
minimum range target of 120 km, lastly the 4th EV can have a maximum of 30 minutes for each
EV and minimum target range of 60 km. So, it’s needed to create a combination of level
chargers that accommodates all EVs by their maximum charge time by its respective EV while it
hits its minimum target range. Lastly the amount maximum willing to spend for a charging
station overall is 500,000 euros

*/

package main
import(
	"fmt"
	"log"
	"os/user"
	"time"
	"strconv"
	"os"
    "os/exec"
    "runtime"

	
)


func Product(a []string, r int) func() []string {
/*This the cartesian product of input iterables. Its python equivalent is
itertools.product() function(https://docs.python.org/3/library/itertools.html)*/
    p := make([]string, r)
    x := make([]int, len(p))
    return func() []string {
        p := p[:len(x)]
        for i, xi := range x {
            p[i] = a[xi]
        }
        for i := len(x) - 1; i >= 0; i-- {
            x[i]++
            if x[i] < len(a) {
                break
            }
            x[i] = 0
            if i <= 0 {
                x = x[0:0]
                break
            }
        }
        return p
    }
}

	

func main() {
    CallClear()
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		log.Println(err)
	}

	username := user.Username
	dt := time.Now()
	
	fmt.Printf("Welcome to our program user@%s. Logged in at: %s\n", username,dt.String())
	fmt.Printf("Welcome to combination program user@%s.\n"+
    	"This program takes inputs and:\n"+
    	"1. Prints all the possible combinations,\n"+
    	"2. Prints the cost for all the combinations,\n"+
    	"3. Choses the combination with the least price.\n" ,username)
    	
    	var nochargers int
    	var noev int
    	
    	//Ask for the details
    	fmt.Print("How many levels(charger types) do you have?\n")
	    fmt.Print("Enter the number of levels of chargers: ")
    	fmt.Scanln(&nochargers)
    	fmt.Printf("%d\n", nochargers)
	
	    fmt.Print("How many Electric Vehicles(EV) do you have?\n")
	    fmt.Print("Enter the number of type of EVs: ")
    	fmt.Scanln(&noev)
    	fmt.Printf("%d\n", noev)
    	
    	//create an array of charger levels
    	var chargers []string
    	for  i := 1; i <= nochargers; i++ {
    		charger:="level"+strconv.Itoa(i)
    		chargers=append(chargers,charger)
    		}
	fmt.Println(chargers)
	np := Product(chargers,  noev)
	fmt.Println("The following are all the possible combinations:\n")
	for {
        product := np()
        if len(product) == 0 {
            break
        }
        fmt.Println(product)
    }
    fmt.Println("\n To find the best combination, please enter the chargers and EV details")
    type charger struct{
    poweroutput int
    price int
    }
    type C map[string]charger

    for i:=1;i<=nochargers;i++{
        var poweroutput int
        var price int
        fmt.Println("Enter the details of charger level/type ",i)
        fmt.Print("Enter the power output of level ",i," charger in kWh: ")
        fmt.Scanln(&poweroutput)
        fmt.Print("Enter the price of level ",i," charger: ")
        fmt.Scanln(&price)
        c:=charger{poweroutput,price}
        EVs:=C{("level"+strconv.Itoa(i)):c}
        fmt.Println(EVs)
    }

    type EV struct{
    bat int
    erange int
    minrange int
    maxtime int
    count int
    }
    for i:=1;i<=noev;i++{
        var bat int
        var erange int
        var minrange int
        var maxtime int
        var count int
        fmt.Println("Enter the details of EV",i)
        fmt.Print("Enter the batter capacity of EV",i," in kWh: ")
        fmt.Scanln(&bat)
        fmt.Print("Enter the range that EV",i," cam travel in kilometres: ")
        fmt.Scanln(&erange)
        fmt.Print("Enter the minimum range EV",i," should travel in KM: ")
        fmt.Scanln(&minrange)
        fmt.Print("Enter the maxmum time that EV",i," should be allowed to charge in the station: ")
        fmt.Scanln(&maxtime)
        fmt.Print("How many EV",i," are allowed to be in the charging station at one particular time: ")
        fmt.Scanln(&count)
    }
}
var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() {
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}
