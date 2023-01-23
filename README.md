<br/>

<div align="center">

<h1>TrickingApi Logo TBD</h1>

[![build status](https://img.shields.io/circleci/build/github/TrickingApi/trickingapi)](https://circleci.com/gh/TrickingApi/trickingapi)
![issues](https://img.shields.io/github/issues/TrickingApi/trickingapi)
![prs](https://img.shields.io/github/issues-pr/TrickingApi/trickingapi)
![repo size](https://img.shields.io/github/repo-size/TrickingApi/trickingapi)
![go version](https://img.shields.io/github/go-mod/go-version/TrickingApi/trickingapi)
[![license](https://img.shields.io/github/license/TrickingApi/trickingapi)](https://github.com/TrickingApi/trickingapi/blob/main/LICENSE.md)
[![discord](https://img.shields.io/discord/1061481749894418533)](https://discord.gg/7r99xBX6eU)

<br/>

</div>

<br/>

A RESTful API for Tricking - [trickingapi.dev](https://trickingapi.dev)

Direct API access - [api.trickingapi.dev](https://api.trickingapi.dev)

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
  
- Run tests to verify all existing tests still work:
  ```
  go test -v ./...
  ```

- Run the server using the following command:
  ```
  go run main.go
  ```
  
- Update API Documentation
  ```
  swag init
  ```
  
Visit localhost:8080/tricks to see the running API! 

Each time you modify main.go or a local package you must run ```go build``` before running again. 
- tbd if hot reloading is necessary/will be added, potentially with [hotswap](https://github.com/edwingeng/hotswap)

## API Routes

### All Tricks
```http
GET /tricks
```

### Get a Specific Trick
```http
GET /tricks/:name
```

e.g.

```http
GET /tricks/pop360
```
```json
{
    "id": "pop360",
    "name": "Pop 360",
    "categories": [
        "Vert Kick"
    ],
    "difficultRank": 0,
    "prerequisites": null,
    "nextTricks": [
        "Pop 360 Shuriken",
        "Pop 720",
        "Swing 360",
        "Illusion Twist"
    ],
    "description": "The Pop 360 starts in Frontside stance, leaves the ground off of both feet, rotates 180° in the air, then lands and finishes with an outside crescent kick towards the target and lands in turbo (both feet). This trick is also a hyper, because Pop 180 Hook (TKT) is not often used in tricking. This trick is vitally important for Illusion Twist and other similar tricks that end in turbo."
}
```

### Get All Category Names
```http
GET /categories
```

### Get Tricks Separated by Categories
```http
GET /categories/tricks
```

### Get All Tricks for a Specified Category
```http
GET /categories/:name
```

e.g.
```http
GET /categories/Vert%20Kick
```

```json
[
    {
        "id": "pop360",
        "name": "Pop 360",
        "categories": [
            "Vert Kick"
        ],
        "prereqs": [],
        "nextTricks": [
            "Pop 360 Shuriken",
            "Pop 720",
            "Swing 360",
            "Illusion Twist"
        ],
        "description": "The Pop 360 starts in Frontside stance, leaves the ground off of both feet, rotates 180° in the air, then lands and finishes with an outside crescent kick towards the target and lands in turbo (both feet). This trick is also a hyper, because Pop 180 Hook (TKT) is not often used in tricking. This trick is vitally important for Illusion Twist and other similar tricks that end in turbo."
    },
    {
        "id": "pop360Shuriken",
        "name": "Pop 360 Shuriken",
        "categories": [
            "Vert Kick"
        ],
        "prereqs": [
            "Pop 360"
        ],
        "nextTricks": [
            "Pop 720",
            "Pop 720 Double",
            "Swing 360 Shuriken"
        ],
        "description": ""
    }
]
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
  - For new API endpoints make sure to do the following
    - [ ] Add declarative comments above request handlers following https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
    - [ ] Add unit tests for the respective route
    - [ ] Run swag init for updating documentation
    - PRs will not be approved for new endpoints without completing the three tasks above ^^^


Data Update?
- Create an issue under the "new trick" label and fill out the template and assign it to yourself (if you intend to make the change yourself)

- Update the raw json-files in ```/data``` with your desired change
  - Tricks.json
     - Do not update an existing trick AND add a new trick in the same commit. These should be separate commits to keep things clean :)

- Commit your code and push it to GitHub

- [Open a new pull request](https://help.github.com/articles/creating-a-pull-request/) and describe the changes you have made.

- We'll accept your changes after review.

Simple!

## Questions
If you have any questions, create an [issue](issue) (protip: do a quick search first to see if someone else didn't ask the same question before!).
You can also reach the owner/developers at our [discord](https://discord.gg/T588bdSVKU)
