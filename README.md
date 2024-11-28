# Go Project Interface
Go cli tool for creating go projects

This project is not intended to be anyone's standard. This was just a fun way to try to get more familiarity
with golang. What better way than to make a project that will help me setup projects!
I leaned heavily on the content of [Go Project Layout Standards](https://github.com/golang-standards/project-layout)

# Tools
I've only ever build cli tools using viper and cobra. I have no ill words about them but I did want to try some tools
that I had read were more user-friendly. I'm using [urfave/cli](https://github.com/urfave/cli) and [koanf](https://github.com/knadh/koanf)

# Usage
* build: `make build`
* test: `make test`
* run: `make run`
* run to alternate location: `OUTPUT_DIRECTORY=alternative/relative/directory make run`

The run command will default output a project `test` in the examples directory
To run the command on "bare metal" (local) I have alternative entry-points. I would suggest these since I have not
fully documented anything or made a robust makefile yet


# TODO
I would like to actually improve on this significantly as I improve my golang skills. Currently, I have only
ever created golang cli applications so that is what this is initially focused on that. I would like to generate some
good standards and templates for myself to you over time.
I also want to add some ci/cd to generate the binary...just not today

# Links:
* [Go.Dev](https://go.dev/doc/)
* [Go Project Layout Standards](https://github.com/golang-standards/project-layout)
* Cli Tooling [urfave/cli](https://github.com/urfave/cli)
* Configuration Management [koanf](https://github.com/knadh/koanf)
