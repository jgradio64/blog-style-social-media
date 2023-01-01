# blog-style-social-media


## Contributors
[Shilei Wu](https://github.com/wushileilei)

[Benjamin Eich](https://github.com/jgradio64)

---

## About this project 

### Summary
A simple full stack web application that allows users to create and comment on text blog posts.

### Purpose
This was our idea for a class project that needed to be written in the language Go.

### Tools used

- [Bootstrap 4.3](https://getbootstrap.com/)
    - For Styling of the webpages.
- [Go](https://go.dev/)
    - Serves as the driver for all the back end logic.
    - Communicates with the server.
    - Handles front end conditional logic when rendering html templates.
- [PostgreSQL](https://www.postgresql.org/)
- [PG Admin](https://www.pgadmin.org/)
    - Used to creating and managing the dataabse.
- [Digital Ocean](https://www.pgadmin.org/)
    - Cloud provider used to host the PostgreSQL database.
- [Google uuid](https://github.com/google/uuid)
    - To create and manage user session tokens.
    - Implemented in application memory.

---

## Functionality of "blog-style-social-media"
1. User can create, manage and delete an account.
2. User can create, manage and delete their own posts.
3. Users can view the profiles of other users.
4. Users can follow other users, and view the posts of all users on their homepage.
5. Users can view the posts of all other users from the appropriately named "all" page.
6. Users can create, edit and delete their own comments on any post.