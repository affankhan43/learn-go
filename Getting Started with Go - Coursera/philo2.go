/* -----
This is an alternative implementation of the dining philosophers problem, where
eating is controlled by a host. A philosopher is allowed to eat when the following
criteria are fulfilled:
// * The maximum number of parallelly eating philosophers has not yet been reached.
// * The philosopher is not already eating, do not exceed the number of allowed dinners and
//   there's enough time elapsed since the philosophers last dinner.
// * Both chopsticks corresponding to the philosopher's seat are free.
// Note: seats are randomized.
----- */

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup

	names := []string{"John", "Robert", "Lilla", "Charles", "Stella"}
	tableCount := len(names)

	host := NewDinnerHostPtr(tableCount, 2, 3)
	requestChannel, finishChannel := host.AskChannels()
	go host.Listen()

	for _, name := range names {
		phi := NewPhilosopherPtr(name)
		accepted := host.Add(*phi)
		if accepted {
			waitGrp.Add(1)
			go phi.GoToDinner(&waitGrp, requestChannel, finishChannel)
		}
	}

	waitGrp.Wait()

	fmt.Println("===== EVERYBODY LEFT THE TABLE. =====")
	fmt.Println()

	time.Sleep(time.Duration(2) * time.Second)
}

/* ---
The host of the dinner.
--- */

// DinnerHost is the main data structure for the host of the dinner.
type DinnerHost struct {
	phiData         map[string]*philosopherData
	requestChannel  chan string
	finishChannel   chan string
	maxParallel     int
	maxDinner       int
	currentlyEating int
	tableCount      int
	chopsticksFree  []bool
	freeSeats       []int
}

// philosopherData contains philosopher specific data. It is used within DinnerHost.
type philosopherData struct {
	respChannel    chan string
	eating         bool
	dinnersSpent   int
	seat           int
	leftChopstick  int
	rightChopstick int
	finishedAt     time.Time
}

// NewDinnerHostPtr creates a new, initialized DinnerHost object and returns a pointer to it.
func NewDinnerHostPtr(tableCount, maxParallel, maxDinner int) *DinnerHost {
	host := new(DinnerHost)
	host.Init(tableCount, maxParallel, maxDinner)
	return host
}

// Init is used to initialize the DinnerHost. Note: seats are randomized.
func (host *DinnerHost) Init(tableCount, maxParallel, maxDinner int) {
	host.phiData = make(map[string]*philosopherData)
	host.requestChannel = make(chan string)
	host.finishChannel = make(chan string)
	host.maxParallel = maxParallel
	if host.maxParallel > tableCount {
		host.maxParallel = tableCount
	}
	host.maxDinner = maxDinner
	host.currentlyEating = 0
	host.tableCount = tableCount
	host.chopsticksFree = make([]bool, 5)
	for i := range host.chopsticksFree {
		host.chopsticksFree[i] = true
	}
	rand.Seed(time.Now().Unix())
	host.freeSeats = rand.Perm(tableCount)
}

// newPhilosopherDataPtr creates and initializes a philosopherData object and returns a pointer to it.
func newPhilosopherDataPtr(respChannel chan string) *philosopherData {
	pd := new(philosopherData)
	pd.Init(respChannel)
	return pd
}

// Init is used to initialize the philosopherData.
func (pd *philosopherData) Init(respChannel chan string) {
	pd.respChannel = respChannel
	pd.eating = false
	pd.dinnersSpent = 0
	pd.seat = -1
	pd.leftChopstick = -1
	pd.rightChopstick = -1
}

// ===== DinnerHost methods =====

// AskChannels can be used to obtain two common channels of the host, the first used to request dinner,
// the second used to indicate that someone finished eating.
func (host *DinnerHost) AskChannels() (chan string, chan string) {
	return host.requestChannel, host.finishChannel
}

// Add registers the philosopher at the host. It first checks if they can join (table full, already at
// the table), then creates a new philosopher data record and assigns a seat to the
func (host *DinnerHost) Add(newPhilosopher Philosopher) bool {
	newName := newPhilosopher.Name()
	fmt.Println(newName + " WANTS TO JOIN THE TABLE.")
	if len(host.phiData) >= host.tableCount {
		fmt.Println(newName + " CANNOT JOIN: THE TABLE IS FULL.")
		fmt.Println()
		return false
	}
	if host.phiData[newName] != nil {
		fmt.Println(newName + " CANNOT JOIN: ALREADY ON THE HOST'S LIST.")
		fmt.Println()
		return false
	}
	host.phiData[newName] = newPhilosopherDataPtr(newPhilosopher.RespChannel())
	host.phiData[newName].TakeSeat(host.freeSeats[0])
	host.freeSeats = host.freeSeats[1:]
	fmt.Println(newName + " JOINED THE TABLE.")
	fmt.Println()
	return true
}

