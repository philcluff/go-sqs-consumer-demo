# Go SQS consumer demo

A "go idiomatic" implementation of a SQS message processor using the new aws-sdk-go.

This implementation long-polls the queue, backing off when the queue is quiet, and dispatches messages onto goroutines for processing.

Warning: This is incomplete and missing some error checking.

TODO:

* Testing
* Full validation of err

## Running:

Set some environment variables for your AWS creds and the app configuration:

    export SQS_QUEUE=https://us-west-2.queue.amazonaws.com/1234567890/queue-name
    export AWS_REGION=us-west-2
    export AWS_ACCESS_KEY_ID=youraccesskey
    export AWS_SECRET_ACCESS_KEY=yourmuchlongersecretkey

Run the app using godep:

    godep go run processor.go
