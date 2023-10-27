package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

func ReadMemoryStats() Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)
	res := Memory{}
	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value
		}
	}
	return res
}

func parseLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	reset := "\033[0m"
	gopherColor := "\033[36m"
	cCyan := "\033[34m"

	Hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	username := currentUser.Username

	dashLength := len(username) + len(Hostname) + 1

	// distroname
	cmd := exec.Command("lsb_release", "-d")

	distro, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	distroString := string(distro)
	CollonIndex := strings.Index(distroString, ":")
	distroName := strings.TrimSpace(distroString[CollonIndex+1:])

	// RAM usage

	mem := ReadMemoryStats()

	totalMem := mem.MemTotal / 1024

	usedMem := (mem.MemTotal - mem.MemAvailable) / 1024

	// Uptime
	cmd = exec.Command("uptime", "-p")

	uptime, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	uptimeString := string(uptime)
	uptimeValue := uptimeString[2:]

	// terminal color designs
	const (
		bold    = "\033[1m"
		inverse = "\033[7m"
	)

	colors := []string{
		"\033[30m", // Black
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // White
	}

	fmt.Println()
	fmt.Print(reset)

	fmt.Println(gopherColor + "         ,_---~~~~~----._         ")
	fmt.Println("  _,,_,*^____      _____``*g*\"*, \t" + cCyan + username + reset + "@" + cCyan + Hostname + gopherColor)
	fmt.Println(" / __/ /'     ^.  /      \\ ^@q   f \t" + reset + strings.Repeat("-", dashLength) + gopherColor)
	fmt.Println("[  @f | @))    |  | @))   l  0 _/  \t" + cCyan + "OS: \t" + reset + distroName + gopherColor)
	fmt.Println(" `\\   \\~____ / __ \\_____/    \\   \t" + cCyan+ "RAM: \t" + reset + strconv.Itoa(usedMem) + " MiB / " + strconv.Itoa(totalMem) + " MiB" + gopherColor)
	fmt.Println("  |           _l__l_           I   \t" + cCyan+ "Cores: \t" + reset + strconv.Itoa(runtime.NumCPU()) + gopherColor)
	fmt.Println("  }          [______]           I  \t" + cCyan+ "Uptime:\t" + reset + strings.TrimSpace(uptimeValue) + gopherColor)
	fmt.Println("  ]          " + reset + "  | | | " + gopherColor + "          |  \t")
	fmt.Print("  ]            " + reset + " ~ ~  " + gopherColor + "           |  \t")
	fmt.Printf("%s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬\n", colors[1], colors[2], colors[3], colors[4], colors[5], colors[6])
	fmt.Print(gopherColor + "  |                            |   \t")
	fmt.Printf("%s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬ %s▬▬▬▬▬\n", bold+colors[1], bold+colors[2], bold+colors[3], bold+colors[4], bold+colors[5], bold+colors[6])
	fmt.Println(gopherColor + "   |                           |   " + reset)

}
