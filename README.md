# Golang Training
 Files associated with [Todd McLeod](https://twitter.com/todd_mcleod) [Go](https://go.dev "The Go programming language") training course on [Udemy](https://www.udemy.com "Main Site").

 ## Index
 [Hands On Exercises][1]
 * [Level 1][2]
 * [Level 2][3]
 * [Level 3][5]

## Useful Documentation
1. [GoDoc](https://godoc.org "Golang documentation organised a bit better than go.dev")
2. [Go By Example](https://gobyexample.com "Golang documentation with working examples")

## TODO : General
Explain why [Todd McLeod](https://twitter.com/todd_mclead)'s [Udemy course](https://www.udemy.com/course/learn-how-to-code) is one of the best I've taken. Ever. For any language.

## TODO : Updates & Releases
* Create bulk file renaming.

 ## Things I've learnt.
 1. When setting up Github, ensure that you've generated a [new SSH key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) and added it to your SSH agent, keychain, and Github user account.
 2. Check your remote repository URLs and change from [HTTPS to SSH](https://docs.github.com/en/get-started/getting-started-with-git/managing-remote-repositories#switching-remote-urls-from-https-to-ssh) if necessary ... if you don't, you'll likely get errors when you try to push.
 3. There are many reasons to love [OhMyZsh!](https://ohmyz.sh/) which offers plugins galore, including the hugely useful [Github short commands](https://kapeli.com/cheat_sheets/Oh-My-Zsh_Git.docset/Contents/Resources/Documents/index).
 4. If you're not using [Homebrew](https://brew.sh/), well ... you really should be.
 5. Spelling mistakes have been the cause of almost **all** of the mistakes I've made ... and, through these mistakes, much of the learning and progress I've made.

## Unfiled
1. [Wesbos](https://twitter.com/wesbos) [Command Line Power User course](https://commandlinepoweruser.com/)
2. [Markdown cheat sheet](https://www.markdownguide.org/cheat-sheet/)
3. [Github does dotfiles](https://dotfiles.github.io/) organisation and [Michael Smalley](https://twitter.com/michaeljsmalley) excellent [blog post](http://blog.smalleycreative.com/tutorials/using-git-and-github-to-manage-your-dotfiles/).
4. [Bazel](https://bazel.build) and [using Golang with Bazel](https://medium.com/@simontoth/golang-with-bazel-2b5310d4ce48)
5. I try not to be a fanboy, but [VScode](https://code.visualstudio.com/) is some next level 'holy shit' IDE stuff.
6. [dotfiles](https://github.com/mpyeager/dotfiles) to help maximise efficiencies with command line and developer workflow.
7. Speaking of dotfiles, you can launch programs directly from terminal using `[program name]` + `.` _or_ `filename`. For example ...
* `code .`, `code [filename]` will launch [VScode][4]
* `open .`, `open [filename]` will launch [Finder](https://support.apple.com/en-gb/HT201732 "Using Finder on Mac")
8. _init_, _condition_, _post_ :-)
9. [diff-so-fancy](https://github.com/so-fancy/diff-so-fancy "Makes `diff` readable without having to pipe to txt file") > Usage `diff -u file1 file 2 | diff-so-fancy`

[1]: https://github.com/mpyeager/GolangTraining/tree/main/Hands%20On%20Exercises
[2]: https://github.com/mpyeager/GolangTraining/tree/main/Hands%20On%20Exercises/Level%201
[3]: https://github.com/mpyeager/GolangTraining/tree/main/Hands%20On%20Exercises/Level%202
[4]: https://code.visualstudio.com
[5]: https://github.com/mpyeager/GolangTraining/tree/main/Hands%20On%20Exercises/Level%203
