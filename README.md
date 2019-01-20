
deebot
======
A Golang program to control your Ecovacs vacuum from Home Automation systems such as Home Assisitant, HomeKit (HomeBridge), Alexa. 


## Usage

To get started, you'll need to have already set up an EcoVacs account
using your smartphone. Update deebot.json with your peticulars and place it in /usr/local/etc/


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


### Thanks

Special thanks to:

* [skburgart](https://github.com/skburgart/go-vacbot) who developed the Go based API calls.
* All the users who have given useful feedback and contributed code!