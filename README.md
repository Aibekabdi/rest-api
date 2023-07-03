# Simple REST API Server in Go

This is a simple REST API server implemented in Go. It provides several functions including substring finding, email validation, counter operations using Redis, and basic CRUD operations for users.

## SubstrFind

The `substrFind` function finds the longest substring in a given string. It can be accessed via the following endpoint:

Replace `{string}` with the input string for which you want to find the longest substring.

## EmailCheck

The `emailCheck` function validates whether an email address is valid or not. It can be accessed via the following endpoint:
Replace `{email}` with the email address you want to validate.

## Counter

The `counter` function performs addition and subtraction operations on a counter value stored in Redis. It supports the following endpoints:

- Get the current counter value:
- Increment the counter:
- Decrement the counter:

  
## User CRUD

The server supports basic CRUD (Create, Read, Update, Delete) operations for users. The endpoints for user CRUD are as follows:

- Create a new user:
- Get information about a specific user:
- Update information of a specific user:
- Delete a specific user:
- 
Replace `{id}` with the unique identifier of the user you want to perform the operation on.
