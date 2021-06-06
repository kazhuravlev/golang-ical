# golang-ical

A ICS / ICal parser and serialiser for Golang.

Because the other libraries didn't quite do what I needed.

Usage, parsing:

```go
package main

import (
	"fmt"
	ics "github.com/kazhuravlev/golang-ical"
	"net/http"
)

func main() {
	const calendarURL = "https://ical.mac.com/ical/Irish32Holidays.ics"
	resp, err := http.Get(calendarURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	cal, err := ics.ParseCalendar(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(cal)
}
```

Creating:

```go
cal := ics.NewCalendar()
cal.SetMethod(ics.MethodRequest)
event := cal.AddEvent(fmt.Sprintf("id@domain", p.SessionKey.IntID()))
event.SetCreatedTime(time.Now())
event.SetDtStampTime(time.Now())
event.SetModifiedAt(time.Now())
event.SetStartAt(time.Now())
event.SetEndAt(time.Now())
event.SetSummary("Summary")
event.SetLocation("Address")
event.SetDescription("Description")
event.SetURL("https://URL/")
event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
return cal.Serialize()
```

Helper methods created as needed feel free to send a P.R. with more.
