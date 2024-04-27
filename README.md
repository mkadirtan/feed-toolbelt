# Feed Toolbelt
Powerful CLI tool for working with various feed types such as RSS, Atom and JSON feeds.
TODO: Add working GIF

# Why
Feeds are awesome, but working with them is cumbersome. There are itchy edges and many standards. Feed Toolbelt makes the heavylifting for you while you get to create awesomeness.

# How

Install using go toolchain:
```bash
go install github.com/mkadirtan/feed-toolbelt@latest
```

Find feeds on a specific website:
```bash
ft find nooptoday.com
```

# Usage
```bash
ft COMMAND [OPTIONS] url
```

Available flags:
* `-filter` -   rss, atom, json

## Examples



To filter out only rss feeds:
```bash
ft find -filter=rss nooptoday.com
```

For more information run:
```bash
ft help
```

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/mkadirtan/feed-toolbelt@latest
cd feed-toolbelt
```

### Build the project

```bash
make build
```

### Run the project

```bash
./ft find nooptoday.com
```

### Run the tests

```bash
make test
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch. Make sure to add some tests, too 🚀
