import { useState, useEffect } from 'react';

export interface QueueData {
    position: number;
    estimatedWaitTime: number;
    userToken: string;
}

export const useVirtualQueueSocket = (endpoint: string) => {
    const [data, setData] = useState<QueueData | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [isConnected, setIsConnected] = useState<boolean>(false);
    const url = `ws://localhost/${endpoint}`

    useEffect(() => {
        const socket = new WebSocket(url);

        socket.onopen = () => setIsConnected(true);
        socket.onmessage = (event) => setData(JSON.parse(event.data));
        socket.onerror = () => setError('Erro no WebSocket');
        socket.onclose = () => setIsConnected(false);

        return () => socket.close();
    }, [url]);

    return { data, error, isConnected };
};