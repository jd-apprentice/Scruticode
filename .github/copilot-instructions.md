---
applyTo: "**"
---

## Code Review

When performing a code review, respond always in english and check for my grammar.
When performing a code review, follow this file for any standard.

## Coding

Do not use if else statements in general, prefer early returns or negate conditions.
Reduce the amount of magic strings and numbers, use constants instead.
Ensure values that are being used more than one time to be stored in a variable with a descriptive name.

## Testing

For each new file created, if there is a public function
 - Ensure to create a test file with the same name and the suffix `_test.go`.
 - Include at least one test for the happy path and one for the sad path.

## Quality standards

Ensure that the code is following the guidelines from the `.golangci.yml` file.

## Commands

Ensure that each new command used if is repeated more than once it should be added to the `Makefile` at the root of the project.

## Project related

Since the project is used to check quality and standards in other projects, all mocks should be located in the `mocks` folder, which by default should be ignored.
Inside that folder, mocks should be separated by module or functionality.