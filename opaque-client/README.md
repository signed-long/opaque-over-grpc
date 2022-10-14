### opaque-over-grpc cli design

#### Commands:

1. `add <name> [-skipCheckHash]`

- Adds a password to the db with a corresponing name
- checks prefix of pw hash against [HIBP API](https://haveibeenpwned.com/API/v3)

2. `copy <name>`

- places the passowrd stored with given name in the user's clipboard
- waits for user input to clear clipboard

3. `print <name>`

- prints out the user's password

4. `delete <name>`

- deletes the record from the db with name

5. `gen <name>`

- Generates, and if name is provided, adds a password to the db with a corresponing name
- Generates a strong password
