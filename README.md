# Standard Go Project Layout based on Clean Architecture pattern

## Overview

This is a basic layout for Go application projects. It's **`not an official standard defined by the core Go dev team`**; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem. Some of these patterns are more popular than others. It also has a number of small enhancements along with several supporting directories common to any large enough real world application.

**`If you are trying to learn Go or if you are building a PoC or a simple project for yourself this project layout is an overkill. Start with something really simple instead (check [`this`](https://git.zuzex.com/golang/templates/simple) out).`** As your project grows keep in mind that it'll be important to make sure your code is well structured otherwise you'll end up with a messy code with lots of hidden dependencies and global state. When you have more people working on the project you'll need even more structure. That's when it's important to introduce a common way to manage packages/libraries. When you have an open source project or when you know other projects import the code from your project repository that's when it's important to have private (aka `internal`) packages and code. Clone the repository, keep what you need and delete everything else! Just because it's there it doesn't mean you have to use it all. None of these patterns are used in every single project.

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.