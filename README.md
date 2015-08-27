#ezconf

====

I did not like any of the configuration options with Go.   Creating a struct with every possible value I might want to ready it out of a JSON or yaml file?  really?

I made one I think is more dynamic, while at the same time simple to use.


## Installation

        go get github.com/chrisftw/ezconf

## Usage

1. Make your conf file in the config directory.

        # This is a comment
        # file named "[[namespace]].conf"
        key            : value
        first-name     : Baron
        middle-name    : Von
        last-name      : Awesomesauce

2. use your settings

        Lastname := ezconf.Get("namespace", "last-name")

3. add more settings with code

        port := ezconf.Set("namespace", "nick-name", "Junior")
