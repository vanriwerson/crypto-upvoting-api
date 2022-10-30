import React, { useState, useEffect } from "react";
import axios from "axios";

const ENDPOINT = 'http://localhost:5000/cryptos'

function Voting() {
  const [cryptos, setCryptos] = useState([]);

  useEffect(() => {
    axios.get(ENDPOINT)
      .then((response) => {
        setCryptos(response.data);
      })
  }, []);

  return(
    <section>
      <h1>Support your favorite Crypto!</h1>
      <form>
        <select>
          {
            cryptos.map((crypto) => (
              <option
                key={ crypto.id }
                value={ crypto.id }
              >
              { crypto.name }
              </ option>
            ))
          }
        </ select>
      </form>
    </section>
  )
}

export default Voting;
