package inspect

import (
	"bufio"
	"errors"
	"net/http"

	"github.com/mkadirtan/feed-toolbelt/pkg/util"
)

type HandlerFunc func(string)

type Inspector struct {
	config      InspectorConfig
	foundFeeds  []string
	visitedURLs []string
	body        *bufio.Reader
	header      http.Header
}

type InspectorConfig struct {
	// strategies option is ignored in case PipedInput is defined
	PipedInput *bufio.Reader

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
}

type Option func(*InspectorConfig)

func WithPipedInput(htmlBody *bufio.Reader) Option {
	return func(c *InspectorConfig) {
		c.PipedInput = htmlBody
	}
}

func WithTargetURL(targetURL string) Option {
	return func(c *InspectorConfig) {
		c.TargetURL = &targetURL
	}
}

func WithStrategyHeader() Option {
	return func(c *InspectorConfig) {
		c.Strategies.Header = true
	}
}

func WithStrategyPage() Option {
	return func(c *InspectorConfig) {
		c.Strategies.Page = true
	}
}

func WithStrategyCommon() Option {
	return func(c *InspectorConfig) {
		c.Strategies.Common = true
	}
}

func WithValidate() Option {
	return func(c *InspectorConfig) {
		c.Validate = true
	}
}

func WithOutputHandler(outputHandler HandlerFunc) Option {
	return func(c *InspectorConfig) {
		c.OutputHandler = outputHandler
	}
}

var (
	errNoTarget        = errors.New("no target specified")
	errNoOutputHandler = errors.New("no output handler specified")
)

func NewInspector(options ...Option) (*Inspector, error) {
	config := &InspectorConfig{}

	for _, option := range options {
		option(config)
	}

	if config.PipedInput == nil && config.TargetURL == nil {
		return nil, errNoTarget
	}

	if config.OutputHandler == nil {
		return nil, errNoOutputHandler
	}

	if config.TargetURL != nil {
		normalizedTargetURL, err := util.NormalizeURL(*config.TargetURL)
		if err != nil {
			return nil, err
		}
		*config.TargetURL = normalizedTargetURL
	}

	return &Inspector{
		config:      *config,
		foundFeeds:  make([]string, 0),
		visitedURLs: make([]string, 0),
	}, nil
}
