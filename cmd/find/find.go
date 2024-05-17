package find

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mkadirtan/feed-toolbelt/pkg/inspect"
)

type FindCMD struct {
	Pipe           bool   `help:"Use this flag if you pipe HTML content into this command. Piping without using this flag will result in interpreting piped input as target url" short:"p"`
	StrategyHeader bool   `help:"Toggle header strategy" negatable:"" short:"l" default:"true"`
	StrategyPage   bool   `help:"Toggle page strategy" negatable:"" short:"c" default:"true"`
	StrategyCommon bool   `help:"Toggle common strategy" negatable:"" short:"b" default:"false"`
	Validate       bool   `help:"Validate feed URLs contain actual feeds" negatable:"" short:"g" default:"false"`
	URL            string `arg:"" help:"target url, optional in case piped input is given" optional:""`
}

func (f *FindCMD) Run() error {
	options := []inspect.Option{
		inspect.WithOutputHandler(func(o string) { fmt.Println(o) }),
	}

	// not implemented yet
	if false {
		options = append(options, inspect.WithDebugHandler(func(d string) { fmt.Println(d) }))
	}

	if f.Pipe {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			options = append(options, inspect.WithPipedInput(bufio.NewReader(os.Stdin)))
		} else {
			return errors.New("pipe option used without piped input")
		}
	}

	// no positional url, pipe option true OK
	// no positional url, pipe option false, url from pipe OK
	// no positional url, pipe option false, no url from pipe NO
	if f.URL == "" {
		if !f.Pipe {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				urlBytes, err := io.ReadAll(os.Stdin)
				if err != nil {
					return err
				}
				targetURL := strings.TrimSpace(string(urlBytes))
				options = append(options, inspect.WithTargetURL(targetURL))
			} else {
				return errors.New("no url specified")
			}
		}
	} else {
		options = append(options, inspect.WithTargetURL(f.URL))
	}

	if f.Validate {
		options = append(options, inspect.WithValidate())
	}

	if f.StrategyHeader {
		options = append(options, inspect.WithStrategyHeader())
	}
	if f.StrategyPage {
		options = append(options, inspect.WithStrategyPage())
	}
	if f.StrategyCommon {
		options = append(options, inspect.WithStrategyCommon())
	}

	options = append(options)

	inspector, err := inspect.NewInspector(options...)
	if err != nil {
		return err
	}

	return inspector.Find()
}
