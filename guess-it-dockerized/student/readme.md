Guess Numbers Project
📌 Description

This project is a number guessing program written in Go.
The program reads numbers from standard input (stdin) one by one and, for each number, prints a range in which the next number is expected to be.

The range is calculated using basic statistical methods learned in the math-skills project, such as:

Average (Mean)

Variance

Standard Deviation

The goal is to guess a correct range for the next number while keeping the range as small as possible to achieve a better score during testing.

📌 How It Works

Each input number represents a point on a graph:

X-axis: line index (0, 1, 2, 3, ...)

Y-axis: the number value

After receiving a number, the program prints:

<lower_bound> <upper_bound>


This range is the prediction for the next input number.

📌 Project Structure
.
├── go.mod
├── README.md
├── script.sh
└── student
    └── main.go

📌 Requirements

Go (golang)

Only standard Go packages are used

📌 How to Run Manually

From the project root:

go run ./student/main.go


Then provide numbers via standard input.

📌 script.sh

The tester will use this script to run the program.

Example script.sh:

#!/bin/sh
go run ./student/main.go


⚠️ Make sure the script is executable:

chmod +x script.sh

📌 go.mod

The project uses Go modules.

Example:

module guess-numbers

go 1.20

📌 Testing

The program is tested using multiple datasets (Data 1, Data 2, Data 3).

Scoring rules:

If the next number falls inside the predicted range, you score points.

Smaller ranges = higher score

Larger ranges = lower score

📌 Learning Objectives

This project helps you practice:

Statistical calculations

Probability-based guessing

Reading from standard input

File and script execution

Performance optimization