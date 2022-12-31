
# Wake My PC remote
Small helper to turn on and off my PC remotelly. 

You will need at least one machine running in the network you want to turn on the machine(a raspberry will work). This program uses wake on lan to turn on the machine when it is off, you need to enable it on your machine.

Setup your machine to run the daemon. This daemon reports current status and receives commands such as shutdown and wake on lan commands to turn on the machine you want.

As backend, this uses firebase realtime database. Create a project and download the credentials file, you will need:
- `FIREBASE_DATABASE_URL`: URL of the firebase real time database, get it from admin firebase page
- `FIREBASE_CREDENTIALS_FILE`: is the path to the firebase credentials file.

## Building
You will need golang version 1.18+ and `make`. Simply run `make build` and the binary should be generated in `bin` folder

## Dependencies
### MacOS
- wakeonlan: `brew install wakeonlan`

### Linux
- wol: `yay -Sy wol`

### Windows
- WIP

## Daemon
Make the service run in the startup of your machine. Currently it supports windows, mac and linux.
When setted up run `wakemypc daemon`

## Cli
### Daemon
`wakemypc daemon`

### List machines
List machines and states
`wakemypc ls`

### Shutdown
Shutdowns the machine if it is connected to the network
`wakemypc shutdown hostname|macaddress`

### Wol
Sends and event to all connected machines to try to wake up the machin
`wakemypc wol hostname|macaddress`
