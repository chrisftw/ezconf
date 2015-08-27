package ezconf

import (
	"testing"
)

func TestSampleSetAndGet(t *testing.T) {
	sampleNothing := Get("sample", "nothing")
	AssertTrue(sampleNothing == "", "Nonexistent settings did not return as empty strings", t)
	Set("sample", "something", "something important.")
	sampleSomething := Get("sample", "something")
	AssertTrue(sampleSomething == "something important.", "Settings did not save or return.", t)
}

func TestSampleReadsFromFile(t *testing.T) {
	sampleNothing := Get("sample", "nothing")
	AssertTrue(sampleNothing == "", "Nonexistent settings did not return as empty strings", t)
	sampleSomething := Get("sample", "sample_1")
	AssertTrue(sampleSomething == "SAMPLE!!!", "Settings from file were not correct.1", t)
	sampleSomething2 := Get("sample", "sample_2")
	AssertTrue(sampleSomething2 == "TEST_2 setting here", "Settings from file were not correct.2", t)
	sampleSomething3 := Get("sample", "sample_3")
	AssertTrue(sampleSomething3 == "value with \"quotes\"", "Settings from file were not correct.3", t)
	sampleSomething4 := Get("sample", "sample4")
	AssertTrue(sampleSomething4 == "space before the colon", "Settings from file were not correct.4", t)
	sampleSomething5 := Get("sample", "sample5-url")
	AssertTrue(sampleSomething5 == "http://chrisreister.com", "Settings from file were not correct.5", t)
}

func TestDifferentTypes(t *testing.T) {
	testString := Get("test", "test-string")
	AssertTrue(testString == "bama bama alabama", "Did not correctly read a string value", t)
	testFloat := GetFloat("test", "test-float")
	AssertTrue(testFloat == 3.14, "Did not correctly read a float", t)
	testInt := GetInt("test", "test-int")
	AssertTrue(testInt == -42, "Did not correctly read an int", t)
	testUint := GetUint("test", "test-uint")
	AssertTrue(testUint == 23, "Did not correctly read an uint", t)
	testBool1 := GetBool("test", "test-bool")
	AssertTrue(testBool1 == true, "Did not correctly read a true bool", t)
	testBool2 := GetBool("test", "test-bool-2")
	AssertTrue(testBool2 == false, "Did not correctly read a false bool", t)
}

func BenchmarkReadingMissingConfigs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("sample", "nothing")
	}
}

func BenchmarkReadingRealConfigs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("sample", "sample4")
	}
}

func AssertTrue(value bool, message string, t *testing.T) {
	if !value {
		t.Log(message)
		t.Fail()
	}
}
