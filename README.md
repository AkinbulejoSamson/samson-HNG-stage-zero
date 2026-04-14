samson-HNG-stage-zero
A simple, production-ready REST API built with Go that provides a basic response when given a name. This project serves as the Stage 0 task for the HNG Internship.

🛠 Tech Stack
Language: Go (Golang)

Tooling: None (Zero external dependencies/frameworks for simplicity and speed).

Hosting: Railway

Version Control: Git & GitHub

📡 API Endpoint
Classify Name
Returns a classification response based on the provided name parameter.

URL: /classify

Method: GET

Query Parameter: name=[string]

Example Request:

Bash
GET /classify?name=Samson
