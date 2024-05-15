# Feed Toolbelt

Powerful CLI tool for working with various feed types such as RSS, Atom and JSON feeds.
TODO: Add working GIF

# Why

Feeds are awesome, but working with them is cumbersome. There are itchy edges and many standards. Feed Toolbelt makes
the heavylifting for you while you get to create awesomeness.

# How

Install using go toolchain:

```bash
go install github.com/mkadirtan/feed-toolbelt@latest
```

Find feeds on a specific website:

```bash
feed-toolbelt find nooptoday.com
```

# Usage

```bash
feed-toolbelt COMMAND [OPTIONS] url
```

Available flags:

* `-filter` - rss, atom, json

## Examples

To filter out only rss feeds:

```bash
feed-toolbelt find -filter=rss nooptoday.com
```

For more information run:

```bash
feed-toolbelt help
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
go build -o feed-toolbelt cmd/main.go
```

### Run the project

```bash
./feed-toolbelt find nooptoday.com
```

### Run the tests

```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch. Make sure to add
some tests, too 🚀
---

The contents below are my personal notes. If you are interested in using this tool you might be interested in the future
plans for this project.

# Plan - What to do

Future plans for this project.

## Accept all url formats

- `feed-toolbelt find https://nooptoday.com`
- `feed-toolbelt find  http://nooptoday.com`
- `feed-toolbelt find nooptoday.com`

all should be treated equally.

## Check target url

Some crazy users might input a feed url as target url. In those cases, the program should report the url back.

## Take validation flag

- `feed-toolbelt find https://nooptoday.com -validate`

`-validate` option should validate found feed urls. Common urls are an exception to this rule, they are always
validated. Those urls are already requested and in many cases those urls result in http status 200 with non-feed
content.

## Debug Output

- `feed-toolbelt find https://nooptoday.com -v`
  `-v`, `-verbose` Should report certain issues from the debug logger, if the flag is set.

There is no need to implement a verbosity level in a near future.

Some important output can be:

- Request errors such as 404, 429 or 500.
- URL format errors, in case the requested url is invalid.

## Structured Output

- `feed-toolbelt find https://nooptoday.com -s`
  `-s`, `-structured` Should report the found urls in a structured json format.
  The format can include:
- url
- detected_in ( first found place: header, link, common, a tag etc. )
- feed_type ( rss, json etc. )

## Concurrent request limit and Rate limit

- `feed-toolbelt find https://nooptoday.com -c 4`
  `-c` or `-concurrent` option should define the concurrent request limit.

Also, if any of the requests get the `429` rate limit error, this should be reported.

## Unique feed urls

Simply keep an Inspector state and don't report the same feed url twice, or do not check the previously found urls in
the common paths strategy.

## Report redirected url or the target url

- `feed-toolbelt find https://nooptoday.com --report-redirected`

This feature will rarely be useful, however it is a nice to have for completeness.

## Publish on Brew

I can publish as tap, if it gets recognition it may turn into an official package even!

Wait until rangefunc is not experimental because it will surely confuse other people and complicate the build process.

If rangefunc doesn't get merged in a short time, I can even remove the experimental part, because it is not necessary
for this project.

## Renaming

If I can't find useful commands to include in this tool, I can rename the tool to `feed-finder`.

# Anti Plan - What not to do

These are some features that will not be implemented in this tool, because they are either out of scope or not preferred
for simplicity.

## Custom Headers

Do not accept custom headers from the user, it is not the responsibility of this program.
If required, user can use curl or other tools to specialize the request.

## Parsed Feed Output

Do not output the parsed feed data. This is out of scope.

## Do not visit links outside the specified domain

An arbitrary rule, sometimes you need to draw boundaries. I have no good argument for or against.