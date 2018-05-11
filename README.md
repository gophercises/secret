# Exercise #17: Secrets CLI & API

[![exercise status: in progress](https://img.shields.io/badge/exercise%20status-in%20progress-yellow.svg?style=for-the-badge)](https://gophercises.com/exercises/secret)

## Exercise details

<!-- TODO(jon): Finish this -->

```
secret set twitter_api_key "some-value"
secret get twitter_api_key # some-value
```

```go
flag.String(encryptionKey)
var c secret.Client{
  EncryptionKey: "some-key"
}
c.Get("twitter_api_key") # returns "some-value"
```

```
./app -encryption_key="laskdfjasldkfj"
```
