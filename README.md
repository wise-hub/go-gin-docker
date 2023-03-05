# Go Gin Dockerized
> Simple and Fast API CRUD Service with LDAP Authentication

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
- Logged-in user authentication (token-based, offline and online validation)
- Common CRUD services

## Setup
1. Install Go 1.19 and Docker
2. Download Dev VM with Oracle database. Use Virtual Box or other VM software
3. Set-up the DB, create your schema, create the tables using the DDL in setup/db.sql
4. Update config.json with connection parameters and login parameters
5. Follow the instructions in setup/instr.txt
6. Use Postman collection for testing (setup/GinWs.postman_collection.json)

## Usage
- Deploy with nginx or other reverse proxy/load balancing/WAF service
