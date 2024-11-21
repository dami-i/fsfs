package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type flags struct {
	port  *int
	watch *bool
}

type args struct {
	dir string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Unable to acquire current working directory.")
	}

	flags := flags{
		port:  flag.Int("p", 5000, "Specify a port to which the server will listen on."),
		watch: flag.Bool("w", false, "Run the server in watch mode, listening for file changes."),
	}
	flag.Parse()

	args := args{
		dir: flag.Arg(0),
	}

	var servedDirPath = mountDirPath(cwd, args.dir)
	// Test if dir exists
	if _, err := os.Stat(servedDirPath); os.IsNotExist(err) {
		log.Fatalln("The specified directory does not exist.")
	}

	var servedPort int
	if wasFlagProvided("p") {
		servedPort = *flags.port
	} else {
		servedPort, err = findFreePort()
		if err != nil {
			log.Fatalln("Unable to acquire a free port.")
		}
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		GETOnly:               true,
	})

	app.Static("/", servedDirPath)

	app.Hooks().OnListen(func(f fiber.ListenData) error {
		println(startupMessage(f.Host, f.Port, servedDirPath))
		return nil
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", servedPort)))
}

func wasFlagProvided(name string) bool {
	var found bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func findFreePort() (int, error) {
	var port int
	for p := 5000; p < 5500; p++ {
		if isPortFree(p) {
			port = p
			break
		}
	}
	if port == 0 {
		return 0, errors.New("unable to acquire a free port")
	}
	return port, nil
}

func isPortFree(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	l.Close()
	return true
}

func mountDirPath(cwd, dir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	return filepath.Join(cwd, dir)
}

func startupMessage(host, port, dir string) string {
	var hostname = host
	if host == "0.0.0.0" {
		hostname = "localhost"
	}
	return "" +
		"┌───────────────────────────────┐\n" + //
		"│ FSFS: Fast Static File Server │\n" +
		"└───────────────────────────────┘\n" +
		"Serving files from\n" +
		"\033[1m" + dir + "\033[0m\n" +
		"over\n" +
		"\033[1mhttp://" + hostname + ":" + port + "/\033[0m\n"

}
