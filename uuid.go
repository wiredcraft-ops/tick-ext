// TODO: 16/11/15 Linux only
package tick

import "os/exec"

func GetUUID() (uuid []byte, err error) {

	uuid, err = exec.Command("dmidecode", "-s", "system-uuid").Output()
	if err != nil {
		return nil, err
	}

	return uuid, nil
}
