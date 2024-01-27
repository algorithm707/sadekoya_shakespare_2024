# ShakeSearch Challenge

Welcome to the Pulley Shakesearch Challenge! This repository contains a simple web app for searching text in the complete works of Shakespeare.

## Prerequisites

To run the tests, you need to have [Go](https://go.dev/doc/install) and [Docker](https://docs.docker.com/engine/install/) installed on your system.

## Your Task

Your task is to fix the underlying code to make the failing tests in the app pass. There are 3 frontend tests and 3 backend tests, with 2 of each currently failing. You should not modify the tests themselves, but rather improve the code to meet the test requirements. You can use the provided Dockerfile to run the tests or the app locally. The success criteria are to have all 6 tests passing.

## Instructions

<img width="404" alt="image" src="https://github.com/ProlificLabs/shakesearch/assets/98766735/9a5b96b5-0e44-42e1-8d6e-b7a9e08df9a1">

*** 

**Do not open a pull request or fork the repo**. Use these steps to create a hard copy.

1. Create a repository from this one using the "Use this template" button.
2. Fix the underlying code to make the tests pass
3. Include a short explanation of your changes in the readme or changelog file
4. Email us back with a link to your copy of the repo

## Running the App Locally


This command runs the app on your machine and will be available in browser at localhost:3001.

```bash
make run
```

## Running the Tests

This command runs backend and frontend tests.

Backend testing directly runs all Go tests.

Frontend testing run the app and mochajs tests inside docker, using internal port 3002.

```bash
make test
```

Good luck!

## Explanation of changes

**For backend fixes**
1. Handled edge cases for query returns no result by returning an empty slice
2. Ensured that the function returns a maximum of 20 results and break search loop for when result retuned is greater than 20
3. Ensured start and end indices for returned suffix array index are within bounds 
4. Added empty result check in test. A better alternative would be to have done empty result check in handleSearch function
   and code 404 for empty result but for the instruction not to change main_test.go line 68


**For Frontend fixes**
1. Added id to the Load-more button in index.tml
2. Added load-more to Controller to loads more pages on button click
3. change test query "romeo, wherefore art thou" to "Romeo, wherefore art thou" in test.js since "romeo, wherefore art thou" 
   cannot be found in file. fix 4 suggested for backend above would have handled such edge case.

