# Cron2e

A small utility written in Go for interpreting cron expressions in plain
english.

![cron2e_example](https://user-images.githubusercontent.com/30610148/78629100-71297480-78d9-11ea-84c0-d354eb22e7a2.png)

Design goals:

- Support as many cron formats as is practical (e.g. [AWS
rate and cron
expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html))

- Be quick to run at the command line, easy to use, and easy to integrate into a workflow

- Perform some basic validation on the values for a given schedule

## Installation

First, ensure your `GOPATH` has been configured correctly.

Fetch the package:

``` shell
go get -u github.com/cursande/cron2e
```

Then install it locally:

``` shell
go install github.com/cursande/cron2e
```

## Usage

``` shell
> cron2e "5 0 * 3-8 *"
Runs every day from months March through August at 00:05
```

