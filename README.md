# Go Gin Dockerized
> Simple and Fast API CRUD service 

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)

## General Information
- CRUD REST API suitable for fast web app development
- Works with Oracle DB (might switch to any other easily)
- Built on Gin Web Framework
- Containerized with Docker
- Basic logic for login authentication
- LDAP authentication
- Postman collection included

## Technologies Used
- Go 1.19

## Features
List the ready features here:
- LDAP Login
- Logged-in user authentication (token-based)
- Common CRUD services

## Setup
1. Install Go 1.19 and download DEV VM with Oracle database. Use Virtual Box or other VM software.
2. Set-up the DB, create your schema, create the tables using the DDL in db.sql.
3. Update config.json with connection parameters and login parameters.
4. Follow the instructions in instr.txt.
5. Use Postman collection (GinWs.postman_collection.json) for testing

## Usage
- Do not expose this service directly to the end users, as no security is implemented
- Deploy with nginx or other reverse proxy/load balancing/WAF service where the main network security logic is implemented
