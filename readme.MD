
Small helper to turn on and off my PC remotelly. 

You will need at least one machine running in the network you want to turn on the machine. This program uses wake on lan to turn on the machine when it is off, you need to enable it on your machine.

Setup your machine to run the daemon. This daemon reports current status and receives commands such as shutdown and wake on lan commands to turn on the machine you want.

As backend, this uses firebae realtime database. Create a project and download the credentials file, you will need:
- FIREBASE_DATABASE_URL: URL of the firebase real time database, get it from admin firebase page
- FIREBASE_CREDENTIALS_FILE(default: `firebase_credentials.json`): is the path to the firebase credentials file.

## Dependencies
### MacOS
- wakeonlan: `brew install wakeonlan`

### Linux
- wol: `yay -Sy wol`
