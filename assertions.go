package httptesting

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/golib/assert"
)

// AssertOK tests that the response status code is 200.
func (test *Client) AssertOK() {
	test.AssertStatus(http.StatusOK)
}

// AssertNotFound tests that the response status code is 404.
func (test *Client) AssertNotFound() {
	test.AssertStatus(http.StatusNotFound)
}

// AssertStatus tests that the response status code is equal with the given.
func (test *Client) AssertStatus(status int) {
	assert.EqualValues(test.t, status, test.Response.StatusCode, "Expected response status code "+strconv.Itoa(status)+", but got "+test.Response.Status+".")
}

// AssertContentType tests that the response includes Content-Type header with the given value.
func (test *Client) AssertContentType(contentType string) {
	test.AssertHeader("Content-Type", contentType)
}

// AssertHeader tests that the response includes named header with the given value.
func (test *Client) AssertHeader(name, value string) {
	name = http.CanonicalHeaderKey(name)
	actual := test.Response.Header.Get(name)

	assert.EqualValues(test.t, value, actual, "Expected response header "+name+" with "+value+", but got "+actual+".")
}

// AssertExistHeader tests that the response includes named header.
func (test *Client) AssertExistHeader(name string) {
	name = http.CanonicalHeaderKey(name)

	_, ok := test.Response.Header[name]
	if !ok {
		assert.Fail(test.t, "Response header: "+name+" (*required)", "Expected response header includes "+name+".")
	}
}

// AssertNotExistHeader tests that the response does not include named header.
func (test *Client) AssertNotExistHeader(name string) {
	name = http.CanonicalHeaderKey(name)

	_, ok := test.Response.Header[name]
	if ok {
		assert.Fail(test.t, "Response header: "+name+" (*not required)", "Expected response header does not include "+name+".")
	}
}

// AssertEmpty tests that the response is empty.
func (test *Client) AssertEmpty() {
	assert.Empty(test.t, string(test.ResponseBody))
}

// AssertNotEmpty tests that the response is not empty.
func (test *Client) AssertNotEmpty() {
	assert.NotEmpty(test.t, string(test.ResponseBody))
}

// AssertContains tests that the response contains the given string.
func (test *Client) AssertContains(s string) {
	assert.Contains(test.t, string(test.ResponseBody), s, "Expected response body contains "+s+".")
}

// AssertNotContains tests that the response does not contain the given string.
func (test *Client) AssertNotContains(s string) {
	assert.NotContains(test.t, string(test.ResponseBody), s, "Expected response body does not contain "+s+".")
}

// AssertMatch tests that the response matches the given regular expression.
func (test *Client) AssertMatch(re string) {
	r := regexp.MustCompile(re)

	if !r.Match(test.ResponseBody) {
		test.t.Errorf("Expected response body to match regexp %s", re)
	}
}

// AssertNotMatch tests that the response does not match the given regular expression.
func (test *Client) AssertNotMatch(re string) {
	r := regexp.MustCompile(re)

	if r.Match(test.ResponseBody) {
		test.t.Errorf("Expected response body does not match regexp %s", re)
	}
}

func (test *Client) AssertContainsJSON(key, value string) {
	actual, err := jsonparser.GetString(test.ResponseBody, strings.Split(key, ".")...)
	if err != nil {
		test.t.Errorf("Expected response body contains json key %s with %s, but got Errr(%v)", key, value, err)
	}

	assert.EqualValues(test.t, value, actual, "Expected response body contains json key "+key+" with "+value+", but got "+actual+".")
}

func (test *Client) AssertNotContainsJSON(key string) {
	value, _, _, err := jsonparser.Get(test.ResponseBody, strings.Split(key, ".")...)
	if err == nil {
		test.t.Errorf("Expected response body does not contain json key %s, but got %s", key, string(value))
	}
}

func (test *Client) AssertionContainsJSONInt(key string, value int) {
	actual, err := jsonparser.GetInt(test.ResponseBody, strings.Split(key, ".")...)
	if err != nil {
		test.t.Errorf("Expected response body contains json key %s with %v, but got Errr(%v)", key, value, err)
	}

	valueStr := strconv.FormatInt(int64(value), 10)
	actualStr := strconv.FormatInt(int64(actual), 10)
	assert.EqualValues(test.t, value, actual, "Expected response body contains json key "+key+" with "+valueStr+", but got "+actualStr+".")
}

func (test *Client) AssertionContainsJSONFloat(key string, value float64) {
	actual, err := jsonparser.GetFloat(test.ResponseBody, strings.Split(key, ".")...)
	if err != nil {
		test.t.Errorf("Expected response body contains json key %s with %v, but got Errr(%v)", key, value, err)
	}

	valueStr := strconv.FormatFloat(value, 'e', 3, 10)
	actualStr := strconv.FormatFloat(actual, 'e', 3, 10)
	assert.EqualValues(test.t, value, actual, "Expected response body contains json key "+key+" with "+valueStr+", but got "+actualStr+".")
}

func (test *Client) AssertionContainsJSONBool(key string, value bool) {
	actual, err := jsonparser.GetBoolean(test.ResponseBody, strings.Split(key, ".")...)
	if err != nil {
		test.t.Errorf("Expected response body contains json key %s with %v, but got Errr(%v)", key, value, err)
	}

	valueStr := "true"
	if !value {
		valueStr = "false"
	}
	actualStr := "true"
	if !actual {
		actualStr = "false"
	}
	assert.EqualValues(test.t, value, actual, "Expected response body contains json key "+key+" with "+valueStr+", but got "+actualStr+".")
}
