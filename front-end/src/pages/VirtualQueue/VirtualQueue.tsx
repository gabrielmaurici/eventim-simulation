import React from 'react';
import { useVirtualQueueSocket } from '../../services/virtual-queue-socket';
import { useParams } from 'react-router-dom';
import './VirtualQueue.css';

const VirtualQueue: React.FC = () => {
  const { userToken } = useParams<{ userToken: string }>();
  const { data, error, isConnected } = useVirtualQueueSocket(`ws/virtual-queue?token=${userToken}`);

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
          <p>Tempo estimado de espera: {data.estimatedWaitTime}</p>
        </div>
      ) : (
        <p className="loading-message">Calculando sua posição...</p>
      )}
    </div>
  );
};

export default VirtualQueue;