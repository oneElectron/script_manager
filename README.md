# Script manager
A little program I wrote to manage one-off scripts

# Quick Start
```sh
go install github.com/oneElectron/script_manager/cmd/s@latest
go install github.com/oneElectron/script_manager/cmd/sm@latest
```

Then you can edit a script with your favorite editor (defined by the $EDITOR variable) with:
```sh
# export EDITOR=nvim
sm edit <SCRIPT NAME>
```

You can run the script with:
```sh
s <SCRIPT NAME>
```

the ```s``` binary is only to run scripts.
the ```sm``` binary is to do everything that isnt running scripts.
To get a full list of options you can run:
```sh
sm --help
```

Have fun :)
