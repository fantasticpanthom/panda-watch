# panda-watch :panda_face:

panda-watch is a configurable reconnaissance tool written in go lang for the purpose of monitoring new subdomains deployed by specific organizations and send a slack alert when it finds new subdomains. The tool is supposed to be run scheduled to run periodically.

# :scroll: Menu 

- [Usage](#Usage)
- [Installation](#Installation)

## Usage

The file containing the messages should be placed in a file with the name new.txt.

```bash
>> panda-watch -webhook <yourwebhook>
```

## Installation 
### From source:
```
>>go get -u -v github.com/fantasticpanthom/panda-watch

or

>>cd $GOPATH
>>git clone https://github.com/fantasticpanthom/panda-watch.git
>>go build main.go -o panda-watch
>>mv panda-watch /usr/local/bin
```

### From binary:
You can download the pre-built binaries from the [releases](https://github.com/fantasticpanthom/panda-watch/releases) page and then move them into your $PATH.

Binaries for other operating systems and architectures will be released soon.
