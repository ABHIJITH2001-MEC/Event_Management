# Event Management System

## Overview

The Event Management System is designed to provide a comprehensive solution for creating, managing, and registering for various events. The system includes a backend developed in Golang using the Gin framework and Cassandra database (gocql package) for data storage. The frontend can be implemented using any UI framework or Bootstrap/HTML/CSS/JS.

## Design Decisions

### Backend
- **Golang and Gin Framework**: Chosen for their performance, simplicity, and ease of creating RESTful APIs.
- **Cassandra Database**: Selected for its scalability and ability to handle large amounts of data. The gocql package facilitates seamless interactions with the database.
- **RESTful APIs**: Designed to ensure a clear and standardized interface for communication between the frontend and backend.

### Frontend
- **Choice of UI Framework**: React.js is being used here.

### Overall
- **Scalability and Security**: The system is designed to be scalable to handle a growing number of events and users. Security measures are implemented to protect user data and ensure the system's integrity.

## Technologies Used

- **Backend**: Golang, Gin Framework, Cassandra (gocql package)
- **Frontend**: React.js
- **Database**: Cassandra

## Setup Instructions

### Backend

1. Clone the repository:

   ```bash
   git clone https://github.com/ABHIJITH2001-MEC/EVENT_MANAGEMENT.git
   cd event-management-system/backend
2. Install dependencies:
go mod tidy
3. Run the backend server:
go run main.go
## Frontend
cd ../frontend

## Additional Notes
Database Schema: Details of the Cassandra database schema and design can be found in the docs directory.
API Documentation: Detailed documentation of the RESTful APIs is available in the docs directory.
Testing and Optimization: Testing reports and optimization details are documented in the tests directory.


