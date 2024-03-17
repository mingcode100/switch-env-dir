package main

import (
	"flag"
	"fmt"
	"os"
)

var basePath = "C:\\Users\\liuyu\\tools\\Java"
var versions = make(map[string]string)

func init() {
	versions["1.8"] = "jdk1.8.0_341"
	versions["11"] = "jdk-11"
	versions["20"] = "jdk-20.0.1"
	versions["21"] = "graalvm-jdk-21+35.1"
}

func main() {
	var jdkVersion string
	flag.StringVar(&jdkVersion, "v", "1.8", "target JDK version")
	flag.Parse()
	fmt.Println(jdkVersion)

	useJavaVersion(basePath, jdkVersion)
}

func useJavaVersion(basePath string, jdkVersion string) {
	targetPath := basePath + "\\current"
	err := os.Remove(targetPath)

	jdkPath := versions[jdkVersion]
	err = os.Symlink(basePath+"\\"+jdkPath, targetPath)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Switch To JDK ", jdkVersion)
}
