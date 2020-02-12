
deebot
======
A Golang program to control your Ecovacs vacuum from Home Automation systems such as Home Assisitant, HomeKit (HomeBridge), Alexa. 


## Usage

To get started, you'll need to have already set up an EcoVacs account
using your smartphone. Update deebot.json with your peticulars and place it in /usr/local/etc/

Determine your device ID using the

```
pip install sucks
pip install pipenv
```

CH: msg.ecouser.net
TW, MY, JP, SG, TH, HK, IN, KR: msg-as.ecouser.net
US: msg-na.ecouser.net
FR, ES, UK, NO, MX, DE, PT, CH, AU, IT, NL, SE, BE, DK: msg-eu.ecouser.net
Any other country: msg-ww.ecouser.net




With that set up, you could have it clean in auto mode:

```
    % deebot clean-auto
```

To tell it to go charge:

```
    % deebot charge
```

Check it's status- useful for Home Automation applications

```
    % deebot status
```


Example Homebridge cmdSwitch2 configuration
```
"platforms": [{
        "platform": "cmdSwitch2",
        "name": "CMD Switch",
        "switches": [{
                "name" : "Vacuum",
                "on_cmd": "deebot clean-auto",
                "off_cmd": "deebot charge",
                "state_cmd": "deebot status 2>&1 | grep RUNNING",
                "type": "Switch"
        }]
}]
```

## Cross compile for Raspberry PI:
```
env GOOS=linux GOARCH=arm GOARM=7 go build -o deebot-arm7
```

### Thanks

Special thanks to:

* [skburgart](https://github.com/skburgart/go-vacbot) who developed the Go based API calls.
* All the users who have given useful feedback and contributed code!