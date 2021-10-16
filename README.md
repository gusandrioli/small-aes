# **small-aes**
Simple Encrypter/Hasher CLI application.

## Installation
### 1 - Download package (Recommended)
```sh
go install github.com/gusandrioli/small-aes
```

### 2 - Download executable
Go to Releases and download the latest binary. Then run it locally:
```sh
./small-aes
```

## Quick Start
```sh
small-aes pdfEncrypt test.pdf                  # AES encrypts pdf and returns 127 byte password
small-aes txtEncrypt test.txt yoursuperpasswrd # AES encrypts txt
```

## Commands
```
pdfDecrypt     Decrypts pdf with a 127 byte key
pdfEncrypt     Encrypts pdf with AES and a 127 byte key
sha256EvalHash Returns bool if txt file equals hash
sha256Hash     Hashes txt file into 32 byte hash
txtDecrypt     Decrypts your message with a key of 16, 24, or 32 bytes.
txtEncrypt     Encrypts your message with a key 16, 24, or 32 bytes.
```

## References
1. PDF encryption wraps the [pdfcpu API](https://github.com/pdfcpu/pdfcpu).
2. Text [encryption](https://golang.org/pkg/crypto/aes/)/[hashing](https://golang.org/pkg/crypto/sha256/) uses the standard library.

## Bugs
Bugs or suggestions? Open an issue [here](https://github.com/gusandrioli/small-aes/issues/new).
