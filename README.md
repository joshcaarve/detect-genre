# genre-tree
Created a WebApp to visualize music genres and subgenres as a tree structure using React.js and Go.
The WebApp supports adding new genres to the tree and deleting genres or subgenres.


# How it looks
![alt text](demo/genre-tree-demo.gif)

# Notes
If you ever receive a segfault on startup, make sure to update the sys module.
```sh
$ go get -u golang.org/x/sys
$ go mod tidy
```

Fixing npm updates
```sh
$ npm audit fix
```

# How to run
Easiest with two terminals (one for frontend and one for backend)

```sh
$ cd backend
$ go mod tidy
$ rm test.db
$ go run main.go
$ cd ../frontend
$ npm install
$ npm run start
```

# Shortcomings
I'm not a fan of frontend development, so the frontend is, well, quite simple.
The sizing is weird and the buttons are in an awkward spot.
That being said, it was fun to play with React and Typescript.