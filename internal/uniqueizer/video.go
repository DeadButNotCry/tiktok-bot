package uniqueizer

import (
	"os"
	"os/exec"
	"strings"
)

func DoUnique(pathIn, pathOut string) error {
	command := "ffmpeg -i " + pathIn + " -map_metadata -1 -vf noise=alls=1:allf=t " + pathOut
	parts := strings.Fields(command)

	//for _, part := range parts {
	//	log.Println(part)
	//}
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return nil
}

//  ffmpeg -i videos/BAACAgIAAxkBAANHZBx1_nvrqsKxvlGHKEpyDrgGM04AAngrAAIKkelIN_V748UewRMvBA.mp4 -map_metadata -1 ./result/BAACAgIAAxkBAANHZBx1_nvrqsKxvlGHKEpyDrgGM04AAngrAAIKkelIN_V748UewRMvBA.mp4
