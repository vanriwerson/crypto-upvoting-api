import React, { useState, useEffect } from 'react';
import { requestData } from '../../services/requestAPI';
import './RankingBoard.css'

function RankingBoard() {
  const [ranking, setRanking] = useState([]);

  const getRanking = () => requestData('/ranking')
    .then((response) => setRanking(response))
    .catch((error) => console.log(error));

  useEffect(() => {
    getRanking();
  }, [])
  
  const loading = <tr><td>Ranking is loading...</td></tr>

  return (
    <table className='ranking'>
      <thead>
        <tr>
          <th>Crypto</th>
          <th>Votes</th>
        </tr>
      </thead>
      <tbody>
        {
          ranking.length ? (
            ranking.map(({ name, votes }) => (
              <tr key={ name }>
                <td>{ name }</td>
                <td>{ votes }</td>
              </tr>
            ))
          ) : loading
        }
      </tbody>
    </table>
  )
}

export default RankingBoard;
