# go_rest_api_static
Running the app
==> go run .
App will run on Port 8080
You can see the response of Api on browser for get requests or via curl requests

Curl requests:

get- curl localhost:8080/books

post- curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

