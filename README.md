# covstats

REST APIs for covid stats of India and its States from mohfw,
with reverse geocoding support to get covid stats of relevant 
State.

Built with Golang, Echo framework and MongoDB


## How to run

 1. Install golang, mongodb. Run mongodb  

 2. Clone the project outside of `$GOPATH` directory 

 3. Go to root directory of project and make changes to `config/config.json`

 4. Run `/path/to/covstats$ go mod tidy` and `/path/to/covstats$ go run server.go"` to start the server
 
 6. For swagger page, open browser and visit `localhost:1323/swagger/index.html`. 
    Set `lat:16.3` and `lon:80.4` for testing reverse gecoding API use case.


## Developer

 1. To run tests go to handlers folder `path/to/covstats/handlers$ go test -v`

 2. If swagger declarative comments are modified, you need to run `path/to/covstats$ swag init -g ./server.go --output ./docs/` 



## Application

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

## Future Work

1. Heroku Deploy
2. Unit tests require more cases
3. Swagger docs can be packed with more info
4. Validation of lat, lon needs review, handled by locationiq server tentatively
5. Static assets
6. User authentication

## Preview of APIs with Swagger

<img src="swagger-preview.PNG" width=1000 />

## References

1.  [project skeleton: code structure and mongodb setup](https://github.com/sangianpatrick/go-echo-mongo)
2.  for sorting results based on time: [link1](https://gist.github.com/border/3489566) [link2](https://pkg.go.dev/labix.org/v2/mgo#Query.Sort)
3.  [statewise covid data](https://www.mohfw.gov.in/data/datanew.json)
4.  [intro to echo and making client api calls](https://betterprogramming.pub/intro-77f65f73f6d3)
5.  [more info on echo](https://blog.logrocket.com/making-http-requests-in-go/)
6.  [go quick intro](https://www.youtube.com/watch?v=C8LgvuEBraI)
7.  [swagger api docs with examples](https://github.com/swaggo/swag)
9.  [testing apis](https://echo.labstack.com/guide/testing/)
10. [setup for go tests](https://stackoverflow.com/questions/28240489/golang-testing-no-test-files/28240537)
11. [command to run go tests](https://ieftimov.com/post/testing-in-go-go-test/)
