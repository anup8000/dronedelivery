# Drone Delivery Scheduling

## Commands

Below are convinience commands to test and run the program with all defaults:

```go
make test
make run
```

Alternatively, you can provide the input file and output file:

```go
./cmd/dronedelivery/dronedelivery -i 
```

If docker is your poison - the image can be found here: <https://cloud.docker.com/u/anup8000/repository/docker/anup8000/dronedelivery>
The docker image is built with multi-stage builds and is just about ~3 MB!!

```docker
docker build -t anup8000/dronedelivery -f ./build/package/dockerfile .
docker run --rm --name dronedelivery -v /Users/anup/projects/goprojects/drone-delivery/assets/files:/app anup8000/dronedelivery
```

Kubernetes:

```yaml
kubectl apply -f ./deployments/kubernetes/deployment.yaml
```

## Assumptions

- I have assumed that even though I have a file at hand to optimize schedules, the algorithm should work generally for any set of orders. That means I have not optimized the algorithm for any specific set of orders. So if the drone schedule time is x and if in the file there is a quicker order to process at x + 10 minutes, the drone scheduler will not know or assume that. It will process based on orders at hand at a given time
- Based on the example inputs, it seemed like there is a ~1 second lag between orders (or might be rounding). For simplicity, I have not considered any lag between orders. Adding a lag is easy enough.
- This problem can be solved to various ways - akin to job scheduling in operating systems. In interest of time, I have chosen a semi-greedy algorithm using behavior mimicing a min-heap/priority-queue
- I have assumed starving of some orders is okay as we maximize promoter score
- I realized that increasing the promoters (deliveries < 120 minutes) and reducing detractors is a similar problem to minimizing the wait times. Most of my tests showed that it did not matter which way we go. I went with the latter approach which yeilded similar (or better) results. But again, as I mentioned earlier, there are multiple ways to solve this problem
- I have assumed that the orders will start at the day of processing (can start coming in before 6AM though)
- I have assumed that the drone can be in flight after 9AM to complete the order in progress

## Programming Assumptions/Omissions

- For sake of simplicity, I have used ints for order numbers. In real world using something like big.Int (or BigInt in other languages) should be considered
- Ideally I would have liked to use a readily available data structure like min heap, but I have stayed away from high level structures to make the code as transparent as possible
- I have not used a logger for this project, but for production applications, a logger writing to multiple destinations is recommended. Ideally it should stream logs to a splunk/fluentd/elastisearch backend
- I have handled many common errors, but in production application it is better to have very detailed error handlers with specific messages. I have followed the principle that the caller of the component should take responsibility to process the error. The helper components will just log and rethrow the error.
- The built docker image is not ready for production. A production image would run the process under a restricted user. Also security testing should be done on the docker images
- Some of the tests output console logs, in production there should be a logger which should be configured to log to a separate location for tests

## Language and Design considerations

### Language consideration

The reasons for using Go as a programming language for this assignment are:

- Go compiles and runs natively, not on a VM or any virtualization layer. This makes it super fast for console based application
- There is no cold-start wait time like in Java or .NET
- Go docker images are tiny! The one I built for this assignment is ~3 Mb!!
- No prerequisite softwares needed for installation once you have the Go binary
  
### Design considerations

Golang is a radically different language than Java or C# or even NodeJS or Python for that matter. The creators and community users of Go have laid down certain precepts to adhere to. Some of the design considerations for this project are based on it.

- Go is not object oriented in the sense as what Java or C# is. Code structure is opinionated and different from other OO based languages
- Functional division: Code is divided based on function, not on layers
- Functional co-location: Similar to earlier point, execution and tests are co-located based on function
- Common artifacts are separated out from application logic code to their own packages
- Have used packages - which in Go are akin to Services/Repositories/Helpers in other languages