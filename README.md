# fav
Favarite sites store for CLI
Manage your favorite site information.
Since it can be opened in the specified browser, it can be opened without sharing session information.

## How to install
1. Download the source code from Github
```
git clone git@github.com:haytty/fav.git
```
2. Run make install.
Deployed in `/usr/local/bin` by default
```
sudo make install
```

## Usage
Instructions on how to use this tool.
Check the help for detailed command usage. `fav --help`

1. First create a config file
```
fav init
```

2. Next, register your favorite site information
```
fav add Github https://github.com
```

3. Next, register your favorite browser information
```
fav browser add Chrome "/Applications/Google Chrome.app"
```

4. Start `fav`. You can use your favorite information.
```
fav
```
