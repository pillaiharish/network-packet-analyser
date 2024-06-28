function updateList(elementId, data) {
    const listElement = document.getElementById(elementId);
    listElement.innerHTML = ''; // Clear previous content
    data.forEach(item => {
        const div = document.createElement('div');
        div.textContent = `${item.Domain}: ${item.Count}`;
        listElement.appendChild(div);
    });
}

function fetchData() {
    fetch('/data')
        .then(response => response.json())
        .then(data => updateList('visitedList', data))
        .catch(error => console.error('Error fetching visited data:', error));

    fetch('/monitor')
        .then(response => response.json())
        .then(data => updateList('blockedList', data))
        .catch(error => console.error('Error fetching blocked data:', error));
}

setInterval(fetchData, 5000); // Update every 5 seconds
window.onload = fetchData; // Initial fetch on load
