// Simple Calculator Step 2

// First Number Input
const firstNumber = parseFloat(prompt("Enter first Number:"));
// Operator Input
const operator = prompt("Enter an operator:");
// Second Number Input
const secondNumber = parseFloat(prompt("Enter second number:"));

// Operation Calculator
let result;

switch (operator) {
  case "+":
    result = firstNumber + secondNumber;
    break;
  case "-":
    result = firstNumber - secondNumber;
    break;
  case "*":
    result = firstNumber * secondNumber;
    break;
  case "/":
    result = firstNumber / secondNumber;
    break;
  default:
    result = "Invalid Operator";
    break;
}
alert(`The result is ${result}`);