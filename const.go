package main

const HOST string = "localhost"
const PORT int = 5432
const DBNAME string = "postgres"

const HOMEPAGE = `
<h1> User Homepage </h1>
<a href="http://localhost:8000/userview">View user profile</a>
<br>
<a href="http://localhost:8000/useredit">Update my profile</a>
<br>
<a href="http://localhost:8000/userdelete">Delete my profile</a>
<br>
<a href="http://localhost:8000/post">Create a new post</a>
<br>
<a href="http://localhost:8000/postview">View posts</a>
<br>`


const INDEXPAGE = `
<h1> Welcome! </h1>
<a href="http://localhost:8000/signup">Sign Up</a>
<br>
<a href="http://localhost:8000/login">Log In</a>`