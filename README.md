# personal-go-graphql
Personal prototype for GraphQL using Go Language

# How to run
```shell
cd graphql-prototype
go run server.go
```

# Sample use
Access GraphQL playground by visiting http://localhost:8080/
```
# Create data
mutation createTodo {
  createTodo(input: {text: "todo", userId: "1"}) {
    user {
      id
    }
    text
    done
  }
}

# Get data
query findTodos {
  todos {
    user {
      id
      name
    }
  }
}
```
