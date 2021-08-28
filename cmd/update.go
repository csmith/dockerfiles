package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/csmith/dockerfiles"
	"github.com/kouhin/envflag"
)

var (
	shouldCommit = flag.Bool("commit", true, "Whether or not to commit changes")
	shouldBuild  = flag.Bool("build", true, "Whether or not to build (and push) the image")
	target       = flag.String("target", "", "Project to build. Blank for all projects.")

	funcs = template.FuncMap{
		"image":           Image,
		"alpine_packages": AlpinePackages,
		"arch_packages":   ArchPackages,
		"github_tag":      GitHubTag,
	}
)

func main() {
	if err := envflag.Parse(); err != nil {
		panic(err)
	}

	if *target != "" {
		Update(*target)
		return
	}

	// NB: These are manually sorted to flatten the dependency hierarchy.
	targets := []string{
		"alpine",
		"arch",
		"base",             // depends on alpine
		"golang",           // depends on alpine
		"distribution",     // depends on golang + base
		"irc-bot",          // depends on golang + base
		"irc-distribution", // depends on golang + base
		"irc-github",       // depends on golang + base
		"irc-goplum",       // depends on golang + base
		"irc-news",         // depends on golang + base
		"postgres-13",      // depends on alpine
		// "vault",        // depends on alpine + golang
	}

	for i := range targets {
		Update(targets[i])
	}
}

var materials map[string]string

func Update(dir string) {
	materials = make(map[string]string)

	outputPath := filepath.Join(dir, "Dockerfile")
	inputPath := filepath.Join(dir, "Dockerfile.gotpl")
	oldMaterials := existingBom(outputPath)

	tpl := template.New(inputPath)
	tpl.Funcs(funcs)

	if _, err := tpl.ParseFiles(inputPath); err != nil {
		log.Fatalf("Unable to parse template file %s: %v", inputPath, err)
	}

	writer := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(writer, "Dockerfile.gotpl", nil); err != nil {
		log.Fatalf("Unable to render template file %s: %v", inputPath, err)
	}

	bom, _ := json.Marshal(materials)
	header := fmt.Sprintf("# Generated from https://github.com/csmith/dockerfiles/blob/master/%s\n# BOM: %s\n\n", inputPath, bom)

	content := append([]byte(header), writer.Bytes()...)
	if err := os.WriteFile(outputPath, content, os.FileMode(0600)); err != nil {
		log.Fatalf("Unable to write Dockerfile to %s: %v", outputPath, err)
	}

	if *shouldCommit {
		gitCommand := exec.Command(
			"git",
			"commit",
			"--no-gpg-sign",
			"-m",
			fmt.Sprintf("[%s] %s", dir, diff(oldMaterials, materials)),
			outputPath,
		)
		gitCommand.Stdout = os.Stdout
		gitCommand.Stderr = os.Stderr
		if err := gitCommand.Run(); err != nil {
			// Either the commit failed, or there was nothing to commit. Either way, we don't want to build.
			log.Printf("Failed to run git commit: %v", err)
			return
		}
	}

	if *shouldBuild {
		imageName := fmt.Sprintf("reg.c5h.io/%s", dir)
		buildahCommand := exec.Command(
			"/usr/bin/buildah",
			"bud",
			"--timestamp",
			"0",
			"--tag",
			imageName,
			dir,
		)
		buildahCommand.Stdout = os.Stdout
		buildahCommand.Stderr = os.Stderr
		if err := buildahCommand.Run(); err != nil {
			log.Fatalf("Failure building image: %v", err)
		}

		pushCommand := exec.Command(
			"/usr/bin/buildah",
			"push",
			imageName,
		)
		pushCommand.Stdout = os.Stdout
		pushCommand.Stderr = os.Stderr
		if err := pushCommand.Run(); err != nil {
			log.Fatalf("Failure pushing image: %v", err)
		}
	}
}

func diff(oldBom, newBom map[string]string) string {
	var res []string
	for i := range newBom {
		if oldBom[i] != newBom[i] {
			if len(newBom[i]) > 10 || oldBom[i] == "" {
				res = append(res, fmt.Sprintf("%s updated", i))
			} else {
				res = append(res, fmt.Sprintf("%s: %s->%s", i, oldBom[i], newBom[i]))
			}
		}
	}
	if len(res) == 0 {
		return "no detected changes"
	}
	return strings.Join(res, ", ")
}

func existingBom(target string) map[string]string {
	res := make(map[string]string)
	bs, err := os.ReadFile(target)
	if err != nil {
		log.Printf("Unable to read existing Dockerfile (%s) for BOM: %v", target, err)
		return res
	}

	bomLine := strings.SplitN(string(bs), "\n", 3)[1]
	if !strings.HasPrefix(bomLine, "# BOM: ") {
		log.Printf("Existing Dockerfile (%s) does not appear to have a BOM", target)
		return res
	}

	if err := json.Unmarshal([]byte(strings.TrimPrefix(bomLine, "# BOM: ")), &res); err != nil {
		log.Printf("Existing Dockerfile (%s) has invalid BOM: %v", target, err)
		return res
	}

	return res
}

func Image(ref string) string {
	res, err := dockerfiles.LatestDigest(ref)
	if err != nil {
		log.Fatalf("Unable to get latest digest for ref %s: %v", ref, err)
	}
	materials[fmt.Sprintf("image:%s", ref)] = res
	return fmt.Sprintf("%s@%s", ref, res)
}

func AlpinePackages(packages ...string) map[string]string {
	res, err := dockerfiles.LatestAlpinePackages(packages...)
	if err != nil {
		log.Fatalf("Unable to get latest packages: %v", err)
	}
	for i := range res {
		materials[fmt.Sprintf("apk:%s", i)] = res[i]
	}
	return res
}

func ArchPackages(packages ...string) map[string]string {
	res, err := dockerfiles.LatestArchPackages(packages...)
	if err != nil {
		log.Fatalf("Unable to get latest packages: %v", err)
	}
	for i := range res {
		materials[fmt.Sprintf("pacman:%s", i)] = res[i]
	}
	return res
}

func GitHubTag(repo string) string {
	tag, err := dockerfiles.LatestGitHubTag(repo)
	if err != nil {
		log.Fatalf("Couldn't determine latest tag: %v", err)
	}
	materials[fmt.Sprintf("github:%s", repo)] = tag
	return tag
}
