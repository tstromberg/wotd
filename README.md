# wotd

Word of the Day (Go Library/CLI)

## Examples

### CLI

The `wotd` command-line tool pulls the most recent word-of-the-day from Wordnik and Dictionary.com - appropriate for a login shell.

```
rareripe —  https://wordnik.com/word-of-the-day
    adjective: Ripening early.
    noun     : A fruit or vegetable that ripens early.

diasporic —  https://www.dictionary.com/e/word-of-the-day/
    adjective: of, being, or relating to any group that has been dispersed outside its traditional homeland, either involuntarily or by migration.
```

To install, run `go install github.com/tstromberg/cmd/wotd/wotd.go`

### Library

See `cmd/wotd/motd.go` as an example, but the API is simple:

```go
r, err := wotd.Wordnik()
if err != nil {
    panic(err.Error())
}
fmt.Printf("The word of the day is: %s", r.Word)
```
