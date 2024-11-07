import React, { useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import eventimSimulationApi from '../../services/eventim-simulaiton-api'
import './Ticket.css';

interface ReserveInput {
  user_token: string,
  quantity: number
}

const Ticket: React.FC = () => {
  const navigate = useNavigate();
  const { userToken } = useParams<{userToken: string}>();
  const { eventId } = useParams<{eventId: string}>();
  const [ticketCount, setTicketCount] = useState<number>(1);

  const handleTicketChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTicketCount(Number(event.target.value));
  };

  const handleReserve = async () => {
    try {
      const endpoint = "api/tickets/reserve"
      const body: ReserveInput = {
        user_token: userToken!,
        quantity: ticketCount
      };
      
      await eventimSimulationApi.post<ReserveInput, ResponseType | null>(endpoint, body);

      navigate(`/buy-tickets/${userToken}/${eventId}`);
    } catch(error) {
      alert("Ocorreu algum erro ao reservar os ingressos: " + error);
    }
  };

  return (
    <div className="ticket-page">
      <div className="ticket-container">
        <h1>Comprar Ingressos</h1>
        <label>
          Quantidade de Ingressos:
          <input
            type="number"
            value={ticketCount}
            min="1"
            max="5"
            onChange={handleTicketChange}
          />
        </label>
        <button onClick={handleReserve}>Reservar</button>
      </div>
    </div>
  );
};

export default Ticket;