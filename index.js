require('dotenv').config()

const express = require("express");
const app = express();

const bp = require("body-parser");

const amqp = require("amqplib");
const amqpServer = process.env.AMQP_URL;
var channel, connection;

connectToQueue();

async function connectToQueue() {
    try {
        connection = await amqp.connect(amqpServer);
        channel = await connection.createChannel();
        await channel.assertQueue("order");
        channel.consume("order", data => {
            console.log(`Order received: ${Buffer.from(data.content)}`);
            console.log("** Will be shipped soon! **\n")
            channel.ack(data);
        });
    } catch (ex) {
        console.error(ex);
    }
}

app.listen(process.env.PORT, () => {
    console.log(`Server running at ${process.env.PORT}`);
});
