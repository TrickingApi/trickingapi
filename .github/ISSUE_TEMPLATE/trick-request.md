---
name: Trick Request
about: Describe the Trick that should be added to the database
title: "[Trick Request] {insert new trick name here}"
labels: new trick
assignees: ''

---

# Fill out the following format with your suggested trick's Info

This template uses markdown styling; please adhere to the [style guidelines](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax)

## Trick Data
For examples, please see [tricks.json](https://github.com/TrickingApi/trickingapi/blob/main/data/tricks.json)

```
{
    "id": "someId", // must be camelcase
    "name": "Some Trick Name", // can have spaces
    "aliases": [], // any aliases for the trick
    "categories": [
      // any of the following: FLIP, VERT_KICK, TWIST, PSEUDO_DUB, SING, DUB, TRIP, QUAD, GROUND_WORK, UNKNOWN
    ],
    "prereqs": [], // any tricks that should be a prerequisite to this trick
    "nextTricks": [], // any tricks that this trick is a prerequisite for
    "description": "" // description of the trick
  }
```

### Has this trick been landed?
- [ ] yes

### Can you provide a video example of this trick?
- [ ] yes

### If yes, please provide a link in this request:
** insert link here **
