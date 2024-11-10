

import React, { useState, useEffect } from 'react';

function RaceSelector({ year, onSelectRace }) {
  const [races, setRaces] = useState([]);

  useEffect(() => {
    if (year) {
      fetch(`http://localhost:8080/races/${year}`)
        .then((response) => response.json())
        .then((data) => {
          const raceData = data.MRData?.RaceTable?.Races || [];
          setRaces(raceData);
        })
        .catch((error) => console.error('Error fetching races:', error));
    }
  }, [year]);

  const handleRaceSelect = (e) => {
    const selectedRace = races.find(race => race.raceName === e.target.value);
    onSelectRace(selectedRace);
  };

  return (
    <div>
      {year ? (
        <>
          <label htmlFor="raceSelect">Select Race:</label>
          <select id="raceSelect" onChange={handleRaceSelect} defaultValue="">
            <option value="" disabled>Select a race</option>
            {races.map((race, index) => (
              <option key={index} value={race.raceName}>{race.raceName}</option>
            ))}
          </select>
        </>
      ) : (
        <p>Please select a year to see the races.</p>
      )}
    </div>
  );
}

export default RaceSelector;

