import React, { useEffect } from 'react';

const Trial = () => {
  useEffect(() => {
    fetch('http://localhost:8080/')
      .then(response => response.json())
      .then(data => {
        console.log(data);
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <div>
      <h1>Check the console for API data</h1>
    </div>
  );
};

export default Trial;
