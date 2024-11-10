

import React, { useState, useEffect } from 'react';

function SeasonSelector({ onSelectYear }) {
  const [years, setYears] = useState([]);

  useEffect(() => {
    const startYear = 1950;
    const currentYear = new Date().getFullYear();
    const allYears = [];
    for (let year = startYear; year <= currentYear; year++) {
      allYears.push(year);
    }
    setYears(allYears);
  }, []);

  return (
    <div>
      <label htmlFor="seasonSelect">Select F1 Season:</label>
      <select
        id="seasonSelect"
        onChange={(e) => onSelectYear(e.target.value)}
        defaultValue=""
      >
        <option value="" disabled>Select a year</option>
        {years.map((year) => (
          <option key={year} value={year}>{year}</option>
        ))}
      </select>
    </div>
  );
}

export default SeasonSelector;

