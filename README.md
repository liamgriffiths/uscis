# uscis

I wrote this little script to text me the current status for a pending
[USCIS](https://www.uscis.gov/) case to save me the trouble of checking it
constantly.

It works by scraping the results from the status checker and uses twilio
to send the result from the response to your phone.

You can run it using the command line options or using the little helper script
`./run` to pull out the options from the environment. I use it with a cron job
to run on a schedule.

If you want to run it yourself make sure to set these:

```
# required settings
CASE_NUMBER=<Your case number>
TWILIO_ACCOUNT_SID=<Your twillio account id>
TWILIO_AUTH_TOKEN=<The auth token for your twilio account>
TWILIO_FROM_PHONE=<The twilio phone number you're sending texts from>
TWILIO_TO_PHONE=<Your phone number>

# (optional if you want to track errors using Sentry - it's awesome and free for small projects)
SENTRY_DSN=<Sentry DSN to hook up error monitoring>
```
