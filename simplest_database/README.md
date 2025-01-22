This is the "world's simplest database" as described
in the "Designing Data-intensive Applications" book.
See page 122 in the book.

# Usage

Run the module from the `simplest_database` folder

```
$ go run .
Please provide arguments.

$ go run . set name John
Setting name to John
Successfully set name to John

$ go run . set name Jane
Setting name to Jane
Successfully set name to Jane

$ go run . get name
Value for name is Jane
```

# Testing

Run the tests from the `simplest_database` folder

```
$ go test
```
