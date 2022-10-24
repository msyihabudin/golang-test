package helper

import (
	"fmt"
	"runtime"
	"testing"

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
	t.Run("Syihab", func(t *testing.T) {
		result := HelloWorld("Syihab")
		require.Equal(t, "Hello Syihab", result, "Result must be 'Hello Syihab'")
	})

	t.Run("Aru", func(t *testing.T) {
		result := HelloWorld("Aru")
		require.Equal(t, "Hello Aru", result, "Result must be 'Hello Aru'")
	})
}

func TestHelloWorld_withtable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Syihab)",
			request:  "Syihab",
			expected: "Hello Syihab",
		},
		{
			name:     "HelloWorld(Aru)",
			request:  "Aru",
			expected: "Hello Aru",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syihab")
	}
}

func BenchmarkHelloWorldAru(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syihab")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Syihab", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syihab")
		}
	})
	b.Run("Aru", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aru")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "Syihab",
			request: "Syihab",
		},
		{
			name:    "Aru",
			request: "Aru",
		},
		{
			name:    "SyihabAru",
			request: "Syihab Dan Aru",
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}
