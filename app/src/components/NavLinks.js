import React from "react";
import { Link } from "react-router-dom"

function NavLinks() {
  return(
    <div>
      <span><Link to="/">Home</Link></span>
      <span><Link to="/login">Login</Link></span>
      <span><Link to="/vote">Vote</Link></span>
    </div>
  )
}

export default NavLinks;
