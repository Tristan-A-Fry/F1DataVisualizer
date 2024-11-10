
import React, { useState } from 'react';
import SeasonSelector from './components/SeasonSelector';
import RaceSelector from './components/RaceSelector';
import Results from './components/Results';

function App() {
  const [selectedYear, setSelectedYear] = useState(null);
  const [selectedRace, setSelectedRace] = useState(null);

  const handleYearSelect = (year) => {
    setSelectedYear(year);
    setSelectedRace(null); // Reset the race when the year changes
  };

  const handleRaceSelect = (race) => {
    setSelectedRace(race);
  };

  return (
    <div>
      <h1>Formula One Data Visualizer</h1>
      <SeasonSelector onSelectYear={handleYearSelect} />
      {selectedYear && <RaceSelector year={selectedYear} onSelectRace={handleRaceSelect} />}
      {selectedYear && selectedRace && <Results year={selectedYear} race={selectedRace} />}
    </div>
  );
}

export default App;



