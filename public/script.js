// Define a function to fetch a random math question from the backend
async function fetchQuestion() {
    const response = await fetch('/question', {
        method: 'GET'
    });
    const data = await response.json();
    return data.question;
}

// Define a function to submit the user's answer to the backend
async function submitAnswer() {
    const userAnswer = document.getElementById('userAnswer').value;
    const response = await fetch('/answer', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ answer: userAnswer })
    });
    const data = await response.json();
    document.getElementById('result').innerText = data.result;
}

// Main function to fetch question when the page loads
window.onload = async function() {
    const question = await fetchQuestion();
    document.getElementById('question').innerText = question;
}
