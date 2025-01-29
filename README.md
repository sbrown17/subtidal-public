# Prototype
Hello, this is a prototype of some grander vision of a message board I had before I stopped working on it for a few reasons. One of the biggest reasons being I didn't want to moderate this.

It was a fun project that helped me learn more about Golang.

Take none of the security or anything else about it seriously, it was going to be a fun toy to have at KCDC2024, but ultimately I decided against making it publicly available.

I enjoyed it very much and have continued working on some other private repos that are completely unrelated to a message board.

I have chosen to leave it in the state it was when I stopped working on it. If you have any questions feel free to ask me about it! I always love chatting about tech/programming.

<br>
# Subtidal
A granular image board for serious, pseudoanonymous discussion of specific topics

Reason:
Other image boards are too plural in scope offering one board for the entirety of science.

## Gitpod development
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#github.com/sbrown17/subtidal)

## Stack
Go
Htmx
Postgresql
Digital ocean

## real checklist
- [x] board only shows threads
- [x] global navigation for boards
- [x] home page
- [x] threads can be posted in
- [x] threads and posts can be deleted by me
  - [x] simple auth with some cookie to know its my device?
  use this - https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
  also this - https://htmx.org/examples/async-auth
  - [x] just give every post an "x" button that makes a post request to remove the post or thread i want removed
  - [ ] make button behind a drop down where i have to click in 2 different spots (eg open the menu then click an x further down to confirm)
- [ ] bad word filter
- [ ] boards limited to certain number of threads
- [x] threads bump to top when talked in most recently
- [ ] threads saved when created and added to

## Checklist goals
- [x] boards based on route
  - [x] new page per board
    - [x] have div per thread with thumbnail/link/placeholder/text of initializing post
    - [x] have form for making new thread
- [ ] thread limit per board
  - [ ] message limit per thread
- [ ] hashed id based on poster/ip/device
- [ ] ability to delete posts (by poster)
- [ ] report system (make sure it cant be spammed, catpcha??)
- [ ] Code blocks
- [ ] references/cross links to other boards
- [ ] Pretty text? like italics etc.... markdown??
- [ ] Spoilers
- [ ] bad word filter (autospoiler them)
- [ ] settings saved to client (browser storage)
- [ ] log in system?
  - [ ] admins/mods?
- [ ] archive of old threads?

## Possible Boards
Philosophy
Mathematics
Science
  Physics
  Biology
  Chemistry
  Etc, eg. Geology, paleontology
Nonfiction lit
Fiction
Television/movies
Tech
  Software engineering
  Hardware
  Etc?

current:
https://coolors.co/0d1b2a-1b263b-415a77-778da9-e0e1dd

maybe:
https://coolors.co/0b132b-1c2541-3a506b-5bc0be-6fffe9
