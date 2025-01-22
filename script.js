document.addEventListener("DOMContentLoaded", () => {
  const dataOutput = document.getElementById("data-output");

  // Fetching data from the backend API
  async function fetchData() {
    try {
      const response = await fetch("http://localhost:8080/api/data");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      dataOutput.textContent = JSON.stringify(data, null, 2);
    } catch (error) {
      dataOutput.textContent = `Error: ${error.message}`;
    }
  }

  // Fetch data every 5 seconds
  fetchData();
  setInterval(fetchData, 5000);
});
