import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import './Ticket.css';

const Ticket: React.FC = () => {
  const { eventId } = useParams<{ eventId: string }>();
  const [ticketCount, setTicketCount] = useState<number>(1);

  const handleTicketChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTicketCount(Number(event.target.value));
  };

  const handlePurchase = () => {
    // lógica de compra de ingressos
    alert(`Comprando ${ticketCount} ingresso(s) para o evento ${eventId}`);
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
        <button onClick={handlePurchase}>Comprar</button>
      </div>
    </div>
  );
};

export default Ticket;