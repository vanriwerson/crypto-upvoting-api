import { Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import Login from '../pages/Login';
import Voting from '../pages/Voting';

function Router() {
  return(
    <Routes >
        <Route exact path="/" element={ <Home /> } />
        <Route exact path="/login" element={ <Login /> } />
        <Route exact path="/vote" element={ <Voting /> } />
    </Routes>
  )
}

export default Router;