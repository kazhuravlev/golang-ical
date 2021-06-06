package ics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertyParse(t *testing.T) {
	tests := []struct {
		Input    string
		Expected func(output *BaseProperty) bool
	}{
		{Input: "ATTENDEE;RSVP=TRUE;ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *BaseProperty) bool {
			return output.IANAToken == "ATTENDEE" && output.Value == "mailto:employee-A@example.com"
		}},
		{Input: "ATTENDEE;RSVP=\"TRUE\";ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *BaseProperty) bool {
			return output.IANAToken == "ATTENDEE" && output.Value == "mailto:employee-A@example.com"
		}},
		{Input: "ATTENDEE;RSVP=T\"RUE\";ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:mailto:employee-A@example.com", Expected: func(output *BaseProperty) bool { return output != nil }},
		{
			Input: `X-APPLE-STRUCTURED-LOCATION;VALUE=URI;X-ADDRESS="Берсеневская набережная 6\\\\nМосква\\nМосква\\n\\n119072";X-APPLE-RADIUS=48.71308504224796;X-TITLE="Берсеневская набережная 6":geo:53.740176,33.608891`,
			Expected: func(out *BaseProperty) bool {
				if out == nil {
					return false
				}

				return out.IANAToken == "X-APPLE-STRUCTURED-LOCATION"
			},
		},
	}
	for i, test := range tests {
		output := ParseProperty(ContentLine(test.Input))
		assert.True(t, test.Expected(output), i)
	}
}
