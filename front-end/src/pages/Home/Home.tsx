import React from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';
import eventimSimulationApi from '../../services/eventim-simulaiton-api'
import './Home.css';

interface VirtualQueuePostResponse {
  token: string,
  position: number
}

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
  }
];

const Home: React.FC = () => {
  const navigate = useNavigate();

  const handleEventClick = async (eventId: number) => {
    const userToken = Cookies.get('userToken');
    if (userToken) {
      navigate(`/virtual-queue/${userToken}/${eventId}`)
    }

    try {
      const data = await eventimSimulationApi.post<ResponseType | null, VirtualQueuePostResponse>("api/virtual-queue")
      if (data) {
        Cookies.set("userToken", data?.token)
        
        data.position > 0 
          ? navigate(`/virtual-queue/${data.token}/${eventId}`)
          : navigate(`/ticket/${eventId}/${data.token}`);
      }
    } catch (error) {
      alert(error)
    };
  };

  return (
    <div className="home-page">
      <h1>Eventos Disponíveis</h1>
      <div className="events-list">
        {events.map((event) => (
          event.ticketOpen ? (
            <div
              className="event-card open-ticket"
              onClick={() => handleEventClick(event.id)}
            >
              <h2>{event.name}</h2>
              <p><strong>Data:</strong> {event.date}</p>
              <p>{event.description}</p>
              <p className="ticket-status open-ticket">Bilheteira Aberta</p>
            </div>
          ) : (
            <div
              key={event.id}
              className="event-card"
            >
              <h2>{event.name}</h2>
              <p><strong>Data:</strong> {event.date}</p>
              <p>{event.description}</p>
              <p className="ticket-status closed-ticket">Bilheteira Fechada</p>
            </div>
          )
        ))}
      </div>
    </div>
  );
};

export default Home;