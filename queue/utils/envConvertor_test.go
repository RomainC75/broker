package utils

import (
	"fmt"
	testCases "queue/test"
	"testing"
)

func TestToEnvString(t *testing.T) {
	for _, test := range testCases.EnvKeyStringTestCase {
		res := ToEnvString(test.StructKey)
		if res != test.EnvKey {
			t.Errorf("FAIL : key : %s Expected : %s, Actual : %s \n", test.StructKey, test.EnvKey, res)
		}
	}
}

func TestToStructKeyString(t *testing.T) {
	for _, test := range testCases.EnvKeyStringTestCase {
		res := ToStructKeyString(test.EnvKey)
		fmt.Println("=> ", res)
		if res != test.StructKey {
			t.Errorf("FAIL : key : %s Expected : %s, Actual : %s \n", test.EnvKey, test.StructKey, res)
		}
	}
}
