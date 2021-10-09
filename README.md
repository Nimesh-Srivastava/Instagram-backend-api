# Instagram Backend API using GoLang and MongoDB
Appointy tech internship task

The complete api is covered in `main.go` file. The api works for all endpoints that were mentioned int the task.

Passwords have been securely stored on the server using `bcrypt` golang default package. This ensures that passwords cannot be reverse engineered by anyone.

There is never more than one simultaneous call to function throughout the api, hence, thread safety is ensured.

The entire 