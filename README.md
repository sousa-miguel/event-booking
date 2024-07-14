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
2. **Run the application:**
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