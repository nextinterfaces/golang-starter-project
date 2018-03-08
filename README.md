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
         |
         +- go-lang-playground
    
Create dependencies to `pkg` 

    $ cd $GOPATH/src
    $ go build nextinterfaces_util/string_util
    $ go install nextinterfaces_util/string_util
    $ go build nextinterfaces_playground/hello
    $ go install nextinterfaces_playground/hello
    
Create executables to `bin` 
    
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

Run executables
 
    $ hello

Alternatively 

    $ go build hello/hello.go
    $ ./hello
    $
    $ go build web/webserver.go
    $ ./webserver
    $
    $ go build web/webserver_content.go
    $ ./webserver_content
    
