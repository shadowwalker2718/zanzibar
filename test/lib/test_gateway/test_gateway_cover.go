// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package testGateway

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	yaml "gopkg.in/yaml.v2"

	config "github.com/uber/zanzibar/examples/example-gateway/config"
)

var cachedBinaryFile *testBinaryInfo

type testBinaryInfo struct {
	binaryFile       string
	coverProfileFile string
}

func (info *testBinaryInfo) Cleanup() {
	if os.Getenv("COVER_ON") != "1" {
		return
	}

	randStr, err := makeRandStr()
	if err != nil {
		panic(err)
	}

	newCoverProfileFile := path.Join(
		path.Dir(info.coverProfileFile),
		"cover-"+randStr+".out",
	)

	bytes, err := ioutil.ReadFile(info.coverProfileFile)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(newCoverProfileFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func writeConfigToFile(config *config.Config) (string, error) {
	tempConfigDir, err := ioutil.TempDir("", "example-gateway-config-yaml")
	if err != nil {
		return "", err
	}

	baseYamlFile := path.Join(tempConfigDir, "production.yaml")
	configBytes, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(baseYamlFile, configBytes, os.ModePerm)
	if err != nil {
		return "", err
	}

	return tempConfigDir, nil
}

func makeRandStr() (string, error) {
	randBytes := make([]byte, 8)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}
	randStr := fmt.Sprintf("%x", randBytes)

	return randStr, nil
}

func createTestBinaryFile(config *config.Config) (*testBinaryInfo, error) {
	if cachedBinaryFile != nil {
		return cachedBinaryFile, nil
	}

	dirName := getProjectDir()
	mainTestPath := path.Join(dirName, "test", "child_process", "start_gateway_test.go")

	randStr, err := makeRandStr()
	if err != nil {
		return nil, err
	}

	coverProfileFile := path.Join(dirName, "coverage", "cover-"+randStr+".out")

	tempConfigDir, err := ioutil.TempDir("", "example-gateway-test-binary")
	if err != nil {
		return nil, err
	}

	binaryFile := path.Join(tempConfigDir, "test-"+randStr+".test")

	novendorCmd := exec.Command("glide", "novendor")
	novendorCmd.Dir = dirName
	outBytes, err := novendorCmd.Output()
	if err != nil {
		return nil, err
	}

	allFolders := strings.Split(string(outBytes), "\n")
	allPackages := []string{}
	for _, folder := range allFolders {
		if folder == "./test/..." ||
			folder == "./main/..." ||
			folder == "./benchmarks/..." ||
			folder == "" {
			continue
		}

		allPackages = append(allPackages, path.Join("github.com/uber/zanzibar", folder))
	}

	allPackagesString := strings.Join(allPackages, ",")

	args := []string{
		"-c", "0", "go", "test", "-v",
		"-c", "-o", binaryFile,
	}
	if os.Getenv("COVER_ON") == "1" {
		args = append(args,
			"-cover", "-coverprofile", coverProfileFile,
			"-coverpkg", allPackagesString)
	}

	args = append(args, mainTestPath)

	var testGenCmd *exec.Cmd
	if runtime.GOOS == "linux" {
		testGenCmd = exec.Command("taskset", args...)
	} else {
		testGenCmd = exec.Command(args[2], args[3:]...)
	}

	// testGenCmd.Stderr = os.Stderr
	// testGenCmd.Stdout = os.Stdout
	err = testGenCmd.Run()
	if err != nil {
		return nil, err
	}

	cachedBinaryFile = &testBinaryInfo{binaryFile, coverProfileFile}
	return cachedBinaryFile, nil
}
