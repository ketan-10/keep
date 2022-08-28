# Google Keep Clone

## Backend

**gqlgen** <br>

- [Getting Started](https://gqlgen.com/getting-started/)

- Initialize
    - To initialize the gql project we use `go run github.com/99designs/gqlgen init`
    - we can use `go get github.com/99designs/gqlgen` to get the tools required to init.  
    - Or include this in `tools/tools.go` and run `go mod tidy` to auto import. So that compiler will not auto remove the tools dependencies. 
    ```go
    //go:build tools
    // +build tools

    package tools

    import (
        _ "github.com/99designs/gqlgen"
    )
    ```
    - Go does not support dev-dependency. So as per my knowledge the above tools will be in prod executable also. and it's not practical to do `tree-shaking` each time we build. <br>
    
    - For dev-dependency go recommend to use direct binary of your project in cmd folder.



  

