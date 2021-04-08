# Password Manager
A secure and simple password manager written in go. 
It stores your credentials using a local database on the device itself with a secure RSA encryption that can only be accessed using the master password that you set.

> Note: If you're going to use it. DO NOT run it with `go run .` since it creates the executable binary as well as the sqlite database somewhere in `\tmp`
ALWAYS USE A COMPILED VERSION OR RISK LOSING THE DATABASE FILE ON EVERY RUN.