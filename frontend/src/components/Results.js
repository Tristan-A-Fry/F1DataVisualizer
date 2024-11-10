
import React, { useState, useEffect } from 'react';
import '../styles/Results.css'; // Import the CSS for custom styling

function Results({ year, race }) {
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (year && race) {
      console.log(`Fetching results for year: ${year}, race: ${race.raceName}, round: ${race.round}`);
      setLoading(true);
      setError(null);
      setResults([]); // Clear existing results to avoid showing stale data

      fetch(`http://localhost:8080/results/${year}?round=${race.round}&limit=100`)
        .then((response) => {
          if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
          }
          return response.json();
        })
        .then((data) => {
          console.log('Fetched data:', data);
          const allRaces = data.MRData?.RaceTable?.Races || [];
          console.log('All races:', allRaces);

          const filteredRace = allRaces.find((r) => r.round === race.round.toString());
          console.log('Filtered race:', filteredRace);

          if (filteredRace && filteredRace.Results) {
            setResults(filteredRace.Results);
          } else {
            setResults([]);
          }

          setLoading(false);
        })
        .catch((error) => {
          console.error('Error fetching results:', error);
          setError('Failed to fetch results. Please try again.');
          setLoading(false);
        });
    }
  }, [year, race]);

  return (
    <div className="results-container">
      {loading ? (
        <p>Loading results...</p>
      ) : error ? (
        <p>{error}</p>
      ) : year && race ? (
        <>
          <h2>Results for {race.raceName} ({year})</h2>
          {results.length > 0 ? (
            <div className="results-graphic">
              {results.map((result, index) => (
                <div className="driver-result" key={index}>
                  <div className="position-number">#{result.position}</div>
                  <div
                    className={`driver-circle ${result.Constructor.constructorId}`}
                  >
                    <span className="driver-code">{result.Driver.code}</span>
                  </div>
                  <div className="time-segment">
                    {result.Time?.time || result.status || 'N/A'}
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <p>No results found for this race.</p>
          )}
        </>
      ) : (
        <p>Select a race to see the results.</p>
      )}
    </div>
  );
}

export default Results;



// import React, { useState, useEffect } from 'react';
//
// function Results({ year, race }) {
//   const [results, setResults] = useState([]);
//   const [loading, setLoading] = useState(false);
//   const [error, setError] = useState(null);
//
//   useEffect(() => {
//     if (year && race) {
//       console.log(`Fetching results for year: ${year}, race: ${race.raceName}, round: ${race.round}`);
//       setLoading(true);
//       setError(null);
//       setResults([]); // Clear existing results to avoid showing stale data
//
//       fetch(`http://localhost:8080/results/${year}?round=${race.round}&limit=100`)
//         .then((response) => {
//           if (!response.ok) {
//             throw new Error(`HTTP error! status: ${response.status}`);
//           }
//           return response.json();
//         })
//         .then((data) => {
//           console.log('Fetched data:', data);
//           const allRaces = data.MRData?.RaceTable?.Races || [];
//           console.log('All races:', allRaces);
//
//           const filteredRace = allRaces.find((r) => r.round === race.round.toString());
//           console.log('Filtered race:', filteredRace);
//
//           if (filteredRace && filteredRace.Results) {
//             setResults(filteredRace.Results);
//           } else {
//             setResults([]);
//           }
//
//           setLoading(false);
//         })
//         .catch((error) => {
//           console.error('Error fetching results:', error);
//           setError('Failed to fetch results. Please try again.');
//           setLoading(false);
//         });
//     }
//   }, [year, race]);
//
//   return (
//     <div key={race ? race.round : 'default'}>
//       {loading ? (
//         <p>Loading results...</p>
//       ) : error ? (
//         <p>{error}</p>
//       ) : year && race ? (
//         <>
//           <h2>Results for {race.raceName} ({year})</h2>
//           {results.length > 0 ? (
//             <ul>
//               {results.map((result, index) => (
//                 <li key={index}>
//                   <strong>Position {result.position}:</strong> {result.Driver.givenName} {result.Driver.familyName} - {result.Constructor.name}
//                   <br />
//                   <em>Time:</em> {result.Time?.time || result.status || 'N/A'}
//                 </li>
//               ))}
//             </ul>
//           ) : (
//             <p>No results found for this race.</p>
//           )}
//         </>
//       ) : (
//         <p>Select a race to see the results.</p>
//       )}
//     </div>
//   );
// }
//
// export default Results;
//
