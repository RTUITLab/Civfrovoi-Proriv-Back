<p align="center">
  <a href="" rel="noopener">
 <img src="https://i.imgur.com/AZ2iWek.png" alt="Project logo"></a>
</p>
<h3 align="center"></h3>

<div align="center">

[![Hackathon](https://img.shields.io/badge/hackathon-Цифровой_Прорыв-orange.svg)](http://hackathon.url.com)
[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)

</div>

---

<p align="center">
    <br>Easy interaction of the dispatcher with the workers.
</p>

## 📝 Table of Contents

- [Links on other parts of the project](#projects)
- [Problem Statement](#problem_statement)
- [Idea / Solution](#idea)
- [Dependencies / Limitations](#limitations)
- [Future Scope](#future_scope)
- [Setting up a local environment](#getting_started)
- [Technology Stack](#tech_stack)
- [Authors](#authors)
- [Acknowledgments](#acknowledgments)
## 🔗 Links on other parts of the project <a name = "projects"></a>
- [Hak21-Dispatcher](https://github.com/RTUITLab/Hak21-Dispatcher) - Dispatcher frontend.
- [sheremetyevo_tasks](https://github.com/mentoster/sheremetyevo_tasks) - Workers frontend.
- [Civfrovoi-Proriv-Back](https://github.com/RTUITLab/Civfrovoi-Proriv-Back) - Backend.
## 🧐 Problem Statement <a name = "problem_statement"></a>

The main problem is that if there is snow on the ground, the airport can not send
and receive flights.
#### **Minor issues:**

is no tool to ensure fast and efficient snow removal at the
airport.

- There is no visualization of the snow removal process and the current status.

- There is a need for a convenient tool for assigning tasks and monitoring their
execution.

- There is a need to calculate and assign the necessary resources for
snow removal, depending on the weather situation

## 💡 Idea / Solution <a name = "idea"></a>

Our solution is based on simplifying job creation and worker management.


## 🚀 Future Scope <a name = "future_scope"></a>

We want to improve the UI, refractor code.
Make the dispatcher work more autonomously based on the flights of planes and the paths of people.
Ability to track the state of snow by cameras

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development
and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

[Go]("https://golang.org")
[Docker]("https://www.docker.com")

Create File ```src/.env``` for launch from docker:
```.env
CP_DATABASE_URL=root:examplepassword@tcp(db_container:3306)/coords?parseTime=True
CP_APP_PORT=8080
CP_AMQP_SERVER_URL=amqp://docker:docker@rabbit_container:5672/
RABBITMQ_DEFAULT_USER=docker
RABBITMQ_DEFAULT_PASS=docker
MYSQL_ROOT_PASSWORD=examplepassword
MYSQL_USER=docker
MYSQL_PASSWORD=docker
```

### Installing


```
docker-compose up --build
```

## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## ⛏️ Built With <a name = "tech_stack"></a>

- [gRPC](https://grpc.io/) -A high performance, open source universal RPC framework
- [Unity](https://unity.com/) - The leading platform for creating interactive, real-time content
- [Flutter](https://flutter.dev/) -Flutter is Google's UI toolkit for building beautiful, natively compiled applications for mobile, web, desktop, and embedded devices from a single codebase.

## ✍️ Authors <a name = "authors"></a>

- [@Kubirill](https://github.com/Kubirill) - Idea & Initial work
- [@0B1t322](https://github.com/0B1t322) -  work
- [@mentoster](https://github.com/mentoster) -  work



## 🎉 Acknowledgments <a name = "acknowledgments"></a>

- Thx to RtuItLab
