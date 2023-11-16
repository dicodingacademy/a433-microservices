require('dotenv').config()

const express = require("express");
const app = express();

const bp = require("body-parser");
app.use(bp.json());

const amqp = require("amqplib");
const amqpServer = process.env.AMQP_URL;
var channel, connection;

connectToQueue();

async function connectToQueue() {
    connection = await amqp.connect(amqpServer);
    channel = await connection.createChannel();
    try {
        const queue = "order";
        await channel.assertQueue(queue);
        console.log("Connected to the queue!")
    } catch (ex) {
        console.error(ex);
    }
}

app.post("/order", (req, res) => {
    const { order } = req.body;
    createOrder(order);
    res.send(order);
});

const createOrder = async order => {
    const queue = "order";
    await channel.sendToQueue(queue, Buffer.from(JSON.stringify(order)));
    console.log("Order succesfully created!")
    process.once('SIGINT', async () => { 
        console.log('got sigint, closing connection');
        await channel.close();
        await connection.close(); 
        process.exit(0);
    });
};

app.listen(process.env.PORT, () => {
    console.log(`Server running at ${process.env.PORT}`);
});

