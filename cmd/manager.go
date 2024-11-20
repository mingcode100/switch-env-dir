package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

var basePath = ""

func init() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	basePath = current.HomeDir + "\\tools\\Java"
}
func UseJDK(jdkVersion string) {

	mp := findFile()
	//for k, v := range mp {
	//	fmt.Println("k:", k, "v:", v)
	//}
	//return
	v := mp[jdkVersion]
	//fmt.Println("jdkVersion:", jdkVersion)
	if v == "" {
		fmt.Println("JDK Version Not Found")
		os.Exit(0)
	}
	useJavaVersion(v)
}

func useJavaVersion(jdkPath string) {
	targetPath := basePath + "\\current"
	err := os.Remove(targetPath)

	sourcePath := jdkPath
	//fmt.Println("sourcePath:", sourcePath)
	err = os.Symlink(sourcePath, targetPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

func FindJDKs() {
	jdks := findFile()
	for version, path := range jdks {
		fmt.Printf("%-20s %s\n", version, path)
	}
}

func findFile() map[string]string {
	target := "javac.exe"
	jdkMap := make(map[string]string, 0)
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if info.Name() == target {
				dir := filepath.Dir(path)
				dd, _ := filepath.Split(dir)
				cmd := exec.Command(path, "-version")
				versionOutput, err := cmd.CombinedOutput()
				if err != nil {
					panic(err)
				}
				versionInfo := string(versionOutput)
				versionInfo = versionInfo[strings.Index(versionInfo, "javac")+len("javac"):]
				versionInfo = versionInfo[:strings.Index(versionInfo, "\n")]
				//fmt.Println("versionInfo:", versionInfo)
				version := strings.TrimSpace(versionInfo)
				version = strings.Trim(version, "\r\n")
				version = strings.Trim(version, "\n")
				//fmt.Println("version:", version, "dd:", dd)
				jdkMap[version] = dd
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	//return jdks
	return jdkMap
}
