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
    - Go does not support dev-dependency.
    - So to avoid tools to end up in prod executable. We add `// +build tools` at top, As it's not practical to do `tree-shaking` each time we build. <br>
    
    - For dev-dependency go recommend following ([Github Issue](https://github.com/golang/go/issues/25922#issuecomment-1038394599)).

    - After tools are setup Run: **`go run github.com/99designs/gqlgen init`**

- Working:
    - gqlgen is a popular Framework for graphql implementation in go. <br>
    It generate the required code according to your graphql schema file. <br>
    - gqlgen also servers your endpoints. and generate code to attach the resolvers to endpoints. 
    
- Here is a description from gqlgen about the generated files ([*refers](https://github.dev/HaswinVidanage/keep-clone-be)):

    * gqlgen.yml — The gqlgen config file, knobs for controlling the generated code.
    * generated.go — The GraphQL execution runtime, the bulk of the generated code.
    * entities/models_gen.go — Generated models required to build the graph. Often you will override these with your own models. Still very useful for input types.
    * schema/schema.graphqls — This is the file where you will add GraphQL schemas.
    * tmp/resolvers.go — This is where your application code lives. generated.go will call into this to get the data the user has requested.
    * server.go — This is a minimal entry point that sets up an http.Handler to the generated GraphQL server. start the server with go run server.go and open your browser and you should see the graphql playground, So setup is right!

- Dependency
    - Go Does not allow cyclic dependency, it gives error as `import cycle not allowed` <br>
        Though interpreted language like python support it on some level [Example 👆](https://stackoverflow.com/a/744410/10066692) <br>
        I think it should have been possible with [linking step like in C++](https://www.youtube.com/watch?v=H4s55GgAg0I&list=PLlrATfBNZ98dudnM48yfGUldqGD0S4FFb&index=7)

    - [Go team approach to from Dependency management tool before `go mod` or `go dep` on Medium](https://medium.com/@sdboyer/so-you-want-to-write-a-package-manager-4ae9c17d9527) 

    
