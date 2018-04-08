/*
Metrics provides an instrumentation layer, which the application can use to
record certain metrics. Prometheus is used for monitoring.

The following metrics are tracked:

- start_count (Counter)
	- Counts number of application startups
    - Labels:
        - status
            - Indicates if the application is running
            - valid values: `ok | error`
- health (Counter)
    - Counts number of health pings
    - Labels:
        - status
            - Indicates if the application is functioning correctly
            - valid values: `ok | error`
- chat_event
	- Labels:
		- effect
			- Indicates how app responded
			- valid values: `handled | unhandled`
		- type
			- Name of chat api event type
			- valid values: `*`
	- count (Counter)
		- Counts number of chat events received
	- duration_seconds (Histogram, lower 0.01s, width 0.01s, count 15)
		- Measures the time it takes to process a chat event
- cmd
	- Labels:
		- name
			- Name of the command
			- valid values: `*`
		- status
			- Indicates if the cmd ran correctly
			- valid values: `ok | error`
	- count (Counter)
		- Counts number of cmd invocations
	- duration_seconds (Histogram, lower 0.01s, width 0.01s, count 15)
		- Measures the time it takes to run a command

# Common labels
All metrics will have a few common labels:

- user_id: One way hashed ID of user metric relates to
- environment: Which level of deployment metric came from
	- Valid values are:
		- develop
		- test
		- staging
		- production
*/
package metrics
