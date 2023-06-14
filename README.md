## 📑 About the Project
<p align="justify">AirBnB Project<br>
  <br>
This RESTful API was developed by using Golang and written based on Clean Architecture principles. Built with Echo as web framework, GORM as ORM, MySQL as DBMS etc.
</p>

## 🛠 Tools
**Backend:** <br>
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/username/repository)
[![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)](https://code.visualstudio.com/)
[![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)](https://jwt.io/)
[![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)](https://swagger.io/)
[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/)
[![MidTrans](https://img.shields.io/badge/MidTrans-%234D4D4D.svg?style=for-the-badge)](https://www.midtrans.com/)


**Deployment:** <br>
![Google Cloud](https://img.shields.io/badge/googlecloud-%230db7ed.svg?style=for-the-badge&logo=googlecloud&logoColor=white)
![Google Cloud Storage](https://img.shields.io/badge/googlecloudstorage-%230db7ed.svg?style=for-the-badge&logo=googlecloudstorage&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

**Communication:**  
![GitHub](https://img.shields.io/badge/github%20Project-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-%237289DA.svg?style=for-the-badge&logo=discord&logoColor=white)

## 🔗 ERD
<img src="ERD-ImmersiveDashboard.png">

## 🔥 Open API

Simply [click here](https://app.swaggerhub.com/apis/dimasyudhana/AirBnB-RestfulAPI/1.0.0#) to see the details of endpoints we have agreed with our FE team.

<details>
  <summary>👶 Users </summary>
  
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /login              | -           |-                   | NO          | Login to the system                     |
| POST        | /register           | -           |-                   | YES         | Register a new user                     |
| PUT         | /users              | user_id     |-                   | YES         | Update user profile by admin            |
| DELETE      | /users              | user_id     |-                   | YES         | Deactive user profile by admin          |
| GET         | /users              | -           |-                   | YES         | List users                              |
| GET         | /users              | user_id     |-                   | YES         | Show user profile                       |
| PUT         | /users              | -           |-                   | YES         | Update user profile                     |
  
</details>

<details>
  <summary>📑 Classes</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /classes            | -           | YES         | Register new class                      |
| GET         | /classes            | -           | YES         | Get list class                          |
| GET         | /classes            | class_id    | YES         | Get class                               |
| PUT         | /classes            | class_id    | YES         | Edit class                              |
| DELETE      | /classes            | class_id    | YES         | Delete book                             |  

</details>

<details>
  <summary>📠 Mentees</summary>
  
| Method      | Endpoint            | Params                | JWT Token   | Function                                |
| ----------- | ------------------- | --------------------- | ----------- | --------------------------------------- |
| POST        | /mentees            | -                     | YES         | Register new mentee                     |
| GET         | /mentees            | -                     | YES         | List Mentees                            |
| GET         | /mentees            | mentee_id             | YES         | Mentee Profile include its Feedbacks    |
| PUT         | /mentees            | mentee_id             | YES         | Update Mentee Profile                   |
| DELETE      | /mentees            | mentee_id             | YES         | Deactive Mentee Profile                 |
| DELETE      | /mentees            | mentee_id/feedbacks   | YES         | Mentee Profile include its Feedbacks    |

  </details>

  <details>
   <summary>🔊 Feedbacks</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /feedbacks          | -           | YES         | Register feedback for mentees           |
| PUT         | /feedbacks          | feedback_id | YES         | Update selected feedback for mentees    |
| DELETE      | /feedbacks          | feedback_id | YES         | Unregister spesif feedback              |

  </details>

# 🛠️ How to Run Locally

- Clone it

```
$ git clone https://github.com/GroupProject3-Kelompok2/BE.git
```

- Go to directory

```
$ cd BE
```
- Run the project
```
$ go run main.go
```

- Voila! 🪄

### 🧰Backend

- [Github Repository for the Backend team](https://github.com/GroupProject3-Kelompok2/BE.git)
- [Swagger OpenAPI](https://app.swaggerhub.com/apis/dimasyudhana/AirBnB-RestfulAPI/1.0.0)


# 🤖 Author

-  Belhi Romdona  <br>  [![GitHub](https://img.shields.io/badge/BelhiRomdona-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/belhiibeng)
-  Dimas A Yudhana  <br>  [![GitHub](https://img.shields.io/badge/DimasYudhana-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/dimasyudhana)

<h5>
<p align="center">Created by Group 2 ©️ 2023</p>
</h5>
