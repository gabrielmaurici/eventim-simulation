import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import eventimSimulationApi from '../../services/eventim-simulaiton-api';
import TicketsModal from '../../components/TicketsModal/TicketsModal';
import Cookies from 'js-cookie';
import './BuyTickets.css';

interface BuyTicketsInput {
    user_token: string;
}

interface BuyTicketsResponse {
  tickets_purchased: string[];
}

interface Ticket {
  userToken: string;
  eventId: string;
  ticketId: string;
}

const BuyTickets: React.FC = () => {
  const navigate = useNavigate();
  const { userToken } = useParams<{ userToken: string }>();
  const { eventId } = useParams<{ eventId: string }>();
  const [tickets, setTickets] = useState<Ticket[] | null>(null);
  const [timeLeft, setTimeLeft] = useState(60);

  useEffect(() => {
    if (timeLeft <= 0) {
      alert("O tempo para finalizar a compra expirou.");
      navigate('/');
    }

    const timer = setInterval(() => {
      setTimeLeft((prevTime) => prevTime - 1);
    }, 1000);

    return () => clearInterval(timer);
  }, [timeLeft, navigate]);

  const handleConfirmPurchase = async () => {
    try {
      const endpoint = "api/tickets/purchase";
      const body: BuyTicketsInput = { user_token: userToken! };
      
      const data = await eventimSimulationApi.post<BuyTicketsInput, BuyTicketsResponse | null>(endpoint, body);

      setTickets(data!.tickets_purchased.map(ticketId => ({
        userToken: userToken!,
        eventId: eventId!,
        ticketId: ticketId
      })));
    } catch (error) {
      alert("Erro ao comprar ingressos: " + error);
    }
  };

  const handleCloseModal = () => {
    setTickets(null);
    Cookies.remove("userToken");
    navigate('/');
  };

  return (
    <div className="buy-tickets-page">
      <div className="buy-tickets-container">
        <div className="timer-box">Seus ingressos expiram em: {timeLeft} s</div>
        <h1>Finalizar Compra</h1>
        <p>Token do Usuário: {userToken}</p>
        <p>Revise suas informações e confirme a compra.</p>
        <button onClick={handleConfirmPurchase}>Confirmar Compra</button>
      </div>
      {tickets && <TicketsModal tickets={tickets} onClose={handleCloseModal} />}
    </div>
  );
};

export default BuyTickets;