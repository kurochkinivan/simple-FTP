# Simple FTP server

## Start

```bash
./create_test_files.sh
cd cmd && go run main.go
```

Now you can connect to an ftp server on localhost:8080 with ftp client.

### This FTP server allows users to:

* **Download files** from the FTP server.
* **Navigate through directories** on the FTP server.

### Implemented Commands:

* PWD - print working directory;
* CWD - change working directory;
* LIST - list all files and directories;
* RETR — Retrieve File
* PORT — Set Active Mode Data Connection
* QUIT - close the connection
* SYST, FEAT, USER, SIZE - client information or system inquiry commands


