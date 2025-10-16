package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		level       string
		expectError bool
	}{
		{"debug", false},
		{"info", false},
		{"warn", false},
		{"error", false},
		{"fatal", false},
		{"panic", false},
		{"invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			logger, err := New(tt.level)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				level, err := zerolog.ParseLevel(tt.level)
				if err != nil {
					t.Errorf("failed to parse level: %v", err)
				}

				assert.NoError(t, err)
				assert.NotNil(t, logger)
				assert.Equal(t, zerolog.Level(level), logger.GetLevel())
			}
		})
	}
}
