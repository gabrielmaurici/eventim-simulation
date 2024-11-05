import React from 'react';
import './Home.css';

interface Event {
  id: number;
  name: string;
  date: string;
  description: string;
  ticketOpen: boolean;
}

const events: Event[] = [
  {
    id: 1,
    name: 'Festival de Rock',
    date: '2024-12-01',
    description: 'Um festival incrível com as melhores bandas de rock.',
    ticketOpen: true,
  },
  {
    id: 2,
    name: 'Show do Pop Star',
    date: '2024-11-20',
    description: 'A estrela do pop em um show inesquecível.',
    ticketOpen: false,
  },
  {
    id: 3,
    name: 'Samba na Praça',
    date: '2024-11-15',
    description: 'Uma tarde de samba e alegria na praça central.',
    ticketOpen: false,
  },
  {
    id: 4,
    name: 'Noite de Jazz',
    date: '2024-11-25',
    description: 'Uma noite sofisticada ao som do melhor jazz.',
    ticketOpen: false,
  },
];

const Home: React.FC = () => {
  return (
    <div className="home-page">
      <h1>Eventim</h1>
      <div className="events-list">
        {events.map(event => (
          <div
            key={event.id}
            className={`event-card ${event.ticketOpen ? 'open-ticket' : ''}`}
          >
            <h2>{event.name}</h2>
            <p><strong>Data:</strong> {event.date}</p>
            <p>{event.description}</p>
            {event.ticketOpen && <p className="ticket-status">Bilheteria Aberta</p>}
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;