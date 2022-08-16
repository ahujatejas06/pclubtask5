# Student Registry!

The backend server has been written using the gin microframework and SQLite database.

## Installation
To run the backend, Go 1.18 must be installed along with GCC version 5. GCC is used to compile the SQLite driver module of the Go language. This can be done using a linux machine running a debian based OS. First install linuxbrew  from here https://docs.brew.sh/Homebrew-on-Linux by following the instructions exactly. 

Then run the command 
```
brew install gcc@5
```
After this run 
```
go run main.go
```
in the repository

Headover to **localhost:8080/** to connect to the server,

## Capabilities
The backend is capable of adding, deleting and updating student records. The routes available in the backend are 

- GET request on /  -->  This returns all the data in the database.
- GET request on /:id --> This returns the data of student having given id.
- POST request on / --> Give input in JSON format to add a record in the database.
- DELETE request on /:id --> Deletes the record having given ID.
- PUT request on /:id --> Updates the record having given ID.

The template for JSON is as follows:-

```
{
  "name": "Tejas Ahuja",
  "branch": "Batti",
  "roll_number":211105,
  "userid":"tejasahuja21"
}
```

For the PUT request, you only need to give the fields you want to edit in the JSON.


This is a very basic server and I have tried to keep the variable names as simple as possible.

### Errors
In case of errors, these are returned automatically in the data part of the returned object.

You can use thunderclient extension on VSCode to give input to the backend server. 




