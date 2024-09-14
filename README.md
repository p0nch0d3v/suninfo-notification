# Sun info notification

The purpose of this simple CLI is send some sun information to cell phone via SMS.

---

Arguments:
- 1. Latitude of the geo location
- 2. Longitude of the geo location
- 3. Phone number to receive the info 

---

The tool uses `https://api.sunrise-sunset.org/` as source of the information.

The tool also uses `twilio` to send the information to the phone number.

The tool sends the sunset and twilight end to the phone number 

The tool also registers the successfully sent in a local database, using `sql-lite`.