# Cron2e

A small utility written in Go for interpreting cron expressions in plain
english.

Design goals:

- Support as many cron formats as is practical (e.g. [AWS
rate and cron
expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html))

- Be quick to run at the command line, easy to use, and easy to integrate into a workflow

- Perform some basic validation on the values for a given schedule

## Usage

``` shell
> cron2e "5 0 * 3-8 *"
Runs from months March through August at 00:05
```

