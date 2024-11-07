import { useState, useEffect } from 'react';

export interface QueueData {
    userToken: string;
    position: number;
    estimatedWaitTime: number;
}

export const useVirtualQueueSocket = (endpoint: string) => {
    const [data, setData] = useState<QueueData | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [isConnected, setIsConnected] = useState<boolean>(false);
    const url = `ws://localhost/${endpoint}`

    useEffect(() => {
        const socket = new WebSocket(url);

        socket.onopen = () => setIsConnected(true);
        socket.onmessage = (event) =>  {
            const data = JSON.parse(event.data);
            const mappedData: QueueData = {
                userToken: data.token,
                position: data.position,
                estimatedWaitTime: data.estimated_wait_time
            };
            setData(mappedData);
        }
        socket.onerror = () => setError('Erro no WebSocket');
        socket.onclose = () => setIsConnected(false);

        return () => socket.close();
    }, [url]);

    return { data, error, isConnected };
};