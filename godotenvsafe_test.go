package godotenvsafe_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/uulwake/godotenvsafe"
)

const ENV_FILENAME = ".env"
const ENV_TEMPLATE_FILENAME = ".env.template"

func cleanUp() {
	os.Clearenv()
	os.Remove(ENV_FILENAME)
	os.Remove(ENV_TEMPLATE_FILENAME)
}

func writeFile(t *testing.T, data string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		t.Fatal("Cannot create file .env", err)
	}

	defer f.Close()

	_, err = f.Write([]byte(data))
	if err != nil {
		t.Fatal("Cannot write file to .env", err)
	}
}

func TestLoadSuccess(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=bar1
FOO2=bar2 # comment
FOO3=bar3

`

	envTemplateData := `FOO1=
FOO2= # comment 
FOO3=
# FOO4=

`

	writeFile(t, envData, ENV_FILENAME)
	writeFile(t, envTemplateData, ENV_TEMPLATE_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err != nil {
		t.Error("Should successfully load env variable.", err)
	}

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"FOO1 ", "FOO1", "bar1"},
		{"FOO2 ", "FOO2", "bar2"},
		{"FOO3 ", "FOO3", "bar3"},
	}

	godotenvsafe.Load(ENV_FILENAME)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := os.Getenv(tt.input)
			if got != tt.want {
				fmt.Printf("Got: %s. Want: %s", got, tt.want)
			}
		})
	}
}

func TestLoadFailInvalidEnv(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=""bar#1""
FOO2=bar2 # comment
FOO3=bar3
`
	writeFile(t, envData, ENV_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err == nil {
		t.Error("Should fail loading env variable.", err)
	}
}

func TestLoadFailInvalidEnvTemplate(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=bar1
FOO2=bar2 # comment
FOO3=bar3
`
	writeFile(t, envData, ENV_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err == nil {
		t.Error("Should fail loading env variable template.", err)
	}
}

func TestLoadFailInvalidFormatEnvTemplate(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=bar1
FOO2=bar2 # comment
FOO3=bar3
`

	envTemplateData := `FOO1=
FOO2= # comment 
FOO3
# FOO4=
`

	writeFile(t, envData, ENV_FILENAME)
	writeFile(t, envTemplateData, ENV_TEMPLATE_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err == nil {
		t.Error("Should fail load env variable template [invalid format].", err)
	}
}

func TestLoadFailMissingOneEnv(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=bar1
FOO2=bar2 # comment
FOO3=bar3
`

	envTemplateData := `FOO1=
FOO2= # comment 
FOO3=
FOO4=
`

	writeFile(t, envData, ENV_FILENAME)
	writeFile(t, envTemplateData, ENV_TEMPLATE_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err == nil {
		t.Error("Should fail: missing one env", err)
	}

	gotErrMsg := err.Error()
	wantErrMsg := "there is 1 missing environment variable: FOO4"
	if gotErrMsg != wantErrMsg {
		t.Errorf("Got error message: %s. Want error message: %s", gotErrMsg, wantErrMsg)
	}
}

func TestLoadFailMissingMultipleEnv(t *testing.T) {
	defer cleanUp()

	envData := `FOO1=bar1
FOO2=bar2 # comment
FOO3=bar3
`

	envTemplateData := `FOO1=
FOO2= # comment 
FOO3=
FOO4=
FOO5=
# FOO6=
`

	writeFile(t, envData, ENV_FILENAME)
	writeFile(t, envTemplateData, ENV_TEMPLATE_FILENAME)

	err := godotenvsafe.Load(ENV_FILENAME)
	if err == nil {
		t.Error("Should fail: missing multiple envs", err)
	}

	gotErrMsg := err.Error()
	wantErrMsg := "there are 2 missing environment variables: FOO4,FOO5"
	if gotErrMsg != wantErrMsg {
		t.Errorf("Got error message: %s. Want error message: %s", gotErrMsg, wantErrMsg)
	}
}
