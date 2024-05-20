# Feed Toolbelt

Powerful CLI tool for working with various feed types such as RSS, Atom and JSON feeds.

![example usage gif](https://github.com/mkadirtan/feed-toolbelt/blob/master/example.gif)

# Why

Feeds are awesome, but working with them is cumbersome. There are itchy edges and many standards. Feed Toolbelt makes
the heavylifting for you while you get to create awesomeness.

# How

Install using go toolchain:

```bash
GOEXPERIMENT=rangefunc go install github.com/mkadirtan/feed-toolbelt@latest
```

Find feeds on a specific website:

```bash
feed-toolbelt find nooptoday.com
```

# Usage

```bash
feed-toolbelt find [<url>] [flags]
```

```bash
Arguments:
[<url>]    target url, optional in case piped input is given

Flags:
-h, --help                    Show context-sensitive help.

-p, --pipe                    Use this flag if you pipe HTML content into this command. Piping without using this flag will result in interpreting piped input as
target url
-l, --[no-]strategy-header    Toggle header strategy
-c, --[no-]strategy-page      Toggle page strategy
-b, --[no-]strategy-common    Toggle common strategy
-g, --[no-]validate           Validate feed URLs contain actual feeds
```

## Examples

```bash
feed-toolbelt find nooptoday.com
```

```bash
curl https://nooptoday.com | feed-toolbelt find nooptoday.com --pipe
```

```bash
feed-toolbelt find nooptoday.com --strategy-common
```

For more information run:

```bash
feed-toolbelt --help
```

```bash
feed-toolbelt find --help
```

## 🤝 Contributing

### Clone the repo

```bash
git clone https://feed-toolbelt@latest
cd feed-toolbelt
```

### Build the project

Use `GOEXPERIMENT=rangefunc`, otherwise build fails

```bash
GOEXPERIMENT=rangefunc go build -o feed-toolbelt cmd/main.go
```

### Run the project

```bash
./feed-toolbelt find nooptoday.com
```

### Run the tests

```bash
GOEXPERIMENT=rangefunc go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch. Make sure to add
some tests, too 🚀
---

The contents below are my personal notes. If you are interested in using this tool you might be interested in the future
plans for this project.

# Roadmap

## Publish on Brew

I can publish as tap, if it gets recognition it may turn into an official package even!

Wait until rangefunc is not experimental because it will surely confuse other people and complicate the build process.

If rangefunc doesn't get merged in a short time, I can even remove the experimental part, because it is not necessary
for this project.

## Renaming

If I can't find useful commands to include in this tool, I can rename the tool to `feed-finder`.