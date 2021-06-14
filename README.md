# Problem statement

Create a REST API to implement set functionalities – addItem(), removeItem() and hasItem()

# Solution

1. Developed a Golang service
2. Hosted on AWS EC2
3. Deployment automation through CloudFormation
4. Implemented without using built-in hashmaps or sets
5. Made addItem() operation 2x as fast

# REST API documentation

### GET /addItem/{number}
    Responses:
    200 OK 
        {
            "status": "OK",
            "message": "Item added."
        }
    400 Bad Request 
        { 
            "status": "Bad Request",
            "message": "Item already exists."
        }
        OR 
        {
            "status" : "Bad Request",
            "message" : "Item has to be a number."
        }

### GET /removeItem/{number}
    Responses:
    200 OK 
        {
            "status": "OK",
            "message": "Item deleted."
        }
    400 Bad Request 
        { 
            "status": "Bad Request",
            "message": "Item to be deleted is not present."
        }
        OR 
        {
            "status" : "Bad Request",
            "message" : "Item has to be a number."
        }

### GET /hasItem/{number}
    Responses:
    200 OK 
        {
            "status": "OK",
            "message": "Item is present."
        }
    400 Bad Request 
        { 
            "status": "Bad Request",
            "message": "Item is not present."
        }
        OR 
        {
            "status" : "Bad Request",
            "message" : "Item has to be a number."
        }

# Repo structure

## cloudformation

To deploy AWS resources for the project

## scripts

1. install - Installs golang on ec2

2. build - Builds go service and runs it as a background process

## code

### main.go

Contains code for the go REST API
This service uses third party router Gorilla Mux - https://github.com/gorilla/mux

### addItem2x

Contains code for making the go REST API addItem() operation 2x faster

