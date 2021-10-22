package tfsyntax

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

// get all files filterd with ext.
func CheckExt(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.WalkDir(pathS, func(path string, d os.DirEntry, _ error) error {
		if d.IsDir() && d.Name() == ".terraform" {
			return filepath.SkipDir
		} else {
			r, err := regexp.MatchString(ext, d.Name())
			if err == nil && r {
				files = append(files, d.Name())
			}
		}
		return nil
	})
	return files
}

// return bool if list contains a string
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Create list of keys from a map
func CreateKeyList(except []string, m ...map[string]interface{}) []string {
	var listOfKeys []string
	for _, v := range m {
		for k, _ := range v {
			if !Contains(except, k) {
				listOfKeys = append(listOfKeys, k)
			}

		}
	}
	return listOfKeys
}

// Convert list of interfaces to list of strings.
func ConvertLItoLS(i interface{}) []string {
	var output []string
	if i != nil {
		for _, v := range i.([]interface{}) {
			output = append(output, v.(string))
		}
	}
	return output
}

func Tags(pat, owner, repo string) []string {
	var tags []string
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repoTags, _, err := client.Repositories.ListTags(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range repoTags {
		tags = append(tags, tag.GetName())
	}
	return tags
}

func PopElements(slice, subslice []string) []string {
	for i := 0; i < len(slice); i++ {
		url := slice[i]
		for _, rem := range subslice {
			if url == rem {
				slice = append(slice[:i], slice[i+1:]...)
				i--
				break
			}
		}
	}
	return slice
}

func SupportedVersions(tags, excluded []string) []string {
	return PopElements(tags, excluded)
}
