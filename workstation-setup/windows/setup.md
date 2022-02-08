# Windows 10 / 11

If you running a modern windows you can install and setup a fully working WSL2 and Docker environment.
Most of your dev tools will run natively in a Linux VM but some tools need to be installed on windows.
The scoop package manager is quite convinient for that.

## Scoop

Details can be found here: [scoop.sh](https://scoop.sh)

Quick setup and install common tools defined in the [setup script](scoop-setup.cmd).
Open a powershell:
```
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
iwr -useb get.scoop.sh | iex
.\scoop-setup.cmd
```