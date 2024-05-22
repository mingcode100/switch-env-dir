package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var basePath = "C:\\Users\\liuyu\\tools\\Java"

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
			if strings.Contains(path, target) {
				dir := filepath.Dir(path)
				dd, _ := filepath.Split(dir)
				cmd := exec.Command(path, "-version")
				versionOutput, err := cmd.CombinedOutput()
				if err != nil {
					panic(err)
				}
				versionInfo := string(versionOutput)
				versionInfo = strings.ReplaceAll(versionInfo, "javac", "")
				versionInfo = strings.TrimSpace(versionInfo)
				versionInfo = strings.Trim(versionInfo, "\n")
				versionInfo = strings.Trim(versionInfo, "\r\n")

				split := strings.Split(versionInfo, "\n")
				var version = ""
				if len(split) > 1 {
					version = split[len(split)-1]
				} else {
					version = split[0]
				}
				version = strings.TrimSpace(version)
				jdkMap[version] = dd

			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	//return jdks
	return jdkMap
}
