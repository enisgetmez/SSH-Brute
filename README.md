
# SSH Brute Force Tool

## Preparing the Data

 - Add your password wordlist to the file lists/passwords.txt.
 - Add potential user names to the file lists/users.txt.
 - Add a list of servers with open SSH ports to lists/servers.txt.

## Running
Run the software using the command:
    go run sshbrute.go
Successfully logged in SSH servers will be saved to the results.txt folder.
