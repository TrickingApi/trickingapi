<br/>

<div align="center">

<h1>TrickingApi Logo TBD</h1>

circleci badges coming soon

<br/>

</div>

<br/>

A RESTful API for Tricking - [trickingapi.dev](https://trickingapi.dev)

## Setup &nbsp; [![goVersion19](https://img.shields.io/github/go-mod/go-version/TrickingApi/trickingapi)](https://go.dev/doc/tutorial/web-service-gin)
- Download this source code into a working directory

- Install the requirements using:
  ```
  go get .
  ```

- Set up the local development environment using:
  ```
  go build
  ```

- Run the server using the following command:
  ```
  go run main.go
  ```
  
Visit localhost:8080/api/tricks to see the running API! 

Each time you modify main.go or a local package you must run ```go build``` before running again. 
- tbd if hot reloading is necessary/will be added, potentially with [hotswap](https://github.com/edwingeng/hotswap)

## Endpoints

### All Tricks
```http
GET /api/tricks
```
The API endpoints return the JSON representation of the resources requested. However, if an invalid request is submitted, or some other error occurs,
the trickingapi server returns a JSON response in the following format:

```javascript
{
  "message" : string,
  "success" : bool,
  "data"    : string
}
```


## Join Us on Discord!

Have a question or just want to discuss new ideas and improvements? Hit us up on discord.
Consider talking with us here before creating a new issue. 
This way we can keep issues here a bit more organized and helpful in the long run. Be excellent to each other :smile:

[Sign up](https://discord.gg/T588bdSVKU) easily!

## Contributing
Interested in contributing? Great! 
All contributions are welcome: bug fixes, data contributions, recommendations.

Please see the [issues on GitHub](https://github.com/TrickingApi/trickingapi/issues) before you submit a pull request or raise an issue, someone else might have beat you to it.

To contribute to this repository:

- [Fork the project to your own GitHub profile](https://help.github.com/articles/fork-a-repo/)

- Download the forked project using git clone:

    ```sh
    git clone git@github.com:<YOUR_USERNAME>/trickingapi.git
    ```

- Create a new branch with a descriptive name:

    ```sh
    git checkout -b my_new_branch
    ```

- New Feature/Bugfix?
  - Write some code, fix something, and add a test to prove that it works. *No pull request will be accepted without tests passing, or without new tests if new features are added.*

Data Update?
- Update the raw json-files in ```/data```
  - Tricks.json
    - Do not update an existing trick AND add a new trick in the same commit. These should be separate commits to keep things clean :)

- Commit your code and push it to GitHub

- [Open a new pull request](https://help.github.com/articles/creating-a-pull-request/) and describe the changes you have made.

- We'll accept your changes after review.

Simple!
