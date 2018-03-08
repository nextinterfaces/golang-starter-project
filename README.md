# go-lang-playground

Go-lang starter project

#### Go Commands

Assume $GOPATH workspace structure is:
    
    |
    +- bin
    |
    +- pkg
    | 
    +- src
    
#### Install ####

    $ go install github.com/nextinterfaces/dummy
    
Which is same as:

    $ cd $GOPATH/src/github.com/nextinterfaces/dummy
    $ go install
    
File gets installed at ./bin directory. Let's try it:

    $ $GOPATH/bin/dummy
    $ dummy
 
 #### Build ####
 
    $ cd $GOPATH/src
    $ go build nextinterfaces/util/string_util
    $ go install nextinterfaces/util/string_util
    
    $ go build nextinterfaces_playground/hello
    $ go install nextinterfaces_playground/hello

After the steps above, your workspace should look like this:

    
    bin/
        hello                 # command executable
    pkg/
        darwin_amd64/          # this will reflect your OS and architecture
            go-lang-playground
                stringutil.a  # package object
    src/
        go-lang-playground/
            hello/
                hello.go      # command source
            stringutil/
                reverse.go    # package source

#### Run ####

    $ cd $GOPATH/src
    $ go run nextinterfaces_playground/web/webserver_content.go
    
Alternatively 

    $ go build nextinterfaces_playground/web/webserver_content.go
    $ ./webserver_content
    
    $ go install nextinterfaces_playground/web/webserver_content
    $ webserver_content
    
#### Go Get - Install Third-Party Packages ####

    $ go get gopkg.in/mgo.v2
    
And import as

    import (        
            "gopkg.in/mgo.v2" 
            "gopkg.in/mgo.v2/bson"       
    )
