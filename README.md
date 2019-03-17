Build the package and check for breached passwords:

```sh
$ go build
$ ./pwned-api pass123
The password pass123 was found in 82170 breaches.
```