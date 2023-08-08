# Web Crawler
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
* same links will not be omitted
* It will take a few seconds to generate response because it needs to check if links are accessible or not (though it is improved with limited concurrency (10 go routines) and local cache)
* concurrency worker and port can be changed in files/configurations.json
* there is no timeout in getting response from the links in the site, might take extra time to gather the data depending on the links

## Instructions
1. Clone from master branch
2. Use provided Makefile to start the server with `make` / `make gochallenge`
3. Open up the endpoint `/pagedetail` in `localhost:1234`,query parameters:
   - `url`    : - mandatory
                - put url input in here
                - string
                - example: `url=https://abc.com`  
   example: `https://localhost:1234/pagedetail?url=https://abc.com`
