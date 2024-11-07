import React, { useState, useEffect } from 'react';
import { useVirtualQueueSocket } from '../../services/virtual-queue-socket';
import { useNavigate, useParams } from 'react-router-dom';
import './VirtualQueue.css';

const formatTime = (seconds: number) => {
  if (isNaN(seconds) || seconds < 0) {
    return "00:00:00";
  }

  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainingSeconds = seconds % 60;

  const formattedHours = hours < 10 ? `0${hours}` : `${hours}`;
  const formattedMinutes = minutes < 10 ? `0${minutes}` : `${minutes}`;
  const formattedSeconds = remainingSeconds < 10 ? `0${remainingSeconds}` : `${remainingSeconds}`;

  return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
};

const VirtualQueue: React.FC = () => {
  const navigate = useNavigate()
  const { userToken } = useParams<{ userToken: string }>();
  const { eventId } = useParams<{ eventId: string }>();
  const { data, error, isConnected } = useVirtualQueueSocket(`ws/virtual-queue?token=${userToken}`);

  const [remainingTime, setRemainingTime] = useState<number>(0);

  useEffect(() => {
    if (data && data.position == 0) {
      navigate(`/ticket/${eventId}/${userToken}`);
    }

    if (data && data.estimatedWaitTime > 0) {
      setRemainingTime(data.estimatedWaitTime);

      const timer = setInterval(() => {
        setRemainingTime((prevTime) => {
          if (prevTime <= 0) {
            clearInterval(timer);
            return 0;
          }
          return prevTime - 1;
        });
      }, 1000);

      return () => clearInterval(timer);
    }
  }, [data]);

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (!isConnected) {
    return <div className="loading-message">Conectando-se à fila...</div>;
  }

  return (
    <div className="virtual-queue-page">
      <h1>Fila Virtual</h1>
      {data ? (
        <div className="queue-status">
          <h2>Posição na fila: {data.position}</h2>
          <p>Tempo estimado de espera: {formatTime(remainingTime)}</p>
        </div>
      ) : (
        <p className="loading-message">Calculando sua posição...</p>
      )}
    </div>
  );
};

export default VirtualQueue;