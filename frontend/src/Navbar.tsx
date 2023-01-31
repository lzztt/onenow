import React, { ReactNode } from 'react';
import './navbar.css';

interface Props {
  children: ReactNode;
}

const Navbar = ({ children }: Props) => (
  <nav>
    {React.Children.map(children, child => child)}
  </nav>
);

export default Navbar;
