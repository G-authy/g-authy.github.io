# Secrets API

This is a simple API built with Go and Gorilla Mux for managing secrets.

## Structures

The API deals with `Secret` objects, which have the following structure:

```go
type Secret struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

## The API hEndpoints
The API has the following endpoints:

* GET /secrets: Get all secrets.
* POST /secrets: Create a new secret.
* PUT /secrets/{id}: Update a secret.
* DELETE /secrets/{id}: Delete a secret.

## Running the Project
To run the project, use the following command ```go run secrets.go ```

This will start the server at port 8080.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
**MIT**