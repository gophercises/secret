# Exercise #17: Secrets CLI & API

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/secret)

## Exercise details

Create a package that will store secrets - things like API keys - in an encrypted file. Then create a CLI that will make it possible to set and get these secrets via the command line as well.

*The CLI should mostly just be a wrapper around the `secret` package you create that uses a secrets file in your home directory. For more info on creating a CLI or finding the home directory on different OSes see the [task exercise](https://gophercises.com/exercises/task).*

The way developers use the final version of the `secret` package should look something like this:

```go
v := secret.FileVault("encoding-key", "path/to/file")
err := v.Set("key-name", "key-value")
value, err := v.Get("key-name")
fmt.Println(value) // "key-value"
```

From the outside the package will look fairly simple, but behind the scenes calling `v.Get` should read and decrypt the file provided and output the value for the provided key. Similarly, `v.Set` should cause your code to open up the file provided, decrypt it, set a new secret key/value pair, then save the file again in an encrypted manner.

Cryptography is easy to get wrong, so in the `p1` branch I have provided an `encrypt` package that has both the `Encrypt` and `Decrypt` functions that can be used to encrypt and decrypt strings for you. If you want, feel free to just jump to that branch and copy the code so you don't have to write this particular part of the application.

*The crypto code used in this exercise is based heavily on the code in the standard library's [CFB Encrypter](https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter) and [CFB Decrypter](https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter) examples.*

The CLI usage should probably end up looking like this:

```bash
$ secret set twitter_api_key "some value here" -k "your-encoding-key"
# Value set!
secret get twitter_api_key -k "your-encoding-key"
# "some value here"
```

You can either provide the key via a flag, or have your program read it via an environment variable. That choice is up to you.

### Readers and Writers

In the screencasts we also cover how to use the [cipher.StreamReader](https://golang.org/pkg/crypto/cipher/#StreamReader) and [cipher.StreamWriter](https://golang.org/pkg/crypto/cipher/#StreamWriter) which are incredibly cool because they allow us to just wrap any reader and writer with ciphers that will automatically encrypt and decrypt data being written to a writer or read from a reader. This means that the rest of our code doesn't have to even think about the fact that the data in our file is encrypted - it can just use a reader like it normally and with interface chaining we hide that complexity.

If you code this on your own, I definitely suggest checking out the later videos to see how this all works. Interface chaining is one of the cooler features of Go if you ask me 😀

## Bonus

Add functionality to list all keys that we have secret values stored for, and add a way to delete a key/value pair.
