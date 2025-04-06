# Reddit Monitor

Reddit Monitor is an application that calculates and retrieves the Top Users and Posts of a subreddit.

## Configuration
In the file main.go, populate the following values in the config struct:
```go
cfg := config.Config{
		Username: "reddit username",
		Password: "reddit user password",
		ClientID: "reddit developer client id",
		Secret: "reddit developer secret",
		ServerPort: 8080,
		Limit: 5,
	}
```
## Execution

To run the application you can execute the following options:

1. With Golang installed on your local, you can run the application from the root of the project.

```bash
go run main.go
```

2. With Docker Engine running, you can build and run the application.
```
docker build -t reddit-monitor:v1 .

docker run -d -p 8080:8080 reddit-monitor:v1
```

To run a reader on a subreddit, in a separate command prompt run the following curl command (replace the "subreddit_name" with the desired subreddit):
```
curl --location 'http://localhost:8080?subreddit=subreddit_name'
```

The application will begin to send responses to the command prompt (Top Users and Posts with the most up votes)