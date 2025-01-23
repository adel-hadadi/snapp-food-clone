# Food Delivery Clone App

Welcome to the Food Delivery Clone App! This project replicates the core features of a food delivery platform, enabling users to browse restaurants, order food, and track deliveries.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

This application is designed to simulate the functionality of a modern food delivery service. Users can browse restaurant menus, place orders, and track delivery status. The platform includes a robust backend.

---

## Features

- **User Authentication**: Sign up, log in, and manage profiles.
- **Restaurant Listings**: Browse restaurants and their menus.
- **Order Management**: Add items to the cart, place orders, and view order history.
- **Admin Panel**: Manage restaurants, menus, and orders.

---

## Tech Stack

### Backend:

- **Language**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL
- **Messaging Queue**: RabbitMQ
- **Authentication**: JWT

### Frontend:

- **Framework**: React and NextJS
- **Styling**: TailwindCS

### Others:

- **Containerization**: Docker, Docker Compose
- **Cloud Storage**: MinIO

---

## Getting Started

To get a local copy of the project up and running, follow these steps.

### Prerequisites

- Docker
- Docker Compose
- Golang (for development)
- Node.js and npm/yarn (for frontend development)

---

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/adel-hadadi/snapp-food-clone.git
   cd snapp-food-clone
   ```

2. **Set up environment variables**
   Create a `.env` file in the root directory and configure the required environment variables. See `.env.example` for reference.

3. **Build and start backend services**

   ```bash
   make up-build
   ```

4. **Start frontend**

   ```bash
   make up-front
   ```

5. **Access the application**
   - Backend API: `http://localhost:8080`
   - Frontend: `http://localhost:3000`

---

## Usage

- **Frontend**: Open `http://localhost:3000` in your browser.
- **API Documentation**: Access Swagger/OpenAPI docs at `http://localhost:8080/docs`.

---

## Project Structure

```
snapp-food-clone/
├── cmd/
│   ├── api/
│   ├── app/
│   ├── cron/
├── data/
├── docker/
├── docs/
├── internal/
│   ├── pkg/
│   ├── Dockerfile
│   └── ...
├── pkg/
├── scripts/
├── tmp/
├── web/
│   ├── src/
│   ├── public/
│   ├── Dockerfile
│   └── ...
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add YourFeature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgements

- Inspired by platforms like Uber Eats, DoorDash, and Zomato.
- Special thanks to the open-source community for the tools and frameworks used in this project.

---

Thank you for checking out the Food Delivery Clone App! If you have any questions or suggestions, feel free to open an issue or contact me directly.
