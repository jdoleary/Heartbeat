# Heartbeat
A http healthchecker written in Go

# Usage
Run the executable with a mandatory argument specifying the path 
of the saved data containing urls to check.  For example:

`$ ./heartbeat ./data/heartbeats.json`

heartbeats.json example

`
{
    "records":[   
        {
            "url":"https://jdoleary.me"
        },
        {
            "url":"https://jdoleary.me/other"
        }
    ]
}
`

It will overwite that file with heartbeat information.

# Reporting

Once you install this package `go install`
Setup crontab as follows (if you have mail setup on your server)

`
# Run heartbeat every hour
0 * * * * Heartbeat PATH_TO_JSON_FILE
# Email heartbeat results every day at 1am
0 1 * * * mail -s 'Heartbeat' YOUREMAIL@gmail.com < PATH_TO_JSON_FILE
`