// Listen is the main function of the host, which handles dinner requests and finish
// indications coming from the philosophers on _requestChannel_ and _finishChannel_.
// Dinner request is authorized with a proper reply to a philosopher on its own
// dedicated response channel.
func (host *DinnerHost) Listen() {
	name := ""
	for {
		select {
		case name = <-host.requestChannel:
			fmt.Println(name + " WOULD LIKE TO EAT.")

			response := host.AllowEating(name)
			kickOut := false
			switch response {
			case "OK":
				fmt.Println(name + " STARTS EATING.")
			case "E:CHOPSTICKS":
				fmt.Println(name + " CANNOT EAT: REQUIRED CHOPSTICKS ARE NOT AVAILABLE.")
			case "E:FULL":
				fmt.Println(name + " CANNOT EAT: TWO OTHER PHILOSOPHERS ARE ALREADY EATING.")
			case "E:JUSTFINISHED":
				fmt.Println(name + " CANNOT EAT: JUST FINISHED THE PREVIOUS MEAL.")
			case "E:EATING":
				fmt.Println(name + " CANNOT EAT: ALREADY EATING.")
			case "E:LIMIT":
				fmt.Println(name + " CANNOT EAT: ALREADY HAD THREE DINNERS; MUST LEAVE.")
				host.freeSeats = append(host.freeSeats, host.phiData[name].Seat())
				kickOut = true
			}
			fmt.Println()

			host.phiData[name].RespChannel() <- response

			if kickOut {
				delete(host.phiData, name)
			}
		case name = <-host.finishChannel:
			host.SomeoneFinished(name)
		}
		host.PrintReport(false)
	}
}

// AllowEating checks if the philosopher is allowed to have dinner. Criteria:
// * No more than _maxParallel_ philosophers can eat in parallel.
// * The philosopher is not already eating, do not exceed the number of allowed dinners and
//   there's enough time elapsed since the philosopher's last dinner.
// * Both chopsticks corresponding to the philosopher's seat are free.
// The function also takes care of chopstick reservation. Note: when only either of the
// chopsticks is free, it is reserved in spite the philosopher cannot start eating.
func (host *DinnerHost) AllowEating(name string) string {
	if host.currentlyEating >= host.maxParallel {
		return "E:FULL"
	}

	data := host.phiData[name]

	canEat := data.CanEat(host.maxDinner)
	if canEat != "OK" {
		return canEat
	}

	seatNumber := data.Seat()
	leftChop := seatNumber
	rightChop := (seatNumber + 1) % host.tableCount

	if host.chopsticksFree[leftChop] {
		host.chopsticksFree[leftChop] = false
		data.SetLeftChop(leftChop)
	}
	if host.chopsticksFree[rightChop] {
		host.chopsticksFree[rightChop] = false
		data.SetRightChop(rightChop)
	}

	if !data.HasBothChopsticks() {
		return "E:CHOPSTICKS"
	}

	host.currentlyEating++
	data.StartedEating()

	return "OK"
}

// SomeoneFinished takes the necessary actions when a philosopher finished eating.
func (host *DinnerHost) SomeoneFinished(name string) {
	if host.currentlyEating > 0 {
		host.currentlyEating--
	}
	host.chopsticksFree[host.phiData[name].LeftChopstick()] = true
	host.chopsticksFree[host.phiData[name].RightChopstick()] = true
	host.phiData[name].FinishedEating()
	fmt.Println(name + " FINISHED EATING.")
	fmt.Println()
}

// PrintReport shows the status of the philosophers in a verbose format.
func (host *DinnerHost) PrintReport(additionalInfo bool) {
	names := make([]string, 0, len(host.phiData))
	maxNameLen := 0
	for i := range host.phiData {
		names = append(names, i)
		if len(i) > maxNameLen {
			maxNameLen = len(i)
		}
	}

	sort.Strings(names)

	fmt.Printf("%*s | SEAT | LEFTCH. | RIGHTCH. | DINNERS | STATUS", maxNameLen, "NAME")
	fmt.Println()

	for _, name := range names {
		data := host.phiData[name]
		status := "waiting"
		if data.eating == true {
			status = "eating"
		}

		leftChopStr := strings.Replace(strconv.Itoa(data.LeftChopstick()), "-1", "X", 1)
		rightChopStr := strings.Replace(strconv.Itoa(data.RightChopstick()), "-1", "X", 1)

		repLine := fmt.Sprintf("%*s | %*d | %*s | %*s | %*d | %s",
			maxNameLen, name, 4, data.seat, 7, leftChopStr,
			8, rightChopStr, 7, data.dinnersSpent, status)
		fmt.Println(repLine)
	}

	if additionalInfo {
		freeChops := fmt.Sprintf("CHOPSTICKS:")
		for chopInd, chopStat := range host.chopsticksFree {
			status := "FREE"
			if chopStat == false {
				status = "RESERVED"
			}
			freeChops += fmt.Sprintf(" %d[%s]", chopInd, status)
		}
		fmt.Println(freeChops)
	}
	fmt.Println()
}

