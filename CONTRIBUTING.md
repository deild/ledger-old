# Contributing to Ledger

Want to help contribute to Ledger? Great! Any and all help is certainly appreciated, whether it's code, documentation, or spelling corrections.


## Filing issues

When filing an issue, make sure to answer these five questions:

1. What did you do?
2. What did you expect to see?
3. What did you see instead?
4. What version of Ledger are you using (`ledger --version`)?
5. What version of Go are you using? (`go version`)?

## Contributing code

Make sure your local environment has the following software installed:

* [Git](https://git-scm.com/)
* [Go](https://golang.org) 1.5+
* [Glide](https://github.com/Masterminds/glide)

Fork this repo and create your own feature branch. We encourage pull requests to discuss code changes.

### Guidelines

Consider the following guidelines when preparing to submit a patch:

* Follow standard Go conventions (document any new exported types, funcs, etc.., ensuring proper punctuation).
* Ensure that you test your code. Any patches sent in for new / fixed functionality must include tests in order to be merged into master.
* If you plan on making any major changes, create an issue before sending a patch. This will allow for proper discussion beforehand.

### Format of the Commit Message

We follow a rough convention for commit messages that is designed to answer two
questions: what changed and why. The subject line should feature the what and
the body of the commit should describe the why.

```
Add 'list' command

This change adds the 'list' command which will print either your current
location, or the location that you specify using the Delve linespec.

Fixes #38
```

The format can be described more formally as follows:

```
<what changed>
<BLANK LINE>
<why this change was made>
<BLANK LINE>
<footer>
```

The first line is the subject and should be no longer than 70 characters, the
second line is always blank, and other lines should be wrapped at 80 characters.
This allows the message to be easier to read on GitHub as well as in various
git tools.

This [blog article](http://chris.beams.io/posts/git-commit/) is a good resource for learning how to write good commit messages