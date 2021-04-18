# Password Manager v0.2
A secure and simple password manager written in go. 
It stores your credentials using a local database on the device itself with a secure AES-256 encryption that can only be accessed using the master password that you set.

![Encrytion and Decryption Flow](https://github.com/nimone/Go-Password-Manager/blob/main/encryptDecrpytFlow.png)

> Note: If you're going to use it. DO NOT run it with `go run .` since it creates the executable binary as well as the sqlite database somewhere in `\tmp`
ALWAYS USE A COMPILED VERSION OR RISK LOSING THE DATABASE FILE ON EVERY RUN.
