## About

Filepicker CLI upload tool enables you to easily upload your files and get uniqe url (FileLink) for each on of them. Url is also automatically copied to your clipboard so you can start using it immediately.

## Installation

### From binary

Just download one of the binaries:

[Darwin x86_64](https://github.com/Ink/fp/blob/master/bin/darwin_x86_64/fp?raw=true)

[Linux x86_64](https://github.com/Ink/fp/blob/master/bin/linux_x86_64/fp?raw=true)

and put it in your $PATH

### From sources

```
go get github.com/atotto/clipboard
go get github.com/briandowns/spinner
go build
```

## Configuration

You have to put your Filepicker API Key either using environment variable:

```
export FILEPICKER_APIKEY=YOUR_API_KEY_HERE
```

or by using config file **~/.fp**

```
[Filepicker]
apikey = YOUR_API_KEY_HERE
```

## Usage

```
fp /path/to/your/file
```

![alt tag](https://www.filepicker.io/api/file/sQ1ZcmLUTcuterQQpGpv/convert?w=800)
