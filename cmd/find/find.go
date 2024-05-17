package find

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mkadirtan/feed-toolbelt/pkg/inspect"
)

type FindCMD struct {
	Pipe           bool   `help:"use piped input" short:"p"`
	StrategyHeader bool   `help:"toggle header strategy" negatable:"" short:"l" default:"true"`
	StrategyPage   bool   `help:"toggle page strategy" negatable:"" short:"c" default:"true"`
	StrategyCommon bool   `help:"toggle common strategy" negatable:"" short:"b" default:"false"`
	Validate       bool   `help:"validate feed urls" negatable:"" short:"g" default:"false"`
	URL            string `arg:"" help:"target url"`
}

func (f *FindCMD) Run() error {
	inspectorOptions := []inspect.InspectorOption{
		inspect.WithTargetURL(f.URL),
		inspect.WithOutputHandler(func(o string) { fmt.Println(o) }),
		// inspect.WithDebugHandler(func(d string) { a.debugLogger.Println(d) }),
	}

	if f.Pipe {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			inspectorOptions = append(inspectorOptions, inspect.WithPipedInput(bufio.NewReader(os.Stdin)))
		}
	}

	if f.Validate {
		inspectorOptions = append(inspectorOptions, inspect.WithValidate())
	}

	if f.StrategyHeader {
		inspectorOptions = append(inspectorOptions, inspect.WithStrategyHeader())
	}
	if f.StrategyPage {
		inspectorOptions = append(inspectorOptions, inspect.WithStrategyPage())
	}
	if f.StrategyCommon {
		inspectorOptions = append(inspectorOptions, inspect.WithStrategyCommon())
	}

	inspectorOptions = append(inspectorOptions)

	inspector, err := inspect.NewInspector(inspectorOptions...)
	if err != nil {
		return err
	}
	inspector.Find()

	return nil
}
