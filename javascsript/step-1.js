// Simple Calculator Step 1

// First Number Input
const firstNumber = parseFloat(prompt("Enter first Number:"));
// Operator Input
const operator = prompt("Enter an operator:");
// Second Number Input
const secondNumber = parseFloat(prompt("Enter second number:"));

// Operation Calculator
let result;

if (operator==="+") {
  result = firstNumber + secondNumber;
} else if (operator==="-") {
  result = firstNumber - secondNumber;
} else if (operator==="*") {
    result = firstNumber * secondNumber;
} else if (operator==="/") {
    result = firstNumber / secondNumber;
} else {
  result = "Invalid Operator";
}
alert(`The result is ${result}`);