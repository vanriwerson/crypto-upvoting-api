import React, { useState, useEffect } from "react";
import { Navigate } from "react-router-dom";
import { requestData, setToken, requestLogin } from "../services/requestAPI";
import "./Voting.css";

function Voting() {
  const [cryptos, setCryptos] = useState([]);
  const [vote, setVote] = useState("1");
  const [voteSent, setVoteSent] = useState(false);

  const getCryptos = () => requestData('cryptos')
    .then((response) => setCryptos(response))
    .catch((error) => console.log(error));

  useEffect(() => { 
    getCryptos();
  }, []);

  useEffect(() => { 
    if (!cryptos.length) getCryptos();
  }, [cryptos]);

  const voting = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem("token")
      setToken(token);
      const result = await requestLogin("/vote", { cryptoId: vote })
      console.log(result);

      setVoteSent(true);
    } catch (error) {
      setVoteSent(false);
    }
  };

  const loading = <option>Loading cryptos...</option>

  if (voteSent) return <Navigate to="/" />;

  return (
    <form className="voting-form">
      <h1>Upvote your favorite Crypto!</h1>
      <select onChange={({ target: { value } }) => setVote(value)} value={vote}>
        {cryptos.length ? (cryptos.map((crypto) => (
          <option key={crypto.id} value={crypto.id}>
            {crypto.name}
          </option>
        ))) : loading }
      </select>

      <button type="submit" className="btn" onClick={(e) => voting(e)}>
        Send my vote
      </button>
    </form>
  );
}

export default Voting;
