# slug

  slug generator

## Installation

    $ go get github.com/travisjeffery/slug

## API

#### slug.Slugify(str, replacement, seperator string)

```go
Slugify("foo bar", "", "")
// > foo-bar
Slugify("foo bar", "", "_")
// > foo_bar
Slugify("foo bar", "o", "")
// > f-bar
```

## Tests

    $ go test

## License

  MIT
