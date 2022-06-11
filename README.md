# Wsend Key Value Store
*a key-value store for wsend*

## Overview
A persistent, external, key-value store

## Install
To build:
`go build -o wkv main.go`

For prebuilt binaries:
Release

For command line shell usage drop `wkv` into `~/bin` and have `~/bin` in your `$PATH`

For a system wide install `cp wkv /usr/local/bin` would be a suitable location

## Requirements
Not a hard requirment but having [wsend](https://github.com/abemassry/wsend)
installed helps with grabbing the `uid`

## Usage

### General
```
∮ ./wkv --help
wsend key value store is command line tool to store
a value based on a key some examples include:

wkv create --name="foo"
        To create a store

wkv store --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --value="baz" --type="string"
        this will store the value "bar" at the key "foo" and it's of type string
        which is the default if type was "file" then it would attempt to upload the
        file specified. In either case a file always gets uploaded because the string
        value can be very large and it makes more sense to be flexible.

wkv get --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"
        will print the contents of the value to stdout

wkv remove --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"
        to remove the key and associated value and file.

Usage:
  wkv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      create a new key value store
  get         wsend-key-value store get a value based on a key
  help        Help about any command
  remove      delete a key and value
  store       store a value based on a key

Flags:
  -h, --help     help for wkv
  -t, --toggle   Help message for toggle

Use "wkv [command] --help" for more information about a command.
```

### Create Usage
```
∮ ./wkv create --help
Initialize a key value store with a name
wkv create --name="foo"

optional pass in a uid

wkv create --name="foo" --uid="0123456789abcdef"

if uid is not passed in wkv will search for the uid in the wsend install path
incase wsend is installed

Usage:
  wkv create [flags]

Flags:
  -h, --help          help for create
  -n, --name string   name of the key value store
  -u, --uid string    access token
```

### Store Usage
```
∮ ./wkv store --help
Store a value based on a key.

A key is a string and value can be anything but defaults to a string and is
stored as a URL

wkv store --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --value="baz" --type="string"

        --store-link name of the key value store container
        --key name of the key inside key value store
        --value either a string (default) or a file based on the --type
                in either case a file is uploaded which allows the string data to be
                incredbily long and is referenced by a URL pointing to the file.
        --type string or file
                value is either a string (default) or a file specified by --type="file"
                if a file is specified the path is either absolute or the default is the
                current directory
        --uid is optionally passed in like in create

Usage:
  wkv store [flags]

Flags:
  -h, --help                help for store
  -k, --key string          name of the key
  -n, --store-link string   link of the key value store
  -t, --type string         the type of value, defaults to string (default "string")
  -u, --uid string          access token
  -v, --value string        either text or the name of a file
```

### Get Usage
```
∮ ./wkv get --help
wsend-key-value store get a value based on a key
This command gets a value given a key and a store

wkv get --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"

        --store-link name of the key value store container
        --key name of the key inside key value store
        --action optional
                "print" prints the text content if this is a link, it prints the link
                "download" downloads the file to the current directory
                "dump" the default, dumps the contents of the linked file if the link is
                a file

Usage:
  wkv get [flags]

Flags:
  -a, --action string       action to print, download, dump
  -h, --help                help for get
  -k, --key string          name of the key
  -n, --store-link string   link of the key value store
```

### Remove Usage
```
∮ ./wkv remove --help
Delete a key and value

Deletes a key and the value, as well as the backing file of the value.

wkv remove --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --uid="0123456789abcdef"

        --store-link name of the key value store container
        --key name of the key inside key value store
        --uid is optionally passed in like in create

Usage:
  wkv remove [flags]

Flags:
  -h, --help                help for remove
  -k, --key string          name of the key
  -n, --store-link string   link of the key value store
  -u, --uid string          access token
```
