import React from "react";
import { Link } from "react-router-dom"
import "./NavLinks.css"

function NavLinks() {
  return(
    <div className="nav">
      <span><Link className="nav-link" to="/">Home</Link></span>
      <span><Link className="nav-link" to="/login">Login</Link></span>
    </div>
  )
}

export default NavLinks;
