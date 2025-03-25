# Script manager
A little program I wrote to manage one-off scripts

# Quick Start
```sh
go install github.com/oneElectron/script_manager/cmd/sm@latest
```

Then you can edit a script with your favorite editor (defined by the $EDITOR variable) with:
```sh
# export EDITOR=nvim
sm --edit <SCRIPT NAME>
```

You can run the script with:
```sh
sm <SCRIPT NAME>
```
Have fun :)
