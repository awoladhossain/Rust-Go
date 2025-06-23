import cors from "cors";
import express from "express";
import mongoose from "mongoose";

// No dotenv import or config since .env file is deleted

const app = express();
app.use(cors());
app.use(express.json());
const PORT = 5000;

// Use hardcoded defaults for MongoDB connection (since no env file)
const MONGO_HOST = "localhost";
const MONGO_PORT = "27017";
const MONGO_DB = "mydatabase";
const MONGO_URI = `mongodb://admin:qwerty@${MONGO_HOST}:${MONGO_PORT}/${MONGO_DB}?authSource=admin`;

const userSchema = new mongoose.Schema(
  {
    name: String,
    email: String,
    // add other fields as needed
  },
  { collection: "User" }
);

mongoose.model("User", userSchema);

app.get("/", (req, res) => {
  res.send("hello world");
});

app.get("/getUsers", async (req, res) => {
  try {
    // Assuming you have a User model defined with mongoose
    const users = await mongoose.model("User").find();
    res.json(users);
  } catch (error) {
    res.status(500).json({ error: "Failed to fetch users" });
  }
});

app.post("/addUser", async (req, res) => {
  try {
    const User = mongoose.model("User");
    const newUser = new User(req.body);
    const savedUser = await newUser.save();
    res.status(201).json(savedUser);
  } catch (error) {
    res.status(400).json({ error: "Failed to add user" });
  }
});

mongoose
  .connect(MONGO_URI)
  .then(() => {
    app.listen(PORT, () => {
      console.log(`Server running on port ${PORT}`);
      console.log(`Connected to MongoDB at ${MONGO_URI}`);
    });
  })
  .catch((err) => console.error("MongoDB connection error:", err));
