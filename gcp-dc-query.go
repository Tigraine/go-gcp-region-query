// Query the GCP Metadata API to retrieve the Google GCP Region name your VM is currently running in
package gcpdcquery

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// GetLocalDCWithTimeout returns the local GCP Region w.
// a configurable Timeout for the HTTP Request to the GCP Metadata API.
// Returns the DC name like europe-west1-c.
func GetLocalDCWithTimeout(timeout time.Duration) (string, error) {
	zoneNameQueryURL := `http://metadata.google.internal/computeMetadata/v1/instance/zone`
	req, err := http.NewRequest(`GET`, zoneNameQueryURL, nil)
	if err != nil {
		return ``, err
	}
	client := &http.Client{
		Timeout: timeout,
	}
	req.Header.Add(`Metadata-Flavor`, `Google`)
	res, err := client.Do(req)
	if err != nil {
		return ``, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ``, err
	}
	dc, err := parseDataCenterName(string(body))
	return dc, err
}

// GetLocalDC returns the local GCP Region with a default Timeout of 1s
func GetLocalDC() (string, error) {
	return GetLocalDCWithTimeout(1 * time.Second)
}

var dcregex = regexp.MustCompile(`\/([\w-]*)-.*$`)

func parseDataCenterName(gceRegionString string) (string, error) {
	res := dcregex.FindAllStringSubmatch(gceRegionString, -1)
	if len(res) == 0 || len(res[0]) < 2 {
		return ``, fmt.Errorf(`Cannot parse DataCenter Name from %s`, gceRegionString)
	}
	return res[0][1], nil
}
