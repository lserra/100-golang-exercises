# 📚 Exercises for Beginners

**Author:** Laercio Serra <laercio.serra@gmail.com>  
**Version:** v0.0.1, March 2026

**Description:** Rewritten from the 100+ Python challenging programming exercises  
**Keywords:** Golang, programming, exercises, challenges  

---

This document contains all the beginner exercises. If you find an error, bug, or something that could be written better, please get in contact with me. Thanks! 🙏

## 🔢 Exercise 001: Find Numbers Divisible by 7 but Not by 5

Write a program that finds all numbers divisible by 7 but are not multiples of 5, between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

### 💡 Hint

Consider using `strconv` and `strings.Join`.

### ✅ Solution

[View Solution](001/exercise001.go)

---

## 🧮 Exercise 002: Compute Factorial

Write a program that computes the factorial of a given number. The results should be printed in a comma-separated sequence on a single line.

**Example:**

- **Input:** `8`
- **Output:** `40320`

### 💡 Hints

In case of input data being supplied to the question, it should be assumed to be console input.

### ✅ Solution

[View Solution](002/exercise002.go)

---

## 📊 Exercise 003: Create a Map with Numbers Squared

Given an integral number `n`, write a program to generate a map that contains `(i, i*i)` where `i` is an integral number between 1 and `n` (both included). The program should then print the map representation.

**Example:**

- **Input:** `8`
- **Expected Output:** `map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]`

### 💡 Hints

- Use `make` to create the map
- Use `%v` verb for output formatting

### ✅ Solution

[View Solution](003/exercise003.go)

---

## 📝 Exercise 004: Create a Slice from Comma-Separated Input String

Write a program that accepts a sequence of comma-separated numbers from the console and generates a slice from them. Return the slice.

**Example:**

- **Input:** `34, 67, 55, 33, 12, 98`
- **Expected Output:** `[34 67 55 33 12 98]`

### 💡 Hints

- Input data should be assumed to be console input
- The `strings` package has a `Split` method

### ✅ Solution

[View Solution](004/exercise004.go)

---

## 🏗️ Exercise 005: Define a "Class" (Interface) with at Least Two Methods

Create a separate file (module) that has at least two methods:

- **ReadString:** Read a string from console input
- **PrintString:** Return the string in upper case

Also create a `main.go` file that acts as the calling class.

### 💡 Hints

- Use `bufio.NewScanner(os.Stdin)` to read a full line of text
- Use `go run .` to execute all files in the directory

### ✅ Solution

[View Solution](005/exercise005.go)
