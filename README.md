## Installation

### From binary

Just download one of the binaries:

[Darwin x86_64](bin/darwin_x86_64/fp)

[Linux x86_64](bin/linux_x86_64/fp)

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