// ===== philosopherData methods

// CanEat checks if the philosopher specific criteria of eating is fulfilled.
func (pd *philosopherData) CanEat(maxDinner int) string {
	switch {
	case pd.eating:
		return "E:EATING"
	case pd.dinnersSpent >= maxDinner:
		return "E:LIMIT"
	case time.Now().Sub(pd.finishedAt) < (time.Duration(150) * time.Millisecond):
		return "E:JUSTFINISHED"
	}
	return "OK"
}

// StartedEating updates philosopher specific data when the philosopher starts eating.
func (pd *philosopherData) StartedEating() {
	pd.eating = true
	pd.dinnersSpent++
}

// FinishedEating updates philosopher specific data when the philosopher finished eating.
func (pd *philosopherData) FinishedEating() {
	pd.eating = false
	pd.leftChopstick = -1
	pd.rightChopstick = -1
	pd.finishedAt = time.Now()
}

// RespChannel returns the philosopher's response channel.
func (pd *philosopherData) RespChannel() chan string {
	return pd.respChannel
}

// LeftChopstick returns the ID of the philosopher's currently reserved left chopstick.
// If no left chopstick is reserved, then -1 is returned.
func (pd *philosopherData) LeftChopstick() int {
	return pd.leftChopstick
}

// RightChopstick returns the ID of the philosopher's currently reserved right chopstick.
// If no right chopstick is reserved, then -1 is returned.
func (pd *philosopherData) RightChopstick() int {
	return pd.rightChopstick
}

// HasBothChopsticks returns true if both of the chopstics are reserved for the philosopher.
func (pd *philosopherData) HasBothChopsticks() bool {
	return (pd.leftChopstick >= 0) && (pd.rightChopstick >= 0)
}

// SetLeftChop can be used to set the left chopstick ID for the philosopher.
func (pd *philosopherData) SetLeftChop(leftChop int) {
	pd.leftChopstick = leftChop
}

// SetRightChop can be used to set the right chopstick ID for the philosopher.
func (pd *philosopherData) SetRightChop(rightChop int) {
	pd.rightChopstick = rightChop
}

// TakeSeat can be used to set the seat number for the philosopher.
func (pd *philosopherData) TakeSeat(seatNumber int) {
	pd.seat = seatNumber
}

// Seat returns the seat number of the philosopher.
func (pd *philosopherData) Seat() int {
	return pd.seat
}

/* ---
The philosophers.
--- */

// Philosopher represents an individual philosopher.
type Philosopher struct {
	name        string
	respChannel chan string
}

// NewPhilosopherPtr creates and initializes a Philosopher object and returns a pointer to it.
func NewPhilosopherPtr(name string) *Philosopher {
	phi := new(Philosopher)
	phi.Init(name)
	return phi
}

// Init can be used to initialize a Philosopher.
func (phi *Philosopher) Init(name string) {
	phi.name = name
	phi.respChannel = make(chan string)
}

// Name returns the name of the philosopher.
func (phi *Philosopher) Name() string {
	return phi.name
}

// RespChannel returns the philosopher's dedicated response channel.
func (phi *Philosopher) RespChannel() chan string {
	return phi.respChannel
}

// GoToDinner is the philosopher's main task. They periodically issue eat requests to the host, unless
// not already eating. When asked so by the host, the philosopher leaves.
func (phi *Philosopher) GoToDinner(waitGrp *sync.WaitGroup, requestChannel, finishChannel chan string) {
	defer waitGrp.Done()

	retryInterval := time.Duration(2000) * time.Millisecond
	eatingDuration := time.Duration(5000) * time.Millisecond

	for {
		requestChannel <- phi.name
		switch <-phi.respChannel {
		case "OK":
			time.Sleep(eatingDuration)
			finishChannel <- phi.name
		case "E:LIMIT":
			fmt.Println(strings.ToUpper("----- " + phi.name + " LEFT THE TABLE. -----"))
			fmt.Println()
			return
		default:
			time.Sleep(retryInterval)
		}
	}
}