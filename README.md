# Basic-project-sejutacita
Project Test 
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

<br>

# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Start Docker](#Getting-Start-With-Docker)
- [Our Feature](#Our-Feature)
- [Endpoints](#endpoints)

<br>


# Description

Project Erajaya Basic Rest-API
<br>
<p>Project API sederhana dengan menggunakan 2 enpoint untuk menambah dan mengurutkan list product, project ini dapat dijalankan di localhost
dan juga dapat dibungkus ke dalam container dengan mengguankn docker</p>

<br>


# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench
* Docker


<br>

# High Level Architecture
<img src="https://github.com/alfiancikoa/project-rest-api/blob/main/img/HLA.jpg">

# How to use
<h5>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instruction</a></h5>

- Clone this repository in your $PATH:
```
$ git clone https://github.com/alfiancikoa/project-rest-api.git
```
<h3>If you are using port from your localhost</h3>

- Create Environment file (.env) and fill with
  
```
MYSQL_USER=root           // your user MySQL 
MYSQL_PASSWORD=password   // your password MySQL
MYSQL_HOST=localhost      // your IP local default localhost
MYSQL_PORT=3306           // default port from MySQL
MYSQL_DBNAME=dbproject    // database name
```

- Create Database If Not Exist On the Terminal


```
# login mysql
$ sudo mysql -u root -p

# on menu mysql
mysql> create database <db-name>;

# to show all databases list
mysql> show databases;

# to exit
mysql> exit;
```
- Running the Program
```
$ go run main.go
```
# Getting Start With Docker

Common commands
```
# pull image from container registry
docker pull <image-name>

# login to container registry
docker login -u <user-name>

# build image
docker build -t <image-name>[:tag] .

# push image to container registry
docker push <image-name>

# create and run container
docker run -p <host-port>:<container-port> -e <env-name>=<env-value> -v <host-volume>:<container-volume> --name <container-name> <image-name>

# run existing container
docker start <container-name>
docker container start <container-name>

# stop running container
docker stop <container-name>
docker container stop <container-name>

# remove container
docker rm <container-name>
docker container rm <container-name>

# show images
docker image list

# show container
docker container list
docker ps
docker ps -a

# access/run command in a container
docker exec -it <container-name> <command>
docker exec -it mysql bash
```
<br>

# Using Docker
If you are using docker-desktop, the containers can access host os by using host.docker.internal name.
Otherwise, you can use default host IP address: 172.17.0.1
<br>
```
# create and run dockerMySQL container
docker run -p 3307:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=dbproject --name dockerContainer -d mysql:latest
```
<br>
- Make sure your IP Address of Docker Container from dockerMySQL

```
# make sure the dockerMySQL container is running
docker container ls

# check ip Address of each container when it is start/run
docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)

# in this command you will get the ip address from container running
# example 172.17.0.1 or 172.17.0.2
# please make sure the ip address from MySQL container
```
<br>
- Set Environtmen on file .env with your configuration

```
MYSQL_USER=root              // your user MySQL 
MYSQL_PASSWORD=password      // make sure this is same with previous command (MYSQL_ROOT_PASSWORD=password)
MYSQL_HOST=172.17.0.1        // default IP Docker || or ip where you have checked in previous command 
MYSQL_PORT=3306              // default port MySQL
MYSQL_DBNAME=dbproject       // make sure this is same with previous command (MYSQL_DATABASE=dbproject)
```

<br>
- Build The Image and then Run the Container

```
# Build the image
docker build -t nama-image:latest .

# create and run appContainer
docker run --name nameContainerAPI -e -p 8080:8080 nama-image:latest

# Stop using ctrl+C
# if you want to run app again
docker container start nameContainerAPI

# if you want to stop app
docker container stop nameContainerAPI
```

<br>

* Open Postman run with your localhost, follow the routes in the Visual Studio Code folder.
* <h4>more instruction details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instruction</a></h4>

<br>

# Our Feature
* CREATE: Products
* READ: Products

<br>

# Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST  | /products | Insert New Products | No | No
| GET | /products | Get All Products List | No | No
| GET    | /products/filter/:request | Get All Products Filter by | No | No

<br>

# Filter Use
Field the filter to use on :require
* ascending ( To show product starting from A )
* descending ( To show product starting from Z )
* lower ( To show product from lower price )
* upper ( To show product from higher price )
* new ( To show product starting from latest Product )


<br>
Example Test API in localhost using Postman
<br>

```
<method> <endpoint>
- POST: localhost:8080/products
- GET: localhost:8080/products
- GET: localhost:8080/products/filter/new
```

<h3>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instructions</a></h3>
<h4>For Unit Testing Coverage Report <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/img/Coverage%20Report%20-%20.png">coverage</a></h4>
<br>
