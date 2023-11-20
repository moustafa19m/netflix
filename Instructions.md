## Building and Running the Project
This project depends on a correspondence with Giphy's API. Before you can run this project, you need to get a valid API key.


Once you have the API key handy, you need to set it up as an environment variable.
```
export GIPHY_API_KEY='YOUR_KEY_HERE'
```


You will then need to ```cd``` in the project root directory and run the commands to build and run the go project:
```
go build .
go run .
```


Upon successful compilation and execution, you should be able to hit ``` http://localhost:8080/query```. You will need to provide at least one query params argument with key ```searchTerm```. If you want to provide multiple terms, you need to provide them as such:
```
http://localhost:8080/query?searchTerm="a"&searchTerm="a"&searchTerm="c"
```

If you provide two params of the same value, the response will contain only one response for both of them. In other words, duplicates are ignored. Number of unique search terms can not be more than 50.


I used a beta API key to test this project myself, and due to the limit of 1,000 requests per day, I set hard limit of 50 queries per request to this service, since every query will trigger a new call to Giphy's API.

For each unique term you provide, there will be a maximum of 25 results. This number can be changed in query params, and ideally it should be a variable that the user can set within bounds of API permissions and bandwith available.

The order of the output is not guaranteed, and it will not necessarily match the order in which the params were provided.



## Running the Tests
There a tests in this project that mock the correspondence with the Giphy API. You don't need to have a valid key to run these tests. To run the tests, execute the following commands from the root directory:
```
go build .
go test ./...
```


## To Be Production-Ready
1. __Secret Management__: Currently, the API key is explicitly defined by the user, and it can be viewed in terminal logs in plaintext. The key is better stored in a vault and retrieved at the startup time of the service.


2. __Dockerized and Deployed__: I used go version 1.21.4, however, this doesn't guarantee that the project will work with any other version. For production stability, having a docker image will be required.


3. __Managing Go Routines__: I created a go routine for every param request. This speeds up the process of capturing all the required data. Even though go routines are light and the likelihood of exceeding memory capacity is low, in the rare event that it happens due to more than expected requests, we need to have a management system for these routines. This would enable us to kill routines that are taking too long, as well as limit the number of running routines if necessary.


4. __Setting Contexts__: In reference to the previous point, where some routines could take longer than expected, we can utilize contexts to expire if we don't receive a response from the API within a set time limit. This would free up compute resources to enable us to make more requests.


5. __Caching__: The limiting factor in the optimization of this service is the wait time from Giphy's API. To minimize the impact on this service's latency to respond back to the caller, we can leverage caching methods. This has two major benefits: decreased wait time and compute time, as well as lowered billing from Giphy's API since their production endpoint is paid.


6. __Paging__: Currently, I am limiting the number of gifs per search term to 25. Ideally, I would set that number variable for the user to set up until a maximum number (for example, 100). And to allow the user to capture more gifs per search term, I would enable paging so that they can view the next 100 gifs and so on.


7. __Parameterizing Other Fields__: Many of the params required to make a call to Giphy's API are hard-coded. Ideally, these would be dynamically decided based on the user's location or preferences, for example, in deciding the language, or passed through from the user directly.


8. __Changing Max Terms Allowed Dynamically__: Based on usage and the allowed number of requests per day, I would change the hard limit that I set of 50 requests per call. I would make the limit variable such that we allow more search terms if we have decreased traffic and a higher bandwidth than usual to make more calls.


9. __Adding More Tests__: I added tests for the API and the set that I created. I would also want to add tests to the service itself and the web package.

10. __Monitoring__: We would want to watch the performance of the service as well as the processes that it runs. To that end, we would need to leverage a monitoring system such as Prometheus and a log aggregator such as Datadog. This will also require the addition of searchable and IDed log statements. I attemted this effort in the project by creating a ULID for each request, but I would like to add more logs and more infomation per log, such as user, calling ip, etc.

