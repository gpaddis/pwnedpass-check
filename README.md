Build the package and check for breached passwords:

```sh
$ go build
$ ./pwnedpass-check pass123
The password pass123 was found in 82170 breaches.
```