package gcpregionquery

import "testing"

func TestParseDataCenterName(t *testing.T) {
	zone, _ := parseDataCenterName(`projects/1077355926250/zones/europe-west1-b`)
	if zone != `europe-west1` {
		t.Errorf("Expected zone to be europe-west1 got `%v`", zone)
	}

	us, _ := parseDataCenterName(`projects/1077355926250/zones/us-central1-b`)
	if us != `us-central1` {
		t.Errorf("Expected zone to be us-central1 got `%v`", us)
	}

	dc, err := parseDataCenterName(`foobar`)
	if err == nil {
		t.Errorf("Expected zone to be unparsable got `%v`", dc)
	}
}
