# Microservices Architecture with Go

## Overview

For a long time, web applications were primarily developed as monolithic applications, handling everything from user authentication to email notifications. While this approach remains popular, many larger-scale applications are shifting towards a microservices architecture. This architectural style structures an application as a loosely coupled collection of smaller, independent services that can be developed, deployed, and maintained independently.

## Features of Microservices

- **Maintainability and Testability:** Each microservice is small and focused, making it easier to maintain and test.
- **Loose Coupling:** Microservices are loosely coupled, reducing dependencies between them.
- **Independent Deployment:** Each service can be deployed independently, allowing for easier updates and scaling.
- **Business-Centric Organization:** Microservices are organized around specific business capabilities.
- **Small Team Ownership:** Each microservice can be owned by a small team, fostering accountability and ownership.

## Project Description

In this project, contained self-contained microservices that communicate with each other and a simple front-end application. The services we will build include:

- **Front End Service:** Displays web pages.
- **Authentication Service:** Handles user authentication with a PostgreSQL database.
- **Logging Service:** Manages logs using a MongoDB database.
- **Listener Service:** Receives messages from RabbitMQ and acts upon them.
- **Broker Service:** An optional entry point into the microservice cluster.
- **Mail Service:** Converts JSON payloads into formatted emails and sends them out.

