package helper

import (
	"testing"
	"fmt"
	"runtime"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld_success(t *testing.T) {
	result := HelloWorld("Syihab")
	if result != "Hello Syihab" {
		t.Fail()
	}

	fmt.Println("TestHelloWorld_success done")
}

func TestHelloWorld_fail(t *testing.T) {
	result := HelloWorld("Syihabudin")
	if result != "Hello Syihab" {
		t.Error("Result must be 'Hello Syihab'")
	}

	fmt.Println("TestHelloWorld_fail done")
}

func TestHelloWorld_withassert(t *testing.T) {
	result := HelloWorld("Syihab")
	assert.Equal(t, "Hello Syihab", result, "Result must be 'Hello Syihab'")
	fmt.Println("TestHelloWorld_withassert done")
}

func TestHelloWorld_withrequire(t *testing.T) {
	result := HelloWorld("Syihab")
	require.Equal(t, "Hello Syihab", result, "Result must be 'Hello Syihab'")
	fmt.Println("TestHelloWorld_withrequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}

	result := HelloWorld("Syihab")
	require.Equal(t, "Hello Syihab", result, "Result must be 'Hello Syihab'")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")

	m.Run()

	//after
	fmt.Println("After unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Syihab", func(t *testing.T){
		result := HelloWorld("Syihab")
		require.Equal(t, "Hello Syihab", result, "Result must be 'Hello Syihab'")
	})

	t.Run("Aru", func(t *testing.T){
		result := HelloWorld("Aru")
		require.Equal(t, "Hi Aru", result, "Result must be 'Hello Aru'")
	})
}