## Implementations
1. fmt.Scanf doesn't take space separated input, find a work around that
2. Use a SQL file to initialize the tables
3. Using the clipboard lib to securely copy and paste the password from/to to user's clipboard
4. Hide the password during entry

## Enhancements
1. Use AES 256 encryption for stored password https://github.com/gtank/cryptopasta
1. Use HMAC-SHA256 hashing for masterpassword https://github.com/gtank/cryptopasta
3. Full text search for faster sqlite queries https://www.sqlitetutorial.net/sqlite-full-text-search/