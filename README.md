# Heartbeat
A http healthchecker written in Go

# Usage
Run the executable with a mandatory argument specifying the path 
of the Heartbeat files.

```
~/HeartbeatData
    # Required file containing info on which urls to test
    /data.json  
    # Pretty print info about the tests (created automatically)
    /pretty.txt 
```

For example:

`$ ./heartbeat ~/HeartbeatData`

data.json example

```
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
```

It will overwite that file with heartbeat information.

# Reporting

Once you install this package `go install`
Setup crontab as follows (if you have mail setup on your server)

```
# Run heartbeat every hour
0 * * * * Heartbeat PATH_TO_DATA_DIR
# Email heartbeat results every day at 1am
0 1 * * * mail -s 'Heartbeat' YOUREMAIL@gmail.com < PATH_TO_DATA_DIR/pretty.txt
```