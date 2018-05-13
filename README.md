# Heartbeat
A golang http healthchecker

# Usage
Run the executable with a mandatory argument specifying the path to a newline-separated file of sites to check.  For example:

`$ ./heartbeat ./conf/sites.txt`

sites.txt example

`
https://jdoleary.me
http://example.com
http://example.com/specificRoute
http://subdomain.example.com
`