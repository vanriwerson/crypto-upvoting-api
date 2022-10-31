import React from "react";
import { Link } from "react-router-dom";
import RankingBoard from "../components/RankingBoard/RankingBoard";
import './Home.css';

function Home() {
  return(
    <section>
      <h1>Our users upvote these Cryptos</h1>
      
      <RankingBoard />

      <Link to='/login' className="cta">Timme to choose your favorite!</Link>
    </section>
  )
}

export default Home;
