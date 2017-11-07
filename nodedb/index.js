"use strict";

const mongodb = require("mongodb");
const MongoStore = require("./taskstore");
const express = require("express");
const app = express();

const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

//connect to MongoDB
mongodb.MongoClient.connect(mongoURL)
    .then(db => {
        //initialize a new task store
        let taskStore = new MongoStore(db);

        //parses posted JSON and makes
        //it available from req.body
        app.use(express.json());

        app.post("/v1/tasks", (req, res) => {
            //insert a new task
            //use taskStore.insert()
        });

        app.get("/v1/tasks", (req, res) => {
            //return all not-completed tasks in the database
        });

        app.patch("/v1/tasks/:taskID", (req, res) => {
            let taskIDToFetch = req.params.taskID;
            //update single task by id
        });

        app.listen(port, host, () => {
            console.log(`server is listening at http://${addr}....`);
        });
    });

