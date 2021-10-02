# covstats

REST APIs for covid stats of India and its States from mohfw,
with reverse geocoding support to get covid stats of relevant 
State. Built with Golang, Echo framework and MongoDB



## How to run

 1. Install golang, mongodb and postman. Run mongodb  

    You can comment middleware portion of code in server.go after this step
    to skip using postman and using browser to make the API calls

 2. Go to root directory of project and fill `config/config.json`

 3. Run `/path/to/covstats$ go mod tidy`
 
 4. Run `/path/to/covstats$ go run "path\to\covstats\server.go"` to start the server
 
 3. Import covstats.postman_collection.json in postman and make calls to server for response.


## Application

>The request header should contain ( middleware part of code in server.go should not be commented for this):
```Content-Type: "application/json"```
>The error response should be:

```json
{
  "code": "<HTTP STATUS CODE: Error>",
  "data": null,
  "message":"Error message",
  "success": false
}
```

>The success response should be:

```json
{
  "code": "<HTTP STATUS CODE: Success>",
  "data": "<MULTI DATA TYPE: array, stirng and object>",
  "message":"Success message",
  "success": true
}
```

## Future Work / Todos

1. Swagger docs can be improved 
2. Data validation for /lat/:lat/lon/:lon before calling location1q server
3. Static assets
4. Heroku Deploy


## References

1.  https://github.com/sangianpatrick/go-echo-mongo           -  skeleton of project: code structure and mongodb setup
2.  https://www.mohfw.gov.in/data/datanew.json                -  for statewise covid data
3.  https://betterprogramming.pub/intro-77f65f73f6d3          -  intro to echo and for making client api calls 
4.  https://blog.logrocket.com/making-http-requests-in-go/    -  more info on echo 
5.  https://www.youtube.com/watch?v=C8LgvuEBraI               -  go quick intro 

6.  https://gist.github.com/border/3489566 
    https://pkg.go.dev/labix.org/v2/mgo#Query.Sort            -  for sorting results based on time