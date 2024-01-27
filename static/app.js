const Controller = {
  currentPage: 1,

  search: async (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    try {
      const response = await fetch(`/search?q=${data.query}`);
      const results = await response.json();
      Controller.updateTable(results);
    } catch (error) {
      console.error('Error fetching search results:', error);
    }  
  },

  loadMore: async (ev) => {
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    try {
      const response = await fetch(`/search?q=${data.query}`);
      const results = await response.json();
      Controller.getMoreTables(results);
    } catch (error) {
      console.error('Error fetching more pages:', error);
    }  
  },

  updateTable: (results) => {
    const table = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      rows.push(`<tr><td>${result}</td></tr>`);
    }
    table.innerHTML = rows;
  },

  getMoreTables: (results) => {
    const table = document.getElementById("table-body");
    const itemsPerPage = 10;

    // Calculate the start and end index for the current page
    const startIndex = (Controller.currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;

    // Extract the results for the current page
    const currentPageResults = results.slice(startIndex, endIndex);

    
    const rows = currentPageResults.map(result => `<tr><td>${result}</td></tr>`);

    // Append the rows to the table
    table.innerHTML += rows.join("");

    Controller.currentPage++;

    // If there are no more results, hide the "Load More" button
    if (results.length <= endIndex) {
      const loadMoreButton = document.getElementById("load-more");
      loadMoreButton.style.display = "none";
    }
  },
}; 

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);

const loadMoreButton = document.getElementById("load-more");
loadMoreButton.addEventListener("click", Controller.loadMore);
