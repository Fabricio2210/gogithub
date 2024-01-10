package gogit

import(
	"fmt"
	"os"
	"github.com/go-git/go-git/v5"
)

func Gogit (repositoryURL string, destinationDir string) {
	_, err := os.Stat(destinationDir)
	if os.IsNotExist(err) {
		fmt.Println("Cloning repository ", repositoryURL)
		_, err := git.PlainClone(destinationDir, false, &git.CloneOptions{
			URL:      repositoryURL,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Repository cloned successfully!")
	} else {
		fmt.Println("Repository already exists. Performing pull...")
		repo, err := git.PlainOpen(destinationDir)
		if err != nil {
			fmt.Println("Error opening repository:", err)
			return
		}
		worktree, err := repo.Worktree()
		if err != nil {
			fmt.Println("Error getting worktree:", err)
			return
		}

		fmt.Println("Fetching from remote repository ", repositoryURL)
		err = repo.Fetch(&git.FetchOptions{
			Progress: os.Stdout,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println("Error fetching updates:", err)
			return
		}

		fmt.Println("Merging fetched changes...")
		err = worktree.Pull(&git.PullOptions{
			RemoteName: "origin",
			Progress:   os.Stdout,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println("Error pulling updates:", err)
			return
		}

		fmt.Println("Pull operation completed successfully!")
	}
}