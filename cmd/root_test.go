package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		stdin   string
		wantOut string
		wantErr bool
	}{
		{
			name:    "char count from args",
			args:    []string{"--chars", "hello"},
			stdin:   "",
			wantOut: "hello\nCharacter count: 5\n",
			wantErr: false,
		},
		{
			name:    "uppercase from args",
			args:    []string{"--upper", "hEllO", "WoRlD"},
			stdin:   "",
			wantOut: "HELLO WORLD\n",
			wantErr: false,
		},
		{
			name:    "lowercase from args",
			args:    []string{"--lower", "hEllO", "WoRlD"},
			stdin:   "",
			wantOut: "hello world\n",
			wantErr: false,
		},
		{
			name:    "word count from args",
			args:    []string{"--words", "three blind mice"},
			stdin:   "",
			wantOut: "three blind mice\nWord count: 3\n",
			wantErr: false,
		},
		{
			name:    "uppercase takes precedence over lowercase",
			args:    []string{"--upper", "--lower", "hEllO", "WoRlD"},
			stdin:   "",
			wantOut: "HELLO WORLD\n",
			wantErr: false,
		},
		{
			name:    "uppercase from stdin",
			args:    []string{"--upper"},
			stdin:   "hello from stdin",
			wantOut: "HELLO FROM STDIN\n",
			wantErr: false,
		},
		{
			name:    "no input",
			args:    []string{},
			stdin:   "",
			wantOut: "",
			wantErr: true,
		},
		{
			name:    "char and word count from args",
			args:    []string{"--chars", "--words", "three blind mice"},
			stdin:   "",
			wantOut: "three blind mice\nCharacter count: 16\nWord count: 3\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()

			outBuf := new(bytes.Buffer)
			rootCmd.SetOut(outBuf)
			rootCmd.SetErr(io.Discard)
			defer rootCmd.SetOut(os.Stdout)
			defer rootCmd.SetErr(os.Stderr)

			if tt.stdin != "" {
				rootCmd.SetIn(strings.NewReader(tt.stdin))
				defer rootCmd.SetIn(os.Stdin)
			}

			rootCmd.SetArgs(tt.args)
			err := Execute()
			gotOut := outBuf.String()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error, got: %q", gotOut)
				}
				return
			}

			if err != nil {
				t.Errorf("Execute() got error: %v, want: %q", err, tt.wantOut)
				return
			}

			if gotOut != tt.wantOut {
				t.Errorf("Execute() got: %q, want: %q", gotOut, tt.wantOut)
			}
		})
	}
}

func resetFlags() {
	chars = false
	lower = false
	upper = false
	words = false
}
