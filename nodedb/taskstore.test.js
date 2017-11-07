"use strict";

const MongoStore = require("./taskstore");
const mongodb = require("mongodb");

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

describe("Mongo Task Store", () => {
    test("CRUD Cycle", () => {
        return mongodb.MongoClient.connect(mongoURL)
            .then(db => {
                let store = new MongoStore(db, "tasks");
                let task = {
                    title: "Learn Node.js to MongoDB",
                    tags: ["mongodb", "node.js", "info344"]
                };

                return store.insert(task)
                    .then(task => {
                        expect(task._id).toBeDefined();
                        return task._id;
                    })
                    .then(() => {
                        db.close();
                    })
                    .catch(err => {
                        db.close();
                        throw err;
                    });
            });
    });
});