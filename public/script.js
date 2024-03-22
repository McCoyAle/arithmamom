document.addEventListener("DOMContentLoaded", function() {
    // DOM elements
    const nameInput = document.getElementById("name");
    const nameBtn = document.getElementById("nameBtn");
    const operationSelection = document.getElementById("operationSelection");
    const startBtn = document.getElementById("startBtn");
    const questionsContainer = document.getElementById("questions");
    const resultContainer = document.getElementById("result");

    // Event listeners
    nameBtn.addEventListener("click", function() {
        const name = nameInput.value.trim();
        if (name !== "") {
            // Hide name input and show operation selection
            nameInput.disabled = true;
            nameBtn.disabled = true;
            operationSelection.style.display = "block";
        } else {
            alert("Please enter your name.");
        }
    });

    startBtn.addEventListener("click", function() {
        const operation = operationSelection.value;
        if (operation !== "") {
            // Hide operation selection
            operationSelection.disabled = true;
            startBtn.disabled = true;
            questionsContainer.style.display = "block";

            // Generate questions based on selected operation
            fetchQuestions(operation);
        } else {
            alert("Please select a mathematical operation.");
        }
    });

// Function to fetch and display questions
function fetchQuestions(operation) {
    // Send a GET request to the backend endpoint to fetch questions based on the selected operation
    fetch("/mathques?operation=" + operation)
        .then(response => {
            // Check if the response is successful (status code 200)
            if (response.ok) {
                // Parse the JSON response
                return response.json();
            } else {
                // If the response is not successful, throw an error
                throw new Error("Failed to fetch questions: " + response.status);
            }
        })
        .then(data => {
            // Once the JSON data is received, populate questionsContainer with the questions
            // Clear previous content in questionsContainer
            questionsContainer.innerHTML = "";
            // Iterate over each question in the data and create HTML elements to display them
            data.questions.forEach((question, index) => {
                // Create a new paragraph element for each question
                const questionElement = document.createElement("p");
                // Set the text content of the paragraph to the question text
                questionElement.textContent = "Question " + (index + 1) + ": " + question;
                // Append the paragraph element to the questionsContainer
                questionsContainer.appendChild(questionElement);
            });
        })
        .catch(error => {
            // If an error occurs during the fetch request, display an error message
            console.error("Error fetching questions:", error.message);
            questionsContainer.innerHTML = "<p>Error fetching questions. Please try again later.</p>";
        });
}

    // Function to submit answers and calculate score (you'll need to implement this)
    function submitAnswers() {
        // Fetch user answers from the page and send them to the server for validation
        // Example: fetch("/submit", { method: "POST", body: JSON.stringify(answers) })
        // Then parse the response and display the score
        // Example: resultContainer.innerHTML = "<p>Your score is: " + score + "</p>";
    }
});
