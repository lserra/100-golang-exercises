# 📚 Exercises for Intermediate

**Author:** Laercio Serra <laercio.serra@gmail.com>  
**Version:** v0.0.1, March 2026

**Description:** Rewritten from the 100+ Python challenging programming exercises  
**Keywords:** Golang, programming, exercises, challenges  

---

This document contains all the intermediate exercises. If you find an error, bug, or something that could be written better, please get in contact with me. Thanks! 🙏

## 📐 Exercise 006: Calculate Values Using a Formula

Write a program that calculates and prints values according to the given formula:

**Q = √[(2 × C × D) / H]**

Following are the fixed values of C and H:

- C is 50
- H is 30
- D is the variable whose values should be input to your program in a comma-separated sequence

**Example:**

- **Input:** `100,150,180`
- **Expected Output:** `18,22,24`

### 💡 Hints

- Use `math.Sqrt` for the square root
- Use `math.Round` for rounding float operations

### ✅ Solution

[View Solution](006/exercise006.go)

---

## 🔢 Exercise 007: Generate a 2D Array

Write a program that takes 2 digits, X and Y, as input and generates a 2-dimensional array. The element value in the i-th row and j-th column of the array should be i × j.

**Note:** i = 0, 1, ..., X-1; j = 0, 1, ..., Y-1

**Example:**

- **Input:** `3,5`
- **Expected Output:** `[[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]`

### 💡 Hints

Input data should be assumed to be console input in a comma-separated format.

### ✅ Solution

[View Solution](007/exercise007.go)

---

## 🔤 Exercise 008: Sort Words Alphabetically

Write a program that accepts a comma-separated sequence of words as input and prints the words in a comma-separated sequence after sorting them alphabetically.

**Example:**

- **Input:** `without,hello,bag,world`
- **Expected Output:** `bag,hello,without,world`

### 💡 Hints

Input data should be assumed to be console input.

### ✅ Solution

[View Solution](008/exercise008.go)

---

## 🔠 Exercise 009: Capitalize All Characters

Write a program that accepts a sequence of lines as input and prints the lines after capitalizing all characters in each sentence.

**Example:**

**Input:**
```
Hello world
Practice makes perfect
```

**Expected Output:**
```
HELLO WORLD
PRACTICE MAKES PERFECT
```

### 💡 Hints

Input data should be assumed to be console input.

### ✅ Solution

[View Solution](009/exercise009.go)

---

## 🔣 Exercise 010: Remove Duplicates and Sort Words

Write a program that accepts a sequence of whitespace-separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically.

**Example:**

- **Input:** `hello world and practice makes perfect and hello world again`
- **Expected Output:** `again and hello makes perfect practice world`

### 💡 Hints

- Input data should be assumed to be console input
- Use a set-like container to remove duplicated data automatically, then sort the data

### ✅ Solution

[View Solution](010/exercise010.go)

---

## 💾 Exercise 011: Filter Binary Numbers Divisible by 5

Write a program that accepts a sequence of comma-separated 4-digit binary numbers as input and checks whether they are divisible by 5 or not. The numbers that are divisible by 5 should be printed in a comma-separated sequence.

**Example:**

- **Input:** `0100,0011,1010,1001`
- **Expected Output:** `1010`

**Note:** Assume the data is input by console.

### 💡 Hints

Input data should be assumed to be console input.

### ✅ Solution

[View Solution](011/exercise011.go)

---

## 🧮 Exercise 012: Find Numbers with All Even Digits

Write a program that finds all numbers between 100 and 300 (both included) such that each digit of the number is an even number. The numbers obtained should be printed in a comma-separated sequence on a single line.

### 💡 Hints

Input data should be assumed to be console input.

### ✅ Solution

[View Solution](012/exercise012.go)

---

## 📊 Exercise 013: Count Letters and Digits

Write a program that accepts a sentence and calculates the number of letters and digits.

**Example:**

- **Input:** `hello world! 123`

**Expected Output:**
```
LETTERS 10
DIGITS 3
```

### 💡 Hints

Input data should be assumed to be console input.

### ✅ Solution

[View Solution](013/exercise013.go)
