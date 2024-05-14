package cmd

import (
	"bufio"
	"flag"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/inspect"
)

type app struct {
	debugLogger    *log.Logger
	outputLogger   *log.Logger
	positionalArgs []string
	optionArgs     []string
}

var a app

func Run() {
	a.debugLogger = log.New(os.Stderr, "", 0)
	a.outputLogger = log.New(os.Stdout, "", 0)

	usage := `Usage: feed-toolbelt <command> [arguments] 
The commands are:
	# find

	Finds feeds on a host. find requires a positional argument for hostname.
	Example usage:
		feed-toolbelt find https://nooptoday.com
	Alternatively, you can pipe http request into feed-toolbelt.
	Example usage:
		curl https://nooptoday.com | feed-toolbelt find
	
	Options:
		-strategy	list of strategies [header, link, jsonld, brute]
	Example usage:
		feed-toolbelt find https://nooptoday.com -strategy=header,link`

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			a.optionArgs = append(a.optionArgs, arg)
		} else {
			a.positionalArgs = append(a.positionalArgs, arg)
		}
	}

	if len(a.positionalArgs) < 1 {
		a.debugLogger.Println("expected a command")
		a.debugLogger.Println(usage)
		os.Exit(1)
	}

	command := a.positionalArgs[0]

	switch command {
	case "find":
		runFindCommand()
	default:
		a.debugLogger.Printf("Unsupported command: %s\n", command)
		a.debugLogger.Print(usage)
		os.Exit(1)
	}
}

func runFindCommand() {
	usage := `Usage: feed-toolbelt find targetURL [arguments]
	Example usage:
		feed-toolbelt find https://nooptoday.com
	
	Alternatively, you can pipe html output into feed-toolbelt. In this case detecting feeds from headers is not possible.
	Example usage:
		curl https://nooptoday.com | feed-toolbelt find
	
	Options:
		-strategy	list of strategies [header, link, jsonld, brute]
	Example usage:
		feed-toolbelt find https://nooptoday.com -strategy=header,link`

	f := flag.NewFlagSet("find", flag.ExitOnError)
	f.Usage = func() {
		a.debugLogger.Print(usage)
		os.Exit(1)
	}

	strategyFlag := f.String("strategy", "", "Choose which strategies to apply for finding feeds. Possible values are [header, page, common]")
	// flag.ExitOnError doesn't return errors.
	_ = f.Parse(a.optionArgs)

	validStrategies := []string{"header", "page", "common"}
	defaultStrategies := []string{"header", "page"}
	var strategies = make([]string, 0)

	if *strategyFlag == "" {
		strategies = defaultStrategies
	} else {
		for _, strategy := range strings.Split(*strategyFlag, ",") {
			strategy = strings.TrimSpace(strategy)
			if !slices.Contains(validStrategies, strategy) {
				a.debugLogger.Printf("[%s] is not recognized as a valid strategy\n", strategy)
				a.debugLogger.Println(usage)
				os.Exit(1)
			}
			strategies = append(strategies, strategy)
		}
	}

	// targetURL is defined
	if len(a.positionalArgs) >= 2 {
		targetURL := a.positionalArgs[1]
		feeds := inspect.InspectURL(
			targetURL,
			slices.Contains(strategies, "header"),
			slices.Contains(strategies, "page"),
			slices.Contains(strategies, "common"),
		)
		for _, feed := range feeds {
			a.outputLogger.Println(feed)
		}
		os.Exit(0)
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		feeds := inspect.InspectHTML(bufio.NewReader(os.Stdin))
		for _, feed := range feeds {
			a.outputLogger.Println(feed)
		}
		os.Exit(0)
	}

	a.debugLogger.Println("expected a targetURL or piped HTML input")
	a.debugLogger.Println(usage)
}
