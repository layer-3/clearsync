# Commit practices

This document explains the importance and benefits of commit practices and the format of the commit message.

> Inspired by and following [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

If you have any questions, open the link above.

## Why commit practices are important?

Commit practices are crucial in software development as they ensure a clear and organized history of changes, facilitating collaboration among team members.
By making frequent, well-documented commits, developers can easily track progress, identify and resolve bugs, and understand the evolution of the codebase.
Good commit practices enhance code quality and maintainability, enabling efficient version control and seamless integration of contributions.

### Doesn’t this discourage rapid development and fast iteration?

It discourages moving fast in a disorganized way. It helps you be able to move fast long term across multiple projects with varied contributors.

## What are the benefits

- Communicate the nature of changes to teammates, the public, and other stakeholders.
- Make it easier for people to contribute to your projects, by allowing them to explore a more structured commit history.
- Automatically generate CHANGELOGs.
- Automatically determine a semantic version bump (based on the types of commits landed).
- Trigger build and publish processes.

## Commit format

The commit message should be structured as follows:

```txt
<type>[optional scope][optional "!"]: <description>

[optional body]

[optional footer(s)]
```

### Type

The type of commit message is a required parameter and should be one of the following:

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refact**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding tests or correcting existing tests
- **build**: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- **ci**: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
- **chore**: Other changes that don't modify src or test files
- **revert**: Reverts a previous commit

### Optional scope

The scope provides additional contextual information about the commit and is enclosed in brackets. It is optional and can be omitted if the commit does not apply to a specific component.

If it does apply, the scope should not take more than 18 characters. If it takes more, that means the scope you are trying to describe is too detailed. Choose a higher level scope instead.

A scope can include a name of a component and its subcomponent, separated by a slash. A subcomponent is usually the lowest level of component that is being modified.

Example components: `UI`, `API`, `Models`, `Hooks`

Example subcomponents: `tx modal`, `estimate`, `user`, `useFetch`

A component and subcomponent should be separated by a slash `/`, e.g. `UI/tx modal`.

If this commit introduces a component or a subcomponent for the first time, then it should not be mentined in the scope, but in the description.
For example, if you are adding a new component `UI/tx modal`, then the scope should be omitted, and the commit message should look like this:

```txt
feat: add tx modal component
```

Then, when modifying this component, you must use the scope `UI/tx modal`.

```txt
fix(UI/tx modal): not closing on click
```

### Description

The description is a brief summary of the changes made in the commit. It should be clear, concise, and written in the imperative style. The imperative style is used to describe _what would happend if the commit is applied to the codebase_.

The type, scope and description can not exceed 50 charecters. Therefore, the description should be short and to the point.

If the description is too long, it should be split into the body.

### Optional "!" / Breaking change

**BREAKING CHANGE**: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with [MAJOR](https://semver.org/#summary) in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.

### Optional body

The body provides additional information about the changes made in the commit. It is optional and can be omitted if the description is sufficient.

#### Example of a good body

Free-form description:
```txt
feat: add tx modal component

Add a new component for displaying transaction details. The component includes a modal window with a close button and a list of transaction details.
```

A breakdown of larger commit:
```txt
refact(api/cache): improve caching system

- Optimize cache key generation to avoid collisions
- Refactor cache invalidation to account for edge cases
- Add tests for cache hit/miss scenarios
```

### Footer(s)

The footer is used to reference issues, provide additional context, or link to external resources. It is optional and can be omitted if not needed.

**BREAKING CHANGE**: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with [MAJOR](https://semver.org/#summary) in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.

_footers_ other than BREAKING CHANGE: \<description\> may be provided and follow a convention similar to [git trailer format](https://git-scm.com/docs/git-interpret-trailers).

## Enforcing the practice

### Developer tools

The commitizen is used to enforce the commit message format. It is a command-line utility that prompts users to fill in the required fields and generates a commit message based on the input.

Each repository integrates the commitizen tool as an `npm` development dependency and configures it to use the conventional commit format.

This repository contains the shared configuration for commitizen, which is used by all projects in the organization.

### CI/CD pipeline

The CI/CD pipeline is configured to run a pre-commit hook that checks the commit message format before allowing the pipeline to proceed.

Using such a mechanism ensures that all commits in the repository follow the commit practices and would allow us to automatically compile a changelog, determine a semantic version bump, and trigger build and publish processes in other repositories that depend on a given one (e.g. `github.com/layer-3/clearsync` triggers `github.com/layer-3/neodax` bump dependency version, which in turn adds a changelog entries for both repos).
