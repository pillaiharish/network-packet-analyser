function fetchData() {
    fetch('/data')
        .then(response => response.json())
        .then(data => {
            const domainList = document.getElementById('domainList');
            domainList.innerHTML = ''; // Clear previous data
            data.forEach(item => {
                const domainItem = document.createElement('div');
                domainItem.className = 'domain-item';
                domainItem.innerHTML = `<strong>${item.Domain}</strong> <span>${item.Count}</span>`;
                domainList.appendChild(domainItem);
            });
        })
        .catch(error => console.error('Error fetching data:', error));
}

// Poll the server every 5 seconds to refresh the data
setInterval(fetchData, 5000);

// Also fetch data on initial load
window.onload = fetchData;
