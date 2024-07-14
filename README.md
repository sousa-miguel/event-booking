# Event Booking - GO 

Dummy project to better understand Golang and work with REST APIs.

## Technologies Used
* Go
* Gin
* SQLite

## Project Description
This project is a REST API built using Golang. It is designed to handle event booking operations, providing endpoints to create, read, update, and delete events. This project aims to offer a hands-on learning experience with Golang and RESTful API development.

## Setup Instructions

1. **Clone the repository:**
    ```bash
    git clone https://github.com/yourusername/event-booking-go.git
    cd event-booking-go
    ```
2. **Create .env file**
    ```bash
    touch .env
    echo "SUPER_SECRET=EnglishOrSpanish" >> .env 
    ```
3. **Run the application:**
    ```bash
    go run main.go
    ```

## API Endpoints

### Event Endpoints
- **Create Event**
    - **URL:** `/events`
    - **Method:** `POST`
    - **Description:** Create a new event.
    - **Request Body:** 
        ```json
        {
            "Name":"Event 1",
            "Description":"Event desc",
            "Location":"plane",
            "DateTime":"2024-07-14T09:18:00.000Z"
        }
        ```
    - **Response:**
        ```json
        {
            "ID": 1,
            "Name": "Event 1",
            "Description": "Event desc",
            "Location": "plane",
            "DateTime": "2024-07-14T09:18:00Z",
            "UserID": 1
        }
        ```

- **Get All Events**
    - **URL:** `/events`
    - **Method:** `GET`
    - **Description:** Retrieve a list of all events.
    - **Response:**
        ```json
        [
            {
                "ID": 1,
                "Name": "Event 1",
                "Description": "Event desc",
                "Location": "plane",
                "DateTime": "2024-07-14T09:18:00Z",
                "UserID": 1
            },
            ...
        ]
        ```

- **Get Event by ID**
    - **URL:** `/events/{id}`
    - **Method:** `GET`
    - **Description:** Retrieve a specific event by ID.
    - **Response:**
        ```json
        {
            "ID": 1,
            "Name": "Event 1",
            "Description": "Event desc",
            "Location": "plane",
            "DateTime": "2024-07-14T09:18:00Z",
            "UserID": 1
        }
        ```
- **Update Event by ID**
    - **URL:** `/events/{id}`
    - **Method:** `PUT`
    - **Description:** Update a specific event by ID.
    - **Request Body:** 
        ```json
        {
            "Name":"Event 1 (updated)",
            "Description":"Event desc",
            "Location":"plane",
            "DateTime":"2024-07-14T09:18:00.000Z"
        }
        ```
    - **Response:**
        ```json
        {
            "message": "Event updated successfully!"
        }
        ```
- **Delete Event by ID**
    - **URL:** `/events/{id}`
    - **Method:** `DELETE`
    - **Description:** Delete a specific event by ID.
    - **Response:**
        ```json
        {
            "message": "Event deleted successfully!"
        }
        ```

### User Endpoints
- **Create User**
    - **URL:** `/signup`
    - **Method:** `POST`
    - **Description:** Create a new user.
    - **Request Body:** 
        ```json
        {
            "email":"email@example.com",
            "password":"supersafe123"
        }
        ```
    - **Response:**
        ```json
        {
            "message": "user created successfully!"
        }
        ```
- **Login**
    - **URL:** `/login`
    - **Method:** `POST`
    - **Description:** Authenticate to the API.
    - **Request Body:** 
        ```json
        {
            "email":"email@example.com",
            "password":"supersafe123"
        }
        ```
    - **Response:**
        ```json
        {
        "message": "login successfull!",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQGV4YW1wbGUuY29tIiwiZXhwIjoxNzIwOTkwNTQzLCJ1c2VySWQiOjB9.Y8vtleH4jv2tb9VBQPVfi7otG1qUZCBc5qDck4kJFBw"
        }
        ```
