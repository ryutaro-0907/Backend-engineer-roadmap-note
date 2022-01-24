# Git

## What is Git?
To understand GitHub, you must first have an understanding of Git. Git is an open-source version control system that was started by Linus Torvalds 


![How git works](https://git-scm.com/book/en/v2/images/areas.png)


## Basic commands
```
#clone repository 
git clone {repository/url}

#Branch 
# Check branch (only local)
git branch	

# Check branch (local and remote)
git branch	-a 

# create branch 
git branch {name}

#delte branch 
git branch -d {name}

#Checkout 
#Move to branch, with b create and move.
git checkout {branch name}
git checkout -b  {branch name}	

# Show difference
git diff	

# Show edited files
git status	

# Add file to index 
git add {file paht 1 } {file path 2}...	

# Commit files in index 
git commit	-m 

# Send change to remote server.
git push origin {local branch name}

# Get newset data from remote server.
git fetch origin	

``` 
## .gitignore
The .gitignore file tells Git which files to ignore when committing your project to the GitHub repository. gitignore is located in the root directory of your repo.

```
#Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
#Test binary, built with `go test -c`
*.test
#Output of the go coverage tool, specifically when used with LiteIDE
*.out
#Dependency directories (remove the comment below to include it)
vendor/ 

```
## GitHub
We’ve established that Git is a version control system, similar but better than the many alternatives available. So, what makes GitHub so special? Git is a command-line tool, but the center around which all things involving Git revolve is the hub—GitHub.com—where developers store their projects and network with like minded people.

### Repository 
A repository (usually abbreviated to “repo”) is a location where all the files for a particular project are stored. Each project has its own repo, and you can access it with a unique URL.

### forks 
“Forking” is when you create a new project based off of another project that already exists. This is an amazing feature that vastly encourages the further development of programs and other projects. If you find a project on GitHub that you’d like to contribute to, you can fork the repo, make the changes you’d like, and release the revised project as a new repo. If the original repository that you forked to create your new project gets updated, you can easily add those updates to your current fork

### Pull Requests
You’ve forked a repository, made a great revision to the project, and want it to be recognized by the original developers—maybe even included in the official project/repository. You can do so by creating a pull request. The authors of the original repository can see your work, and then choose whether or not to accept it into the official project. Whenever you issue a pull request, GitHub provides a perfect medium for you and the main project’s maintainer to communicate.


## Reference
1. https://git-scm.com/book/en/v2/Getting-Started-What-is-Git%3F
2. https://www.bmc.com/blogs/gitignore/#:~:text=gitignore%20file%20tells%20Git%20which,is%20a%20plain%20text%20document.