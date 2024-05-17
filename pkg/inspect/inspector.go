package inspect

import (
	"errors"
	"io"
)

type HandlerFunc func(string)

type Inspector struct {
	config      InspectorConfig
	foundFeeds  []string
	visitedURLs []string
}

type InspectorConfig struct {
	// strategies option is ignored in case TargetHTML is defined
	// TargetHTML takes precedence over targetURL
	TargetHTML *io.Reader

	TargetURL  *string
	Strategies struct {
		// HTTP Link Header with feed attributes
		Header bool
		// a and link tags with feed attributes
		Page bool
		// try and fetch common feed urls, if they contain valid feeds they will be reported
		Common bool
	}
	// fetch and Validate reported feed links, common strategy validates the feed content regardless of this option
	Validate bool
	// found feeds will be reported to this function
	OutputHandler HandlerFunc
	// debug logs will be reported to this function
	DebugHandler HandlerFunc
}

type InspectorOption func(*InspectorConfig)

func WithTargetHTML(htmlBody io.Reader) InspectorOption {
	return func(c *InspectorConfig) {
		c.TargetHTML = &htmlBody
	}
}

func WithTargetURL(targetURL string) InspectorOption {
	return func(c *InspectorConfig) {
		c.TargetURL = &targetURL
	}
}

func WithStrategyHeader() InspectorOption {
	return func(c *InspectorConfig) {
		c.Strategies.Header = true
	}
}

func WithStrategyPage() InspectorOption {
	return func(c *InspectorConfig) {
		c.Strategies.Page = true
	}
}

func WithStrategyCommon() InspectorOption {
	return func(c *InspectorConfig) {
		c.Strategies.Common = true
	}
}

func WithValidate() InspectorOption {
	return func(c *InspectorConfig) {
		c.Validate = true
	}
}

func WithOutputHandler(outputHandler HandlerFunc) InspectorOption {
	return func(c *InspectorConfig) {
		c.OutputHandler = outputHandler
	}
}

func WithDebugHandler(debugHandler HandlerFunc) InspectorOption {
	return func(c *InspectorConfig) {
		c.DebugHandler = debugHandler
	}
}

var (
	errNoTarget        = errors.New("no target specified")
	errNoOutputHandler = errors.New("no output handler specified")
)

func NewInspector(options ...InspectorOption) (*Inspector, error) {
	config := &InspectorConfig{}

	for _, option := range options {
		option(config)
	}

	if config.TargetHTML == nil && config.TargetURL == nil {
		return nil, errNoTarget
	}

	if config.OutputHandler == nil {
		return nil, errNoOutputHandler
	}

	return &Inspector{
		config:      *config,
		foundFeeds:  make([]string, 0),
		visitedURLs: make([]string, 0),
	}, nil
}
