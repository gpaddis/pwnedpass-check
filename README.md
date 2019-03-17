Build the package and check for breached passwords:

```
$ go build
$ ./pwnedpass-check pass123
The password pass123 was pwned and appears 82170 times in the database.
```