# Sparta Demo

A mostly straight copy of the API Gateway example from [Sparta][1]

I used `dep` for vendoring so be sure to `go get -u github.com/golang/dep/cmd/dep` if you haven't.

## Setup

Check things out
```
$ git clone ssh://git@github.com/turgon/sparta-demo
```

Handle dependencies
```
$ dep ensure
```

Make sure your S3 bucket is set up. I used "the-bucket" below.

And naturally you're using AWS so be sure your env vars are present and correct:
```
$ export AWS_REGION=us-east-1
$ export AWS_ACCESS_KEY_ID=XXXXXXXXXX
$ export AWS_SECRET_ACCESS_KEY=YYYYYYYYYY
```

Finally, provision the service!
```
$ go run main.go provision --s3Bucket the-bucket
```

## Hitting the service

You will need to grab the new API Gateway's base url from the `provision` output. It's in an output log line with `Key=APIGatewayURL`. It should look something like `https://blablahblah.execute-api.us-east-1.amazonaws.com/prod`

```
$ export baseURL=https://blablahblah.execute-api.us-east-1.amazonaws.com/prod
```

Now you can make requests:

```
$ curl -sv $baseURL/prod/hello/world/test
```


[1]: http://gosparta.io/docs/apigateway/example1/
