import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home/Home';
import Ticket from './pages/Ticket/Ticket';
import VirtualQueue from './pages/VirtualQueue/VirtualQueue';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/ticket/:eventId" element={<Ticket />} />
        <Route path="/virtual-queue/:userToken/:eventId" element={<VirtualQueue />} />
      </Routes>
    </Router>
  );
};

export default App;
