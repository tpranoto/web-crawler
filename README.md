# gochallenge
## Summary
a web application served in a REST API to take URL in query parameter and provides several information in the json response:
* HTML Version
* Page Title
* Heading Counts
* Internal and External Links:
   - Count
   - Links
* Inaccessible Links:
   - Count
   - Links
   - Reason links are inaccessible
* Has Login Form

## Instructions
1. Use provided Makefile to start the server with `make` / `make gochallenge`
2. open up the endpoint `/pagedetail` with `url` as the query parameter  in `localhost:1234`
    example: `https://localhost:1234/pagedetail?url=abc`
3. replace `abc` in query parameter with website URL input
