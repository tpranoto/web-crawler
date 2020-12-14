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

## Notes
* href with `tel:`, `mailto:`, `#`, `javascript` will be counted to internal links
* links that are inaccessible will be removed from internal/external links
* It will take a few seconds to generate response because it needs to check if links are accessible or not (though it is improved with concurrency)

## Instructions
1. Clone from master branch
2. Use provided Makefile to start the server with `make` / `make gochallenge`
3. open up the endpoint `/pagedetail` with `url` as the query parameter  in `localhost:1234`
    example: `https://localhost:1234/pagedetail?url=abc`
4. replace `abc` in query parameter with website URL input
