import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost/',
  headers: {
    'Content-Type': 'application/json',
  },
});

const virtualQueueApiMethods = {
  async post<I, R>(endpoint: string, data?: I | null): Promise<R | null> {
    try {
      const response = await api.post(endpoint, data);
      if (response.status != 200)
        throw new Error("Erro ao realizar request " + endpoint)
      return response.data;
    } catch (error) {
      throw error;
    }
  },
};
  
export default virtualQueueApiMethods;