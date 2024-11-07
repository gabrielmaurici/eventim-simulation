import React from 'react';
import './TicketsModal.css';

interface Ticket {
  userToken: string;
  eventId: string;
  ticketId: string;
}

interface TicketsModalProps {
  tickets: Ticket[];
  onClose: () => void;
}

const TicketsModal: React.FC<TicketsModalProps> = ({ tickets, onClose }) => {
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>Ingressos Comprados</h2>
        <div className="tickets-list">
          {tickets.map((ticket, index) => (
            <div key={index} className="ticket">
              <p><strong>Token do Usuário:</strong> {ticket.userToken}</p>
              <p><strong>ID do Evento:</strong> {ticket.eventId}</p>
              <p><strong>ID do Ingresso:</strong> {ticket.ticketId}</p>
            </div>
          ))}
        </div>
        <button className="close-button" onClick={onClose}>Fechar</button>
      </div>
    </div>
  );
};

export default TicketsModal;