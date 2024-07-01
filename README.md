# Movies Server

This Go application serves as a basic RESTful API for managing movie data. It allows you to perform CRUD (Create, Read, Update, Delete) operations on movie records using HTTP requests.

## Features

- List all movies
- Get a specific movie by ID
- Create a new movie record
- Update an existing movie record
- Delete a movie record

## Prerequisites

To run this server, you need to have [Go](https://golang.org/doc/install) installed on your system.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Cyberguru1/GO-Movies-Server.git
   ```
2. Navigate to the project directory:

   ```bash
   cd Go-Movies-Server
   ```
3. Run the server:

   ```bash
   go run main.go
   ```

   The server will start on port 8000 by default.
4. Access the server on Browser:

```http
http:localhost:8000
```

## Usage

- Use an API client (e.g., [Postman](https://www.postman.com/)) or make HTTP requests to interact with the server's endpoints.
- You can perform various operations such as listing all movies, retrieving a specific movie by ID, creating a new movie record, updating an existing movie record, and deleting a movie record.

## Endpoints

### List all movies

- **GET** `/movies`
- **Mock Up FrontEnd**
- ![1696507010258.png](https://imgur.com/WZNKuqS.png)

### Get a specific movie by ID

- **GET** `/movies/{id}`

### Create a new movie record

- **POST** `/movies`
- **Mock Up FrontEnd**
- ![1696507079035.png](https://imgur.com/jZx5AuP.png)

### Update an existing movie record

- **PUT** `/movies/{id}`

### Delete a movie record

- **DELETE** `/movies/{id}`

## JSON Data Format

The server accepts and returns JSON data in the following format:

```json
{
  "id": "1",
  "isbn": "343434",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doewy"
  }
}
```

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

You can copy and paste this Markdown content into a README.md file in your project repository. Make sure to replace the placeholders such as the repository URL, your username, and any additional information specific to your project.
