# Quacker - Single Page Web App

### Webpage (quacker.com)  --  Route  |  View 
1. Logged in
    * "/home"                         | Timeline
    * "/[user_tag]"                   | User's Quacks
    * "/[user_tag]/with_replies"      | User's Quacks&Replies
    * "/[user_tag]/likes"             | User's Likes
    * "/settings/profile"             | Profile Settings
    * "/settings/account"             | Account Settings
    * "/[user_tag]/quack/[quack_id]"  | Quack

2. Guest
    * "/"                             | Signup&Login
    * "/[user_tag]"                   | User's Qaucks           
    * "/[user_tag]/with_replies"      | User's Quacks&Replies   (Can access?)
    * "/[user_tag]/likes"             | User's Likes            (Can access?)   
    * "/[user_tag]/quack/[quack_id]"  | Quack


### Web Server  --  Method  |  Endpoint
* GET   |   "/"
    * returns html template
* POST, PUT, DELETE  |  "/api/user"
* POST, PUT, DELETE  |  "/api/quack"
* POST, DELETE       |  "/api/like"
* POST, DELETE       |  "/api/connection"
* POST               |  "/api/login"
* POST               |  "/api/signup"

* GET                |  "/api/user/tag=?"
    * returns user info with quacks and likes
* GET                |  "/api/user/timeline"
    * returns quacks from followings
* GET                |  "/api/quack/id=?"
    * returns quack info with replies
* GET                |  "/api/quack/id=?/requacks"
    * returns users' info who requacked this quack
* GET                |  "/api/quack/id=?/likes"
    * returns users' info who liked this quack
* GET                |  "/api/quack/id=?/quoted"
    * returns quacks that quoted this quack

### Relational Database  --  Er Diagram
![](https://imgur.com/Hr3BxQT.jpg)
