# Instagram Backend API using GoLang and MongoDB

![Language](https://img.shields.io/badge/Language-GO-blue.svg?style=for-the-badge&logo=appveyor)&nbsp;&nbsp;
![Testing](https://img.shields.io/badge/Testing-Postman-orange.svg?style=for-the-badge&logo=appveyor)&nbsp;&nbsp;
![DB](https://img.shields.io/badge/DB-MongoDB-green.svg?style=for-the-badge&logo=appveyor)&nbsp;&nbsp;
![Type](https://img.shields.io/badge/Type-RestfulAPI-brown.svg?style=for-the-badge&logo=appveyor)&nbsp;&nbsp;

<br>

## 💢 Appointy tech internship task

The complete api is covered in `main.go` file.

<br>

## 💢 Pre-requisites
⚠️ Before running the application, please ensure you have the following packages installed :
* Mongodb for system [Refer this link](https://docs.mongodb.com/manual/installation/)
* GoLang for system [Refer this link](https://golang.org/doc/install)
* Gorilla MUX [Refer this link](https://github.com/gorilla/mux)

<br>

## 💢 Password
Passwords have been securely stored on the server using `bcrypt` golang default package. This ensures that passwords cannot be reverse engineered by anyone.

<br>

## 💢 Thread Safety
There is never more than one simultaneous call to function throughout the api, hence, thread safety is ensured.

<br>

## 💢 API testing
The entire api has been tested using *Postman* API platform. Some screenshots are attached below for reference :<br>
<br>

:atom: **Add User :**
<br>

![Image 09-10-21 at 9 15 PM](https://user-images.githubusercontent.com/30381993/136665838-b9ff3388-bf6d-4c1b-a4db-a7ef468fd2ec.jpg)

<br>
<br>

:atom: **Add post :**
<br>

![Image 09-10-21 at 9 16 PM](https://user-images.githubusercontent.com/30381993/136665853-434bd2ba-8fd6-4c71-afc0-b31b429cbc31.jpg)

<br>
<br>

:atom: **List posts by one user :**
<br>

![Image 09-10-21 at 9 16 PM](https://user-images.githubusercontent.com/30381993/136665899-52f62963-086a-40ab-a1d3-a74ce444dc72.jpg)


<br>
<br>

## Conclusion
The api is fully functioning and works for all endpoints mentioned in the task.
