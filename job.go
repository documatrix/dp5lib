package dp5lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// SetJobVar can be called to set the job variable of the given job.
// If id is nil, the variables will be set for the current job (if there
// is a current job).
func SetJobVar(id *uint64, key string, value string) error {
	dp5SetJobVar := "dp5_set_job_var"

	// Get the dp5_set_job_var executable
	dmBin := os.Getenv("DM_BIN")
	if dmBin != "" {
		dp5SetJobVar = filepath.Join(dmBin, dp5SetJobVar)
	}

	params := []string{
		"-key",
		key,
		"-value",
		value,
	}

	if id != nil {
		params = append(params, "-job-id", strconv.FormatUint(*id, 10))
	}

	cmd := exec.Command(dp5SetJobVar, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error while setting job variable %s to %s using command %s! %s", key, value, cmd.String(), err)
	}

	return err
}

// SetJobVars can be called to set multiple job variables of the given job.
// If id is nil, the variables will be set for the current job (if there
// is a current job).
func SetJobVars(id *uint64, vars map[string]string) error {
	for k, v := range vars {
		err := SetJobVar(id, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
