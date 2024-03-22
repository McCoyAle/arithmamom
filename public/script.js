// Define a function to fetch a random math question from the backend
async function fetchQuestion(name) {
    const response = await fetch('/mathques?name=' + encodeURIComponent(name), {
        method: 'GET'
    });
    const data = await response.text();
    return data;
}

// Define a function to submit the user's answer to the backend
async function submitAnswer(name, answer) {
    // Check if the answer is empty
    if (!answer.trim()) {
        alert("Please provide an answer.");
        return;
    }
    
    const response = await fetch('/mathques?name=' + encodeURIComponent(name) + '&answer=' + encodeURIComponent(answer), {
        method: 'POST'
    });
    const data = await response.text();
    return data;
}

// Main function to fetch question when the page loads
window.onload = async function() {
    const name = prompt("Please enter your name:");
    if (name !== null && name.trim() !== "") {
        const question = await fetchQuestion(name);
        document.getElementById('question').innerText = question;
        document.getElementById('submit').addEventListener('click', async function() {
            const userAnswer = document.getElementById('answer').value;
            const result = await submitAnswer(name, userAnswer);
            document.getElementById('result').innerText = result;
        });
    } else {
        alert("Name cannot be empty. Please refresh the page and try again.");
    }
}
